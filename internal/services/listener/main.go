package listener

import (
	"context"
	"fmt"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/system-contracts/generated"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/odin/client"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/big"
	"sync"
	"time"
)

type Service struct {
	log  *logan.Entry
	eth  *ethclient.Client
	odin *client.Client
	ch   chan TransferDetails

	sync.RWMutex
}

type TransferDetails struct {
	UserAddress     common.Address
	OdinAddress     string
	DepositAmount   *big.Int
	TransactionHash string
	BlockTime       time.Time
}

func New(cfg config.Config) *Service {
	ch := make(chan TransferDetails)
	return &Service{
		log:     cfg.Log(),
		eth:     cfg.EtherClient(),
		odin:    cfg.OdinClient(),
		ch:      ch,
		RWMutex: sync.RWMutex{},
	}
}

func (s *Service) Run(ctx context.Context) error {
	err := s.subscribe(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe on the contract events")
	}
	s.log.Info("Finished listening new transfers")

	return nil
}

func (s *Service) subscribe(ctx context.Context) error {
	contractAddress, err := s.odin.GetBridgeAddress()
	if err != nil {
		return errors.Wrap(err, "failed to get contract address")
	}

	contract, err := generated.NewEtherBridge(*contractAddress, s.eth)
	if err != nil {
		return errors.Wrap(err, "failed to create a contract instance")
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

			err = s.processTransfer(*contract, event, ctx)
			if err != nil {
				return errors.Wrap(err, "failed to process transfer")
			}
		case <-ctx.Done():
			break runner
		}
	}

	return nil
}

func (s *Service) processTransfer(contract generated.EtherBridge, event types.Log, ctx context.Context) error {
	if event.Removed {
		return nil
	}
	s.Lock()
	defer s.Unlock()

	parsed, err := contract.ParseEtherDeposited(event)
	if err != nil {
		return errors.Wrap(err, "failed to parse log")
	}

	block, err := s.eth.BlockByHash(ctx, event.BlockHash)
	if err != nil {
		return errors.Wrap(err, "failed to get block", logan.F{
			"block_hash":   event.BlockHash.String(),
			"block_number": event.BlockNumber,
		})
	}

	transferDetails := TransferDetails{
		UserAddress:     parsed.UserAddress,
		OdinAddress:     parsed.OdinAddress,
		DepositAmount:   parsed.DepositAmount,
		TransactionHash: event.TxHash.String(),
		BlockTime:       time.Unix(int64(block.Time()), 0),
	}

	s.log.Info(fmt.Sprintf(
		"%s deposited %s ETH to %s at %s",
		transferDetails.UserAddress,
		transferDetails.DepositAmount,
		transferDetails.OdinAddress,
		transferDetails.BlockTime.UTC().String(),
	))

	// s.ch <- transferDetails

	return nil
}

func (s *Service) GetTransferDetails() <-chan TransferDetails {
	return s.ch
}
