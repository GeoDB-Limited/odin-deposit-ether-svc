package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

// Config defines an interface of global service configurations.
type Config interface {
	DeployerConfig() DeployerConfig

	comfig.Logger
	Ether
	Odin
}

// Config defines global service configurations.
type config struct {
	deployerConfig comfig.Once

	getter kv.Getter
	once   comfig.Once
	comfig.Logger
	Ether
	Odin
}

// NewConfig returns global service configurations.
func NewConfig(getter kv.Getter) Config {
	return &config{
		getter: getter,
		Logger: comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Ether:  NewEther(getter),
		Odin:   NewOdin(getter),
	}
}
