package listener

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/system-contracts/generated"
	data "github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

// Service defines a service that listens to events of a bridge contract.
type Service struct {
	log      *logrus.Logger
	eth      *ethclient.Client
	contract *bind.BoundContract

	ethTransferDetails   chan<- data.ETHTransferDetails
	erc20TransferDetails chan<- data.ERC20TransferDetails
}

// HandleEventFunc defines a function to handle the events
type HandleEventFunc func(context.Context, types.Log) error

// New creates a service that listens to events of a bridge contract.
func New(
	cfg config.Config,
	contractAddr common.Address,
	ethTransferDetails chan<- data.ETHTransferDetails,
	erc20TransferDetails chan<- data.ERC20TransferDetails,
) *Service {
	parsed, err := abi.JSON(strings.NewReader(generated.BridgeABI))
	if err != nil {
		panic(errors.Wrap(err, "failed to parse contract ABI"))
	}

	etherClient := cfg.EtherClient()
	contract := bind.NewBoundContract(
		contractAddr,
		parsed,
		etherClient,
		etherClient,
		etherClient,
	)

	return &Service{
		log:      cfg.Logger(),
		eth:      etherClient,
		contract: contract,

		ethTransferDetails:   ethTransferDetails,
		erc20TransferDetails: erc20TransferDetails,
	}
}

// Run listens to events of a bridge contract.
func (s *Service) Run(ctx context.Context) {
	go s.subscribe(ctx, "ETHDeposited", s.ProcessETHTransfer)
	go s.subscribe(ctx, "ERC20Deposited", s.ProcessERC20Transfer)
}

// subscribe subscribes on events of a bridge contract.
func (s *Service) subscribe(ctx context.Context, eventName string, handler HandleEventFunc) {
	watchOpts := &bind.WatchOpts{Context: ctx}

	logs, subscription, err := s.contract.WatchLogs(watchOpts, eventName)
	if err != nil {
		panic(errors.Wrap(err, "failed to subscribe on event logs"))
	}
	defer subscription.Unsubscribe()

	for event := range logs {
		if event.Removed {
			s.log.WithField("event", event.BlockHash).Warn("Log was reverted due to a chain reorganisation")
			continue
		}
		err = handler(ctx, event)
		if err != nil {
			panic(errors.Wrap(err, "failed to process transfer"))
		}
	}
}

// ProcessETHTransfer handles the event of depositing ETH
func (s *Service) ProcessETHTransfer(ctx context.Context, event types.Log) error {
	block, err := s.eth.BlockByHash(ctx, event.BlockHash)
	if err != nil {
		return errors.Wrap(err, "failed to get block")
	}

	parsed := new(data.ETHTransfer)
	if err := s.contract.UnpackLog(parsed, "ETHDeposited", event); err != nil {
		return errors.Wrap(err, "failed to parse log")
	}

	transferDetails := data.ETHTransferDetails{
		UserAddress:     parsed.UserAddress,
		OdinAddress:     parsed.OdinAddress,
		DepositAmount:   parsed.DepositAmount,
		TransactionHash: event.TxHash.String(),
		BlockTime:       time.Unix(int64(block.Time()), 0),
	}

	s.log.WithFields(logrus.Fields{
		"ethereum_address": transferDetails.UserAddress,
		"odin_address":     transferDetails.OdinAddress,
		"amount":           transferDetails.DepositAmount,
		"block_time":       transferDetails.BlockTime.UTC().String(),
	}).Info("User deposited ETH")

	s.ethTransferDetails <- transferDetails

	return nil
}

// ProcessERC20Transfer handles the event of depositing ERC20 compatible tokens
func (s *Service) ProcessERC20Transfer(ctx context.Context, event types.Log) error {
	block, err := s.eth.BlockByHash(ctx, event.BlockHash)
	if err != nil {
		return errors.Wrap(err, "failed to get block")
	}

	parsed := new(data.ERC20Transfer)
	if err := s.contract.UnpackLog(parsed, "ERC20Deposited", event); err != nil {
		return errors.Wrap(err, "failed to parse log")
	}

	transferDetails := data.ERC20TransferDetails{
		UserAddress:     parsed.UserAddress,
		OdinAddress:     parsed.OdinAddress,
		TokenAddress:    parsed.TokenAddress,
		DepositAmount:   parsed.DepositAmount,
		TransactionHash: event.TxHash.String(),
		BlockTime:       time.Unix(int64(block.Time()), 0),
	}

	s.log.WithFields(logrus.Fields{
		"ethereum_address": transferDetails.UserAddress,
		"odin_address":     transferDetails.OdinAddress,
		"token_address":    transferDetails.TokenAddress,
		"amount":           transferDetails.DepositAmount,
		"block_time":       transferDetails.BlockTime.UTC().String(),
	}).Info("User deposited ERC20")

	s.erc20TransferDetails <- transferDetails

	return nil
}
