package depositer

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/system-contracts/generated"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/types"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/odin/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	odinClient := cfg.OdinClient()
	bridgeAddr, err := odinClient.GetBridgeAddress()
	if err != nil {
		panic(errors.Wrap(err, "failed to get bridge address"))
	}

	etherClient := cfg.EtherClient()
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
	ethTransferDetails := make(chan types.ETHTransferDetails)
	erc20TransferDetails := make(chan types.ERC20TransferDetails)

	go s.subscribeETHTransfer(ctx, ethTransferDetails)
	go s.exchangeETH(ctx, ethTransferDetails)

	go s.subscribeERC20Transfer(ctx, erc20TransferDetails)
	s.exchangeERC20(ctx, erc20TransferDetails)
}

// subscribe subscribes on events of a bridge contract.
func (s *Service) subscribeETHTransfer(ctx context.Context, ethTransferDetails chan<- types.ETHTransferDetails) {
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

		transferDetails := types.ETHTransferDetails{
			EthAddress:      event.UserAddress,
			OdinAddress:     event.OdinAddress,
			DepositAmount:   event.DepositAmount,
			TransactionHash: event.Raw.TxHash.String(),
			BlockTime:       time.Unix(int64(block.Time()), 0),
		}

		s.log.WithFields(logrus.Fields{
			"ethereum_address": transferDetails.EthAddress,
			"odin_address":     transferDetails.OdinAddress,
			"amount":           transferDetails.DepositAmount,
			"block_time":       transferDetails.BlockTime.UTC().String(),
		}).Info("User deposited ETH")

		ethTransferDetails <- transferDetails
	}
}

// subscribe subscribes on events of a bridge contract.
func (s *Service) subscribeERC20Transfer(ctx context.Context, erc20TransferDetails chan<- types.ERC20TransferDetails) {
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

		transferDetails := types.ERC20TransferDetails{
			EthAddress:      event.UserAddress,
			OdinAddress:     event.OdinAddress,
			TokenAddress:    event.TokenAddress,
			DepositAmount:   event.DepositAmount,
			TransactionHash: event.Raw.TxHash.String(),
			BlockTime:       time.Unix(int64(block.Time()), 0),
		}

		s.log.WithFields(logrus.Fields{
			"ethereum_address": transferDetails.EthAddress,
			"odin_address":     transferDetails.OdinAddress,
			"token_address":    transferDetails.TokenAddress,
			"amount":           transferDetails.DepositAmount,
			"block_time":       transferDetails.BlockTime.UTC().String(),
		}).Info("User deposited ERC20")

		erc20TransferDetails <- transferDetails
	}
}

// exchangeETH exchanges ETH to odin
func (s *Service) exchangeETH(ctx context.Context, ethTransferDetails <-chan types.ETHTransferDetails) {
	for data := range ethTransferDetails {
		rate, err := s.odin.GetExchangeRate("ETH")
		if err != nil {
			panic(errors.Wrap(err, "failed to get the exchange rate"))
		}
		withdrawalAmount := rate.Mul(rate, data.DepositAmount) // TODO: implement calculations

		s.log.WithFields(logrus.Fields{
			"odin_address":      data.OdinAddress,
			"deposit_amount":    data.DepositAmount,
			"rate":              rate,
			"withdrawal_amount": withdrawalAmount,
		}).Info("Exchange")

		if err := s.odin.ClaimWithdrawal(data.OdinAddress, withdrawalAmount); err != nil {
			if err := s.payBackETH(ctx, data.EthAddress, data.DepositAmount); err != nil {
				panic(errors.Wrapf(err, "failed to pay back"))
			}

			s.log.WithFields(logrus.Fields{
				"eth_address":       data.EthAddress,
				"odin_address":      data.OdinAddress,
				"deposit_amount":    data.DepositAmount,
				"rate":              rate,
				"withdrawal_amount": withdrawalAmount,
				"time":              data.BlockTime,
			}).Error(err, "Failed to claim withdrawal")
		}
	}
}

// exchangeERC20 exchanges ERC20 to odin
func (s *Service) exchangeERC20(ctx context.Context, erc20TransferDetails <-chan types.ERC20TransferDetails) {
	for data := range erc20TransferDetails {
		rate, err := s.odin.GetExchangeRate(data.TokenSymbol)
		if err != nil {
			panic(errors.Wrap(err, "failed to get the exchange rate"))
		}
		withdrawalAmount := rate.Mul(rate, data.DepositAmount) // TODO: implement calculations

		s.log.WithFields(logrus.Fields{
			"odin_address":      data.OdinAddress,
			"deposit_amount":    data.DepositAmount,
			"rate":              rate,
			"withdrawal_amount": withdrawalAmount,
		}).Info("Exchange")

		if err := s.odin.ClaimWithdrawal(data.OdinAddress, withdrawalAmount); err != nil {
			if err := s.payBackERC20(ctx, data.EthAddress, data.TokenAddress, data.DepositAmount); err != nil {
				panic(errors.Wrapf(err, "failed to pay back"))
			}

			s.log.WithFields(logrus.Fields{
				"eth_address":       data.EthAddress,
				"odin_address":      data.OdinAddress,
				"deposit_amount":    data.DepositAmount,
				"rate":              rate,
				"withdrawal_amount": withdrawalAmount,
				"time":              data.BlockTime,
			}).Error(err, "Failed to claim withdrawal")
		}
	}
}

func (s *Service) getTxOpts(ctx context.Context) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(s.config.EthereumConfig().PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse private key")
	}
	chainId, err := s.eth.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create the transactor")
	}

	cfg := s.config.DepositerConfig()
	opts.GasLimit = cfg.GasLimit
	opts.GasPrice = cfg.GasPrice

	return opts, nil
}

func (s *Service) payBackETH(ctx context.Context, userAddr common.Address, amount *big.Int) error {
	opts, err := s.getTxOpts(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get tx options")
	}

	_, err = s.contract.PayBackETH(opts, userAddr, amount)
	if err != nil {
		return errors.Wrap(err, "failed send tx to pay back ETH")
	}

	return nil
}

func (s *Service) payBackERC20(ctx context.Context, userAddr common.Address, tokenAddr common.Address, amount *big.Int) error {
	opts, err := s.getTxOpts(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get tx options")
	}

	_, err = s.contract.PayBackERC20(opts, userAddr, tokenAddr, amount)
	if err != nil {
		return errors.Wrap(err, "failed send tx to pay back ETH")
	}

	return nil
}
