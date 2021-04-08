package deposit

import (
	"context"
	app "github.com/GeoDB-Limited/odin-core/app"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/system-contracts/generated"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/odin/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"math/big"
)

const (
	ETHDepositType = iota
	ERC20DepositType
)

// Service defines a service that allows exchanging ETH and ERC20 for odin.
type Service struct {
	config   config.Config
	log      *logrus.Logger
	eth      *ethclient.Client
	contract *generated.Bridge
	odin     client.Client
}

// WithdrawalDetails defines withdrawal options
type WithdrawalDetails struct {
	EthereumAddress common.Address
	OdinAddress     string
	DepositAmount   *big.Int // amount of user's deposit
	TokenAddress    common.Address
	TokenSymbol     string
	DepositType     int
}

// New creates a service that allows exchanging ETH and ERC20 for odin.
func New(cfg config.Config) *Service {
	odinClient := client.New(cfg)
	bridgeAddr, err := odinClient.GetBridgeAddress()
	if err != nil {
		panic(errors.Wrap(err, "failed to get bridge address"))
	}

	etherClient := cfg.EthereumClient()
	bridge, err := generated.NewBridge(bridgeAddr, etherClient)
	if err != nil {
		panic(errors.Wrap(err, "failed to create bridge instance"))
	}

	return &Service{
		config:   cfg,
		log:      cfg.Logger(),
		odin:     odinClient,
		eth:      etherClient,
		contract: bridge,
	}
}

// Run performs events listening and querying the Odin minting module.
func (s *Service) Run(ctx context.Context) {
	app.SetBech32AddressPrefixesAndBip44CoinType(sdk.GetConfig())

	withdrawals := make(chan WithdrawalDetails)
	go s.subscribeETHTransfer(ctx, withdrawals)
	go s.subscribeERC20Transfer(ctx, withdrawals)
	s.processTransfer(ctx, withdrawals)
}

// subscribeETHTransfer subscribes on events of a bridge contract.
func (s *Service) subscribeETHTransfer(ctx context.Context, withdrawals chan<- WithdrawalDetails) {
	logs := make(chan *generated.BridgeETHDeposited)
	subscription, err := s.contract.WatchETHDeposited(&bind.WatchOpts{Context: ctx}, logs, []common.Address{})
	if err != nil {
		panic(errors.Wrap(err, "failed to subscribe on event logs"))
	}
	defer subscription.Unsubscribe()

	for event := range logs {
		if event.Raw.Removed {
			s.log.WithField("event", event.Raw.BlockHash).Warn("Log was reverted due to a chain reorganisation")
			continue
		}

		withdrawals <- WithdrawalDetails{
			EthereumAddress: event.UserAddress,
			OdinAddress:     event.OdinAddress,
			DepositAmount:   event.DepositAmount,
			DepositType:     ETHDepositType,
		}
	}
}

// subscribeERC20Transfer subscribes on events of a bridge contract.
func (s *Service) subscribeERC20Transfer(ctx context.Context, withdrawals chan<- WithdrawalDetails) {
	logs := make(chan *generated.BridgeERC20Deposited)
	subscription, err := s.contract.WatchERC20Deposited(&bind.WatchOpts{Context: ctx}, logs, []common.Address{}, []common.Address{})
	if err != nil {
		panic(errors.Wrap(err, "failed to subscribe on event logs"))
	}
	defer subscription.Unsubscribe()

	for event := range logs {
		if event.Raw.Removed {
			s.log.WithField("event", event.Raw.BlockHash).Warn("Log was reverted due to a chain reorganisation")
			continue
		}

		withdrawals <- WithdrawalDetails{
			EthereumAddress: event.UserAddress,
			OdinAddress:     event.OdinAddress,
			DepositAmount:   event.DepositAmount,
			TokenAddress:    event.TokenAddress,
			TokenSymbol:     event.Symbol,
			DepositType:     ERC20DepositType,
		}
	}
}

// processTransfer handles events from ethereum smart contract
func (s *Service) processTransfer(ctx context.Context, withdrawals <-chan WithdrawalDetails) {
	for w := range withdrawals {
		switch w.DepositType {
		case ETHDepositType:
			if err := s.processETHTransfer(w); err != nil {
				s.log.WithFields(logrus.Fields{
					"eth_address": w.EthereumAddress,
					"amount":      w.DepositAmount,
				}).Error(err, "Failed to process ETH transfer")

				if err := s.payBackETH(ctx, w.EthereumAddress, w.DepositAmount); err != nil {
					s.log.WithFields(logrus.Fields{
						"eth_address": w.EthereumAddress,
						"amount":      w.DepositAmount,
					}).Error(err, "Failed to pay back ETH")
				}
			}
		case ERC20DepositType:
			if err := s.processERC20Transfer(w); err != nil {
				s.log.WithFields(logrus.Fields{
					"eth_address":   w.EthereumAddress,
					"amount":        w.DepositAmount,
					"token_address": w.TokenAddress,
				}).Error(err, "Failed to process ERC20 transfer")

				if err := s.payBackERC20(ctx, w.EthereumAddress, w.TokenAddress, w.DepositAmount); err != nil {
					s.log.WithFields(logrus.Fields{
						"eth_address":   w.EthereumAddress,
						"amount":        w.DepositAmount,
						"token_address": w.TokenAddress,
					}).Error(err, "Failed to pay back ERC20")
				}
			}
		}
	}
}

