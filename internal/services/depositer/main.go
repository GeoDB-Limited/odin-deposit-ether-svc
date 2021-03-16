package depositer

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/types"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/listener"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/odin/client"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Service defines a service that allows exchanging ETH and ERC20 for odin.
type Service struct {
	config config.Config
	log    *logrus.Logger
	odin   client.Client
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

	transferDetails := make(chan types.TransferDetails)
	go s.odin.ClaimMinting(transferDetails)

	listenerService := listener.New(s.config, contractAddress, transferDetails)
	go listenerService.Run(ctx)

	return nil
}
