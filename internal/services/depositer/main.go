package depositer

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/listener"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/odin/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"math/big"
	"time"
)

// Service defines a service that allows exchanging ETH and ERC20 for odin.
type Service struct {
	config config.Config
	log    *logrus.Logger
	odin   client.Client
}

// TransferDetails defines unpacked data of the event.
type TransferDetails struct {
	DepositAmount   *big.Int
	UserAddress     common.Address
	OdinAddress     string
	TransactionHash string
	BlockTime       time.Time
}

// New creates a service that allows exchanging ETH and ERC20 for odin.
func New(cfg config.Config) *Service {
	return &Service{
		config: cfg,
		log:    cfg.Logger(),
		odin:   cfg.OdinClient(),
	}
}

// Run performs events listening and querying the Odin  minting module.
func (s *Service) Run(ctx context.Context) error {
	contractAddress, err := s.odin.GetBridgeAddress()
	if err != nil {
		return errors.Wrap(err, "failed to get contract address")
	}

	transferDetails := make(chan TransferDetails)
	go s.odin.ClaimMinting(transferDetails)

	listenerService := listener.New(s.config, transferDetails, *contractAddress)
	go listenerService.Run(ctx)

	return nil
}
