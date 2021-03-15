package depositer

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/listener"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/odin/client"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Service defines a service that allows exchanging ETH and ERC20 for odin.
type Service struct {
	config   config.Config
	log      *logrus.Logger
	listener *listener.Service
	odin     client.Client
}

// New creates a service that allows exchanging ETH and ERC20 for odin.
func New(cfg config.Config) *Service {
	return &Service{
		config:   cfg,
		log:      cfg.Logger(),
		odin:     cfg.OdinClient(),
		listener: listener.New(cfg),
	}
}

// Run performs events listening and querying the Odin  minting module.
func (s *Service) Run(ctx context.Context) (err error) {
	if err := s.listener.Run(ctx); err != nil {
		err = errors.Wrap(err, "failed to run event listener")
	}

	return nil
}