// processETHTransfer exchanges ETH and claims withdrawal from odin mint module
func (s *Service) processETHTransfer(withdrawal WithdrawalDetails) error {
	withdrawalAmount, err := s.exchangeETH(withdrawal.EthereumAddress, withdrawal.DepositAmount)
	if err != nil {
		return errors.Wrapf(
			err,
			"failed to exchange deposited ETH for user: %s deposit: %s",
			withdrawal.OdinAddress,
			withdrawal.DepositAmount,
		)
	}

	if err := s.odin.ClaimWithdrawal(withdrawal.OdinAddress, withdrawalAmount); err != nil {
		s.log.WithFields(logrus.Fields{
			"odin_address":      withdrawal.OdinAddress,
			"eth_address":       withdrawal.EthereumAddress,
			"withdrawal_amount": withdrawalAmount,
		}).Error(err, "Failed to claim withdrawal")

		return errors.Wrapf(err, "failed to claim withdrawal for user: %s ", withdrawal.OdinAddress)
	}

	return nil
}

// processERC20Transfer exchanges ERC20 and claims withdrawal from odin mint module
func (s *Service) processERC20Transfer(withdrawal WithdrawalDetails) error {
	withdrawalAmount, err := s.exchangeERC20(withdrawal.EthereumAddress, withdrawal.DepositAmount, withdrawal.TokenSymbol)
	if err != nil {
		return errors.Wrapf(
			err,
			"failed to exchange deposited ERC20 for user: %s deposit: %s %s",
			withdrawal.OdinAddress,
			withdrawal.DepositAmount,
			withdrawal.TokenSymbol,
		)
	}

	if err := s.odin.ClaimWithdrawal(withdrawal.OdinAddress, withdrawalAmount); err != nil {
		s.log.WithFields(logrus.Fields{
			"odin_address":      withdrawal.OdinAddress,
			"eth_address":       withdrawal.EthereumAddress,
			"withdrawal_amount": withdrawalAmount,
		}).Error(err, "Failed to claim withdrawal")

		return errors.Wrapf(err, "failed to claim withdrawal for user: %s ", withdrawal.OdinAddress)
	}

	return nil
}

// exchangeETH exchanges deposited ETH to odin tokens
func (s *Service) exchangeETH(ethereumAddress common.Address, amount *big.Int) (*big.Int, error) {
	rate, err := s.odin.GetExchangeRate("ETH")
	if err != nil {
		return &big.Int{}, errors.Wrap(err, "failed to get the exchange rate")
	}

	// TODO: implement exchange logic
	withdrawalAmount := rate.Mul(rate, amount)

	s.log.WithFields(logrus.Fields{
		"eth_address":       ethereumAddress,
		"deposit_amount":    amount,
		"rate":              rate,
		"withdrawal_amount": withdrawalAmount,
	}).Info("Exchanged")

	return withdrawalAmount, nil
}

// exchangeERC20 exchanges deposited ERC20 to odin tokens
func (s *Service) exchangeERC20(ethereumAddress common.Address, amount *big.Int, tokenSymbol string) (*big.Int, error) {
	rate, err := s.odin.GetExchangeRate(tokenSymbol)
	if err != nil {
		return &big.Int{}, errors.Wrap(err, "failed to get the exchange rate")
	}

	// TODO: implement logic
	withdrawalAmount := rate.Mul(rate, amount)

	s.log.WithFields(logrus.Fields{
		"eth_address":       ethereumAddress,
		"deposit_amount":    amount,
		"token_symbol":      tokenSymbol,
		"rate":              rate,
		"withdrawal_amount": withdrawalAmount,
	}).Info("Exchanged")

	return withdrawalAmount, nil
}

// payBackETH pays back the deposit amount
func (s *Service) payBackETH(ctx context.Context, userAddress common.Address, amount *big.Int) error {
	opts, err := s.getTxOpts(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get tx options")
	}

	_, err = s.contract.PayBackETH(opts, userAddress, amount)
	if err != nil {
		return errors.Wrap(err, "failed send tx to pay back ETH")
	}

	s.log.WithFields(logrus.Fields{
		"eth_address":    userAddress,
		"deposit_amount": amount,
	}).Info(err, "Payed back ETH")

	return nil
}

// payBackERC20 pays back the deposit amount
func (s *Service) payBackERC20(
	ctx context.Context,
	userAddress common.Address,
	tokenAddress common.Address,
	amount *big.Int,
) error {
	opts, err := s.getTxOpts(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get tx options")
	}

	_, err = s.contract.PayBackERC20(opts, userAddress, tokenAddress, amount)
	if err != nil {
		return errors.Wrap(err, "failed send tx to pay back ETH")
	}

	s.log.WithFields(logrus.Fields{
		"eth_address":    userAddress,
		"deposit_amount": amount,
		"token_address":  tokenAddress,
	}).Info(err, "Payed back ERC20")

	return nil
}

// getTxOpts returns the options to broadcast signed tx
func (s *Service) getTxOpts(ctx context.Context) (*bind.TransactOpts, error) {
	_, pk := s.config.EthereumSigner()

	chainId, err := s.eth.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}
	opts, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create the transactor")
	}

	cfg := s.config.EthereumConfig()
	opts.GasLimit = cfg.GasLimit.Uint64()
	opts.GasPrice = cfg.GasPrice

	return opts, nil
}
