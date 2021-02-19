package depositer

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"gitlab.com/distributed_lab/logan/v3"
	"sync"
)

type Service struct {
	log    *logan.Entry
	config config.Config
	sync.WaitGroup
}

func New(cfg config.Config) *Service {
	return &Service{
		log:       cfg.Log(),
		config:    cfg,
		WaitGroup: sync.WaitGroup{},
	}
}

func (s *Service) Run(ctx context.Context) {

}
