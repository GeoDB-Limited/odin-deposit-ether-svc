package listener

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/system-contracts/generated"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"math/big"
	"time"
)

// Service defines a service that listens to events of a bridge contract.
type Service struct {
	log *logrus.Logger
	eth *ethclient.Client
	ch  chan TransferDetails
}

// TransferDetails defines unpacked data of the event.
type TransferDetails struct {
	DepositAmount   *big.Int
	UserAddress     common.Address
	OdinAddress     string
	TransactionHash string
	BlockTime       time.Time
}

// New creates a service that listens to events of a bridge contract.
func New(cfg config.Config) *Service {
	ch := make(chan TransferDetails)
	return &Service{
		log: cfg.Logger(),
		eth: cfg.EtherClient(),
		ch:  ch,
	}
}

// Run listens to events of a bridge contract.
func (s *Service) Run(ctx context.Context, contractAddress common.Address) error {
	err := s.subscribe(ctx, contractAddress)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe on the contract events")
	}

	return nil
}

// subscribe subscribes on events of a bridge contract.
func (s *Service) subscribe(ctx context.Context, contractAddress common.Address) error {
	contract, err := generated.NewEtherBridge(contractAddress, s.eth)
	if err != nil {
		return errors.Wrap(err, "failed to create a contract instance")
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	subscription, err := s.eth.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to new logs")
	}
	defer subscription.Unsubscribe()

	go func() error {
		for event := range logs {
			err = s.processTransfer(contract, event, ctx)
			if err != nil {
				return errors.Wrap(err, "failed to process transfer")
			}
		}
		return nil
	}()

	return nil
}

// processTransfer parses events from smart contract.
func (s *Service) processTransfer(contract *generated.EtherBridge, event types.Log, ctx context.Context) error {
	if event.Removed {
		s.log.WithField("event", event.BlockHash).Warn("Log was reverted due to a chain reorganisation")
		return nil
	}

	parsed, err := contract.ParseEtherDeposited(event)
	if err != nil {
		return errors.Wrap(err, "failed to parse log")
	}

	block, err := s.eth.BlockByHash(ctx, event.BlockHash)
	if err != nil {
		return errors.Wrap(err, "failed to get block")
	}

	transferDetails := TransferDetails{
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
