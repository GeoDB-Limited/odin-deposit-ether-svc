package depositer

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"gitlab.com/distributed_lab/logan/v3"
)

type Service struct {
	log    *logan.Entry
	config config.Config
}

func New(cfg config.Config) *Service {
	return &Service{
		log:    cfg.Log(),
		config: cfg,
	}
}

func (s *Service) Run(ctx context.Context) {}
