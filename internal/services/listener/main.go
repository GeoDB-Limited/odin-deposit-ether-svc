package listener

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/system-contracts/generated"
	eth "github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

// Service defines a service that listens to events of a bridge contract.
type Service struct {
	log      *logrus.Logger
	eth      *ethclient.Client
	contract *bind.BoundContract
	ch       chan<- eth.TransferDetails
}

// New creates a service that listens to events of a bridge contract.
func New(cfg config.Config, contractAddr common.Address, ch chan<- eth.TransferDetails) *Service {
	parsed, err := abi.JSON(strings.NewReader(generated.EtherBridgeABI))
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
		ch:       ch,
	}
}

// Run listens to events of a bridge contract.
func (s *Service) Run(ctx context.Context, wg *sync.WaitGroup) {
	err := s.subscribe(ctx)
	if err != nil {
		panic(errors.Wrap(err, "failed to subscribe on the contract events"))
	}

	wg.Done()
}

// subscribe subscribes on events of a bridge contract.
func (s *Service) subscribe(ctx context.Context) error {
	watchOpts := &bind.WatchOpts{Context: ctx}
	logs, subscription, err := s.contract.WatchLogs(watchOpts, "EtherDeposited")
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to new logs")
	}
	defer subscription.Unsubscribe()

	for event := range logs {
		err = s.processTransfer(ctx, event)
		if err != nil {
			return errors.Wrap(err, "failed to process transfer")
		}
	}

	return nil
}

// processTransfer parses events from smart contract.
func (s *Service) processTransfer(ctx context.Context, event types.Log) error {
	if event.Removed {
		s.log.WithField("event", event.BlockHash).Warn("Log was reverted due to a chain reorganisation")
		return nil
	}

	parsed := new(eth.Transfer)
	if err := s.contract.UnpackLog(parsed, "EtherDeposited", event); err != nil {
		return errors.Wrap(err, "failed to parse log")
	}

	block, err := s.eth.BlockByHash(ctx, event.BlockHash)
	if err != nil {
		return errors.Wrap(err, "failed to get block")
	}

	transferDetails := eth.TransferDetails{
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
	}).Info("User deposited")

	s.ch <- transferDetails

	return nil
}
