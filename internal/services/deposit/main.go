package deposit

import (
	"context"
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
	ETHExchangeSymbol       = "ETH" // used to get an exchange rate of ETH
	ETHPrecision      int64 = 18    // the precision of ETH

	ETHDepositType = iota
	ERC20DepositType
)

// Service defines a service that allows exchanging ETH and ERC20 for odin.
type Service struct {
	config   config.Config
	context  context.Context
	logger   *logrus.Logger
	ethereum *ethclient.Client
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
	TokenPrecision  int64
	DepositType     int
}

// New creates a service that allows exchanging ETH and ERC20 for odin.
func New(ctx context.Context, cfg config.Config) *Service {
	odinClient := client.New(ctx, cfg).WithSigner()
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
		context:  ctx,
		logger:   cfg.Logger(),
		odin:     odinClient,
		ethereum: etherClient,
		contract: bridge,
	}
}

// Run performs events listening and querying the Odin minting module.
func (s *Service) Run() {
	withdrawals := make(chan WithdrawalDetails)
	go s.subscribeETHTransfer(withdrawals)
	go s.subscribeERC20Transfer(withdrawals)
	s.logger.Info("Starting deposits processing service...")
	s.processTransfer(withdrawals)
}

// subscribeETHTransfer subscribes on events of a bridge contract.
func (s *Service) subscribeETHTransfer(withdrawals chan<- WithdrawalDetails) {
	logs := make(chan *generated.BridgeETHDeposited)
	subscription, err := s.contract.WatchETHDeposited(&bind.WatchOpts{Context: s.context}, logs, []common.Address{})
	if err != nil {
		panic(errors.Wrap(err, "failed to subscribe on event logs"))
	}
	defer subscription.Unsubscribe()

	for event := range logs {
		if event.Raw.Removed {
			s.logger.WithField("block_hash", event.Raw.BlockHash).Warn("Log was reverted due to a chain reorganisation")
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
func (s *Service) subscribeERC20Transfer(withdrawals chan<- WithdrawalDetails) {
	logs := make(chan *generated.BridgeERC20Deposited)
	subscription, err := s.contract.WatchERC20Deposited(
		&bind.WatchOpts{Context: s.context},
		logs,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		panic(errors.Wrap(err, "failed to subscribe on event logs"))
	}
	defer subscription.Unsubscribe()

	for event := range logs {
		if event.Raw.Removed {
			s.logger.WithField("block_hash", event.Raw.BlockHash).Warn("Log was reverted due to a chain reorganisation")
			continue
		}

		withdrawals <- WithdrawalDetails{
			EthereumAddress: event.UserAddress,
			OdinAddress:     event.OdinAddress,
			DepositAmount:   event.DepositAmount,
			TokenAddress:    event.TokenAddress,
			TokenSymbol:     event.Symbol,
			TokenPrecision:  int64(event.TokenPrecision),
			DepositType:     ERC20DepositType,
		}
	}
}

// processTransfer handles events from ethereum smart contract
func (s *Service) processTransfer(withdrawals <-chan WithdrawalDetails) {
	for w := range withdrawals {
		switch w.DepositType {
		case ETHDepositType:
			if err := s.processETHTransfer(w); err != nil {
				s.logger.WithFields(logrus.Fields{
					"eth_address": w.EthereumAddress,
					"amount":      w.DepositAmount,
				}).Error(err, "Failed to process ETH transfer")

				if err := s.payBackETH(w.EthereumAddress, w.DepositAmount); err != nil {
					s.logger.WithFields(logrus.Fields{
						"eth_address": w.EthereumAddress,
						"amount":      w.DepositAmount,
					}).Error(err, "Failed to pay back ETH")
				}
			}
		case ERC20DepositType:
			if err := s.processERC20Transfer(w); err != nil {
				s.logger.WithFields(logrus.Fields{
					"eth_address":   w.EthereumAddress,
					"amount":        w.DepositAmount,
					"token_address": w.TokenAddress,
				}).Error(err, "Failed to process ERC20 transfer")

				if err := s.payBackERC20(w.EthereumAddress, w.TokenAddress, w.DepositAmount); err != nil {
					s.logger.WithFields(logrus.Fields{
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
		return errors.Wrapf(
			err,
			"failed to claim withdrawal for user: %s amount: %s",
			withdrawal.OdinAddress,
			withdrawalAmount,
		)
	}

	s.logger.WithFields(logrus.Fields{
		"ethereum_address": withdrawal.EthereumAddress,
		"odin_address":     withdrawal.OdinAddress,
		"amount":           withdrawal.DepositAmount,
	}).Info("User deposited ETH")

	return nil
}

// processERC20Transfer exchanges ERC20 and claims withdrawal from odin mint module
func (s *Service) processERC20Transfer(withdrawal WithdrawalDetails) error {
	withdrawalAmount, err := s.exchangeERC20(
		withdrawal.EthereumAddress,
		withdrawal.DepositAmount,
		withdrawal.TokenSymbol,
		withdrawal.TokenPrecision,
	)
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
		return errors.Wrapf(
			err,
			"failed to claim withdrawal for user: %s amount: %s",
			withdrawal.OdinAddress,
			withdrawalAmount,
		)
	}

	s.logger.WithFields(logrus.Fields{
		"ethereum_address": withdrawal.EthereumAddress,
		"odin_address":     withdrawal.OdinAddress,
		"amount":           withdrawal.DepositAmount,
		"token_address":    withdrawal.TokenAddress,
		"token_symbol":     withdrawal.TokenSymbol,
		"token_precision":  withdrawal.TokenPrecision,
	}).Info("User deposited ERC20")

	return nil
}

// exchangeETH exchanges deposited ETH to odin tokens
func (s *Service) exchangeETH(ethereumAddress common.Address, amount *big.Int) (sdk.Coin, error) {
	rate, err := s.odin.GetExchangeRate(ETHExchangeSymbol)
	if err != nil {
		return sdk.Coin{}, errors.Wrap(err, "failed to get the exchange rate")
	}

	withdrawalAmount, err := s.exchange(amount, rate, ETHPrecision)
	if err != nil {
		return sdk.Coin{}, errors.Wrapf(
			err,
			"failed to exchange the deposit: %s with rate: %s",
			amount.String(),
			rate.String(),
		)
	}

	if withdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(0)) {
		return sdk.Coin{}, errors.New("insufficient deposit amount")
	}

	s.logger.WithFields(logrus.Fields{
		"eth_address":       ethereumAddress,
		"deposit_amount":    amount,
		"rate":              rate,
		"withdrawal_amount": withdrawalAmount.Amount,
	}).Info("Exchanged ETH")

	return withdrawalAmount, nil
}

// exchangeERC20 exchanges deposited ERC20 to odin tokens
func (s *Service) exchangeERC20(
	ethereumAddress common.Address,
	amount *big.Int,
	tokenSymbol string,
	tokenPrecision int64,
) (sdk.Coin, error) {
	rate, err := s.odin.GetExchangeRate(tokenSymbol)
	if err != nil {
		return sdk.Coin{}, errors.Wrapf(err, "failed to get the exchange rate for %s", tokenSymbol)
	}

	withdrawalAmount, err := s.exchange(amount, rate, tokenPrecision)
	if err != nil {
		return sdk.Coin{}, errors.Wrapf(
			err,
			"failed to exchange the deposit: %s with rate: %s",
			amount.String(),
			rate.String(),
		)
	}

	if withdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(0)) {
		return sdk.Coin{}, errors.New("insufficient deposit amount")
	}

	s.logger.WithFields(logrus.Fields{
		"eth_address":       ethereumAddress,
		"deposit_amount":    amount,
		"token_symbol":      tokenSymbol,
		"token_precision":   tokenPrecision,
		"rate":              rate,
		"withdrawal_amount": withdrawalAmount.Amount,
	}).Info("Exchanged ERC20")

	return withdrawalAmount, nil
}

// exchange calculates new coin with the given exchange rate
func (s *Service) exchange(amount *big.Int, rate sdk.Dec, precision int64) (sdk.Coin, error) {
	withdrawalAmount := sdk.NewDecFromBigIntWithPrec(amount, precision).Mul(rate)
	return sdk.NewCoin(s.config.OdinConfig().Denom, withdrawalAmount.TruncateInt()), nil
}

// payBackETH pays back the deposit amount
func (s *Service) payBackETH(userAddress common.Address, amount *big.Int) error {
	opts, err := s.getTxOpts()
	if err != nil {
		return errors.Wrap(err, "failed to get tx options")
	}

	tx, err := s.contract.SetRefundETH(opts, userAddress, amount)
	if err != nil {
		return errors.Wrap(err, "failed send tx to pay back ETH")
	}

	s.logger.WithFields(logrus.Fields{
		"eth_address":    userAddress,
		"deposit_amount": amount,
		"tx_hash":        tx.Hash(),
	}).Info("Payed back ETH")

	return nil
}

// payBackERC20 pays back the deposit amount
func (s *Service) payBackERC20(userAddress common.Address, tokenAddress common.Address, amount *big.Int) error {
	opts, err := s.getTxOpts()
	if err != nil {
		return errors.Wrap(err, "failed to get tx options")
	}

	tx, err := s.contract.SetRefundERC20(opts, userAddress, tokenAddress, amount)
	if err != nil {
		return errors.Wrap(err, "failed send tx to pay back ETH")
	}

	s.logger.WithFields(logrus.Fields{
		"eth_address":    userAddress,
		"deposit_amount": amount,
		"token_address":  tokenAddress,
		"tx_hash":        tx.Hash(),
	}).Info("Payed back ERC20")

	return nil
}

// getTxOpts returns the options to broadcast signed tx
func (s *Service) getTxOpts() (*bind.TransactOpts, error) {
	_, pk := s.config.EthereumSigner()

	chainId, err := s.ethereum.NetworkID(s.context)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}
	opts, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create the transactor")
	}

	gasPrice, err := s.ethereum.SuggestGasPrice(s.context)
	if err != nil {
		return nil, errors.Wrap(err, "failed to suggest gas price")
	}

	opts.GasLimit = s.config.DeployConfig().RefundGasLimit.Uint64()
	opts.GasPrice = gasPrice

	return opts, nil
}
