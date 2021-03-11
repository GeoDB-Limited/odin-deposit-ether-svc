package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Config interface {
	DeployerConfig() DeployerConfig

	comfig.Logger
	Ether
	Odin
}

type config struct {
	deployerConfig comfig.Once

	getter kv.Getter
	once   comfig.Once
	comfig.Logger
	Ether
	Odin
}

func NewConfig(getter kv.Getter) Config {
	return &config{
		getter: getter,
		Logger: comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Ether:  NewEther(getter),
		Odin:   NewOdin(getter),
	}
}
