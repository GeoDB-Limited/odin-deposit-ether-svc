package config

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/big"
)

// Config defines an interface of global service configurations.
type Config interface {
	Logger() *logrus.Logger
	EthereumClient() *ethclient.Client
	DeployConfig() *DeployConfig
	OdinConfig() *OdinChainConfig
	EthereumConfig() *EthereumChainConfig
	EthereumSigner() (common.Address, *ecdsa.PrivateKey)
	OdinSigner() (sdk.AccAddress, *secp256k1.PrivKey)

	BridgeAddressStorage() string
}

// Config defines global service configurations.
type config struct {
	Log      string         `yaml:"log"`
	Ethereum EthereumConfig `yaml:"ethereum"`
	Odin     OdinConfig     `yaml:"odin"`
	Deploy   DeployConfig   `yaml:"deploy"`
}

// OdinConfig defines the configurations of odin client.
type OdinConfig struct {
	Chain  OdinChainConfig  `yaml:"chain"`
	Signer OdinSignerConfig `yaml:"signer"`

	BridgeAddressStorage string `yaml:"bridge_address_storage"`
}

// OdinSignerConfig defines configs for odin signer
type OdinSignerConfig struct {
	Mnemonic   string `yaml:"mnemonic"`
	Password   string `yaml:"password"`
	Derivation string `yaml:"derivation"`
}

// OdinChainConfig defines configs for Odin chain
type OdinChainConfig struct {
	Endpoint string   `yaml:"endpoint"`
	ChainId  string   `yaml:"chain_id"`
	Denom    string   `yaml:"denom"`
	Memo     string   `yaml:"memo"`
	GasPrice *big.Int `yaml:"gas_price"`
	GasLimit *big.Int `yaml:"gas_limit"`
}

// EthereumConfig defines the configurations of ethereum client.
type EthereumConfig struct {
	Chain  EthereumChainConfig `yaml:"chain"`
	Signer EthereumSigner      `yaml:"signer"`
}

// EthereumSigner defines configs for ethereum signer
type EthereumSigner struct {
	PrivateKey string `yaml:"private_key"`
}

// EthereumChainConfig defines configs for Ethereum
type EthereumChainConfig struct {
	Endpoint string   `yaml:"endpoint"`
	GasLimit *big.Int `yaml:"gas_limit"`
	GasPrice *big.Int `yaml:"gas_price"`
}

// DeployConfig defines the configurations of Deploy service.
type DeployConfig struct {
	RefundGasLimit             *big.Int         `yaml:"refund_gas_limit"`
	DepositingAllowed          bool             `yaml:"depositing_allowed"`
	LockingFundsAllowed        bool             `yaml:"locking_funds_allowed"`
	ClaimingLockedFundsAllowed bool             `yaml:"claiming_locked_funds_allowed"`
	SupportedTokens            []common.Address `yaml:"supported_tokens"`
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

// DeployConfig returns the configurations of deploy service.
func (c *config) DeployConfig() *DeployConfig {
	return &c.Deploy
}

// EthereumClient returns ethereum client.
func (c *config) EthereumClient() *ethclient.Client {
	ethClient, err := ethclient.Dial(c.Ethereum.Chain.Endpoint)
	if err != nil {
		panic(errors.Wrapf(err, "failed to dial %s", c.Ethereum.Chain.Endpoint))
	}

	return ethClient
}

// EthereumConfig returns the configurations of ethereum.
func (c *config) EthereumConfig() *EthereumChainConfig {
	return &c.Ethereum.Chain
}

// EthereumSigner returns address and private key of ethereum signer.
func (c *config) EthereumSigner() (common.Address, *ecdsa.PrivateKey) {
	pk, err := crypto.HexToECDSA(c.Ethereum.Signer.PrivateKey)
	if err != nil {
		panic(errors.Wrap(err, "error casting private key to ECDSA"))
	}

	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(errors.Wrap(err, "error casting public key to ECDSA"))
	}

	accAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return accAddress, pk
}

// OdinConfig returns the configurations of odin.
func (c *config) OdinConfig() *OdinChainConfig {
	return &c.Odin.Chain
}

// OdinSigner returns address and private key of odin signer.
func (c *config) OdinSigner() (sdk.AccAddress, *secp256k1.PrivKey) {
	seed := bip39.NewSeed(c.Odin.Signer.Mnemonic, c.Odin.Signer.Password)
	master, ch := hd.ComputeMastersFromSeed(seed)

	key, err := hd.DerivePrivateKeyForPath(master, ch, c.Odin.Signer.Derivation)
	if err != nil {
		panic(errors.Wrap(err, "failed to derive odin private key for path"))
	}

	pk := secp256k1.PrivKey{Key: key}
	accAddress := sdk.AccAddress(pk.PubKey().Address())

	return accAddress, &pk
}

// BridgeAddressStorage returns the path to bridge address storage.
func (c *config) BridgeAddressStorage() string {
	return c.Odin.BridgeAddressStorage
}
