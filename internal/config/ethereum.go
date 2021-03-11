package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Ether interface {
	EtherClient() *ethclient.Client
}

type ether struct {
	getter kv.Getter
	once   comfig.Once
	client *ethclient.Client
}

func NewEther(getter kv.Getter) Ether {
	return &ether{getter: getter}
}

func (e *ether) EtherClient() *ethclient.Client {
	e.once.Do(func() interface{} {
		var config struct {
			Endpoint string `fig:"endpoint,required"`
		}

		if err := figure.Out(&config).From(kv.MustGetStringMap(e.getter, "ethereum")).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out rpc"))
		}

		eth, err := ethclient.Dial(config.Endpoint)
		if err != nil {
			panic(fmt.Sprintf("failed to dial %s", config.Endpoint))
		}

		e.client = eth
		return nil
	})

	return e.client
}
