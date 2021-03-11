package depositer

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/listener"
	"gitlab.com/distributed_lab/logan/v3"
)

type Service struct {
	config   config.Config
	log      *logan.Entry
	listener *listener.Service
}

func New(cfg config.Config) *Service {
	return &Service{
		config:   cfg,
		log:      cfg.Log(),
		listener: listener.New(cfg),
	}
}

func (s *Service) Run(ctx context.Context) (err error) {
	localCtx, _ := context.WithCancel(ctx)
	/*
		defer func() {
			if rvr := recover(); rvr != nil {
				// we are spending actual ether here,
				// so in case of emergency abandon the operations completely
				cancel()
				err = errors.Wrap(errors.FromPanic(rvr), "service panicked")
			}
		}()
	*/

	go s.listener.Run(localCtx)
	s.log.Info("Started listening for deposits")

	return nil
}
