package config

import (
	"fmt"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/odin/client"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/big"
)

// Config defines an interface of global service configurations.
type Config interface {
	Logger() *logrus.Logger
	EtherClient() *ethclient.Client
	OdinClient() client.Client
	DeployerConfig() *DeployerConfig
	DepositerConfig() *DepositerConfig
	EthereumConfig() *EthereumConfig
}

// Config defines global service configurations.
type config struct {
	Log       string          `yaml:"log"`
	Ethereum  EthereumConfig  `yaml:"ethereum"`
	Odin      OdinConfig      `yaml:"odin"`
	Deployer  DeployerConfig  `yaml:"deployer"`
	Depositer DepositerConfig `yaml:"depositer"`
}

// DeployerConfig defines the configurations of Deployer service.
type DeployerConfig struct {
	GasLimit        uint64           `yaml:"gas_limit"`
	GasPrice        *big.Int         `yaml:"gas_price"`
	SupportedTokens []common.Address `yaml:"supported_tokens"`
}

// DeployerConfig defines the configurations of odin client.
type OdinConfig struct {
	Endpoint             string `yaml:"odin_endpoint"`
	BridgeAddressStorage string `yaml:"bridge_address_storage"`
}

// DeployerConfig defines the configurations of ethereum client.
type EthereumConfig struct {
	Endpoint   string `yaml:"endpoint"`
	PrivateKey string `yaml:"private_key"`
}

// DepositerConfig defines the configurations of the depositer svc
type DepositerConfig struct {
	GasLimit uint64   `yaml:"gas_limit"`
	GasPrice *big.Int `yaml:"gas_price"`
}

// NewConfig returns global service configurations.
func NewConfig(path string) Config {
	cfg := config{}

	yamlConfig, err := ioutil.ReadFile(path)
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("failed to read config: %s", path)))
	}

	err = yaml.Unmarshal(yamlConfig, &cfg)
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("failed to unmarshal config: %s", path)))
	}

	return &cfg
}

// Logger returns new configured logger.
func (c *config) Logger() *logrus.Logger {
	level, err := logrus.ParseLevel(c.Log)
	if err != nil {
		panic(errors.Wrapf(err, "failed to parse logging level %s", c.Log))
	}

	logger := logrus.New()
	logger.SetLevel(level)

	return logger
}

// EtherClient returns ethereum client.
func (c *config) EtherClient() *ethclient.Client {
	ethClient, err := ethclient.Dial(c.Ethereum.Endpoint)
	if err != nil {
		panic(errors.Wrapf(err, "failed to dial %s", c.Ethereum.Endpoint))
	}

	return ethClient
}

// OdinClient returns odin client.
func (c *config) OdinClient() client.Client {
	clientConn, err := grpc.Dial(c.Odin.Endpoint, grpc.WithInsecure())
	if err != nil {
		panic(errors.Wrapf(err, "failed to dial %s", c.Odin.Endpoint))
	}

	serviceClient := tx.NewServiceClient(clientConn)
	odinClient := client.New(&serviceClient, c.Odin.BridgeAddressStorage, c.Logger())

	return odinClient
}

// DeployerConfig returns the configurations of deployer service.
func (c *config) DeployerConfig() *DeployerConfig {
	return &c.Deployer
}

// DepositerConfig returns the configurations of depositer service.
func (c *config) DepositerConfig() *DepositerConfig {
	return &c.Depositer
}

// DepositerConfig returns the configurations of ethereum.
func (c *config) EthereumConfig() *EthereumConfig {
	return &c.Ethereum
}
