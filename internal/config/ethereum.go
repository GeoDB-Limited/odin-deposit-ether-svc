package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type EthereumConfig struct {
	Checkpoint    uint64 `fig:"checkpoint"`
	Confirmations int64  `fig:"confirmations"`
}

func (c *config) EthereumConfig() EthereumConfig {

	c.once.Do(func() interface{} {
		var result EthereumConfig

		if err := figure.Out(&result).From(kv.MustGetStringMap(c.getter, "ethereum")).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out deposit"))
		}

		c.ethereumConfig = result
		return nil
	})

	return c.ethereumConfig
}
