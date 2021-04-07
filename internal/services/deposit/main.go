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
	"time"
)

// Service defines a service that allows exchanging ETH and ERC20 for odin.
type Service struct {
	config   config.Config
	log      *logrus.Logger
	eth      *ethclient.Client
	contract *generated.Bridge
	odin     client.Client
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

	go s.subscribeETHTransfer(ctx)
	s.subscribeERC20Transfer(ctx)
}

// subscribeETHTransfer subscribes on events of a bridge contract.
func (s *Service) subscribeETHTransfer(ctx context.Context) {
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

		block, err := s.eth.BlockByHash(ctx, event.Raw.BlockHash)
		if err != nil {
			panic(errors.Wrap(err, "failed to get block"))
		}

		s.log.WithFields(logrus.Fields{
			"ethereum_address": event.UserAddress,
			"odin_address":     event.OdinAddress,
			"amount":           event.DepositAmount,
			"block_time":       time.Unix(int64(block.Time()), 0).UTC(),
		}).Info("User deposited ETH")

		s.processETHTransfer(ctx, *event)
	}
}

// processETHTransfer exchanges ETH to odin and claims withdrawal
func (s *Service) processETHTransfer(ctx context.Context, details generated.BridgeETHDeposited) {
	rate, err := s.odin.GetExchangeRate("ETH")
	if err != nil {
		panic(errors.Wrap(err, "failed to get the exchange rate"))
	}

	// TODO: implement logic
	withdrawalAmount := rate.Mul(rate, details.DepositAmount)

	s.log.WithFields(logrus.Fields{
		"odin_address":      details.OdinAddress,
		"deposit_amount":    details.DepositAmount,
		"rate":              rate,
		"withdrawal_amount": withdrawalAmount,
	}).Info("Exchanged")

	if err := s.odin.ClaimWithdrawal(details.OdinAddress, withdrawalAmount); err != nil {
		s.log.WithFields(logrus.Fields{
			"eth_address":       details.UserAddress,
			"odin_address":      details.OdinAddress,
			"deposit_amount":    details.DepositAmount,
			"rate":              rate,
			"withdrawal_amount": withdrawalAmount,
		}).Error(err, "Failed to claim withdrawal")

		if err := s.payBackETH(ctx, details.UserAddress, details.DepositAmount); err != nil {
			panic(errors.Wrapf(err, "failed to pay back"))
		}

	}
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

// subscribeERC20Transfer subscribes on events of a bridge contract.
func (s *Service) subscribeERC20Transfer(ctx context.Context) {
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

		block, err := s.eth.BlockByHash(ctx, event.Raw.BlockHash)
		if err != nil {
			panic(errors.Wrap(err, "failed to get block"))
		}

		s.log.WithFields(logrus.Fields{
			"ethereum_address": event.UserAddress,
			"odin_address":     event.OdinAddress,
			"token_address":    event.TokenAddress,
			"amount":           event.DepositAmount,
			"block_time":       time.Unix(int64(block.Time()), 0).UTC(),
		}).Info("User deposited ERC20")

		s.processERC20Transfer(ctx, *event)
	}
}

// processERC20Transfer exchanges ERC20 to odin and claims withdrawal
func (s *Service) processERC20Transfer(ctx context.Context, details generated.BridgeERC20Deposited) {
	rate, err := s.odin.GetExchangeRate(details.Symbol)
	if err != nil {
		panic(errors.Wrap(err, "failed to get the exchange rate"))
	}

	// TODO: implement logic
	withdrawalAmount := rate.Mul(rate, details.DepositAmount)

	s.log.WithFields(logrus.Fields{
		"odin_address":      details.OdinAddress,
		"deposit_amount":    details.DepositAmount,
		"rate":              rate,
		"withdrawal_amount": withdrawalAmount,
	}).Info("Exchanged")

	if err := s.odin.ClaimWithdrawal(details.OdinAddress, withdrawalAmount); err != nil {
		s.log.WithFields(logrus.Fields{
			"eth_address":       details.UserAddress,
			"odin_address":      details.OdinAddress,
			"deposit_amount":    details.DepositAmount,
			"rate":              rate,
			"withdrawal_amount": withdrawalAmount,
		}).Error(err, "Failed to claim withdrawal")

		if err := s.payBackERC20(ctx, details.UserAddress, details.TokenAddress, details.DepositAmount); err != nil {
			panic(errors.Wrapf(err, "failed to pay back"))
		}
	}
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
