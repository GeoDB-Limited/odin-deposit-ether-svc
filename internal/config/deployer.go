package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/big"
)

// DeployerConfig defines the configurations of Deployer service.
type DeployerConfig struct {
	KeyPair  string   `fig:"private_key,required"`
	GasLimit *big.Int `fig:"gas_limit,required"`
}

// DeployerConfig returns the configurations of Deployer service.
func (c *config) DeployerConfig() DeployerConfig {
	return c.deployerConfig.Do(func() interface{} {
		var result DeployerConfig

		if err := figure.Out(&result).From(kv.MustGetStringMap(c.getter, "deployer")).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out deployer"))
		}

		return result
	}).(DeployerConfig)
}
