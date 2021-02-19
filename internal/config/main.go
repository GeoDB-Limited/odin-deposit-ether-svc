package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Config interface {
	EthereumConfig() EthereumConfig

	comfig.Logger
	Ether
}

type config struct {
	ethereumConfig EthereumConfig

	getter kv.Getter
	once   comfig.Once
	comfig.Logger
	Ether
}

func NewConfig(getter kv.Getter) Config {
	return &config{
		getter: getter,
		Logger: comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Ether:  NewEther(getter),
	}
}
