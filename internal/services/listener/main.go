package listener

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"math/big"
	"sync"
)

type Service struct {
	log    *logan.Entry
	client *ethclient.Client
	cfg    config.EthereumConfig

	ch chan TransactionInfo

	sync.RWMutex
	decimals uint
}

type TransactionInfo struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func New(cfg config.Config) *Service {
	ch := make(chan TransactionInfo, 10)

	return &Service{
		log:     cfg.Log(),
		cfg:     cfg.EthereumConfig(),
		ch:      ch,
		client:  cfg.EtherClient(),
		RWMutex: sync.RWMutex{},
	}
}

func (s *Service) Run(ctx context.Context) {}
