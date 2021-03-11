package listener

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/system-contracts/generated"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/odin/client"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"
)

type Service struct {
	log  *logan.Entry
	eth  *ethclient.Client
	odin *client.Client

	ch chan TransferDetails

	sync.RWMutex
}

type Transfer struct {
	From     common.Address
	To       string
	Value    *big.Int
	Decimals int64
	Raw      types.Log
}

type TransferDetails struct {
	TransactionHash string
	Destination     string
	BlockTime       time.Time
	Amount          *big.Int
	Decimals        int64
}

func New(cfg config.Config) *Service {
	ch := make(chan TransferDetails)
	return &Service{
		log:  cfg.Log(),
		eth:  cfg.EtherClient(),
		odin: cfg.OdinClient(),
		ch:   ch,

		RWMutex: sync.RWMutex{},
	}
}

func (s *Service) Run(ctx context.Context) {
	go running.WithBackOff(ctx, s.log, "transfer-listener", func(ctx context.Context) error {
		err := s.subscribe(ctx)
		if err != nil {
			return errors.Wrap(err, "failed to subscribe on the contract events")
		}
		s.log.Info("Finished listening new transfers")

		return nil
	}, time.Minute, time.Minute, time.Hour)
}

func (s *Service) subscribe(ctx context.Context) error {
	contractAddress, err := s.odin.GetBridgeAddress()
	if err != nil {
		return errors.Wrap(err, "failed to get contract address")
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{*contractAddress},
	}

	logs := make(chan types.Log)
	subscription, err := s.eth.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to new logs")
	}

	defer subscription.Unsubscribe()

runner:
	for {
		select {
		case err := <-subscription.Err():
			return errors.Wrap(err, "subscription returned error")
		case event, ok := <-logs:
			if !ok {
				return errors.New("channel closed unexpectedly")
			}

			err = s.processTransfer(ctx, event)
			if err != nil {
				return errors.Wrap(err, "failed to process transfer")
			}
		case <-ctx.Done():
			break runner
		}
	}

	return nil
}

func (s *Service) processTransfer(ctx context.Context, event types.Log) error {
	if event.Removed {
		return nil
	}
	s.Lock()
	defer s.Unlock()

	contractAbi, err := abi.JSON(strings.NewReader(generated.BridgeABI))
	if err != nil {
		log.Fatal(err)
	}

	parsed := new(Transfer)
	if err = contractAbi.UnpackIntoInterface(&parsed, "Transfer", event.Data); err != nil {
		return errors.Wrap(err, "failed to unpack log", logan.F{
			"event_data":   event.Data,
			"event_topics": event.Topics,
		})
	}

	block, err := s.eth.BlockByHash(ctx, event.BlockHash)
	if err != nil {
		return errors.Wrap(err, "failed to get block", logan.F{
			"block_hash":   event.BlockHash.String(),
			"block_number": event.BlockNumber,
		})
	}

	s.ch <- TransferDetails{
		TransactionHash: event.TxHash.String(),
		Destination:     parsed.To,
		Amount:          parsed.Value,
		BlockTime:       time.Unix(int64(block.Time()), 0),
		Decimals:        parsed.Decimals,
	}

	return nil
}
