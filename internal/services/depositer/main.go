package depositer

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/listener"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/odin/client"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// Service defines a service that allows exchanging ETH and ERC20 for odin.
type Service struct {
	config   config.Config
	log      *logan.Entry
	listener *listener.Service
	odin     *client.Client
}

// New creates a service that allows exchanging ETH and ERC20 for odin.
func New(cfg config.Config) *Service {
	return &Service{
		config:   cfg,
		log:      cfg.Log(),
		odin:     cfg.OdinClient(),
		listener: listener.New(cfg),
	}
}

// Run performs events listening and querying the Odin  minting module.
func (s *Service) Run(ctx context.Context) (err error) {
	localCtx, cancel := context.WithCancel(ctx)

	defer func() {
		if rvr := recover(); rvr != nil {
			cancel()
			err = errors.Wrap(errors.FromPanic(rvr), "service panicked")
		}
	}()

	if err := s.listener.Run(localCtx); err != nil {
		err = errors.Wrap(err, "failed to run event listener")
	}

	return nil
}
