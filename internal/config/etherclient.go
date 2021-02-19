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
	value  *ethclient.Client
}

func NewEther(getter kv.Getter) Ether {
	return &ether{getter: getter}
}

func (h *ether) EtherClient() *ethclient.Client {
	h.once.Do(func() interface{} {
		var config struct {
			Endpoint string `fig:"endpoint,required"`
		}

		if err := figure.Out(&config).From(kv.MustGetStringMap(h.getter, "rpc")).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out rpc"))
		}

		eth, err := ethclient.Dial(config.Endpoint)
		if err != nil {
			panic(fmt.Sprintf("failed to dial %s", config.Endpoint))
		}

		h.value = eth
		return nil
	})

	return h.value
}
