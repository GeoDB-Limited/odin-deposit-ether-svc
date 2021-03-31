package depositer

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/types"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/exchanger"
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

// Run performs events listening and querying the Odin minting module.
func (s *Service) Run(ctx context.Context) error {
	contractAddress, err := s.odin.GetBridgeAddress()
	if err != nil {
		return errors.Wrap(err, "failed to get contract address")
	}

	ethTransferDetails := make(chan types.ETHTransferDetails)
	erc20TransferDetails := make(chan types.ERC20TransferDetails)
	listenerService := listener.New(s.config, contractAddress, ethTransferDetails, erc20TransferDetails)
	listenerService.Run(ctx)

	withdrawalDetails := make(chan types.WithdrawalDetails)
	exchangerService := exchanger.New(s.config, ethTransferDetails, erc20TransferDetails, withdrawalDetails)
	exchangerService.Run()

	s.odin.ClaimWithdrawal(withdrawalDetails)

	return nil
}
