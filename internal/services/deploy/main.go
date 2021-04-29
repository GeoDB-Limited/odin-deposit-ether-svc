package deploy

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/system-contracts/generated"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"math/big"
	"os"
)

// Service defines a service that deploys a bridge contract.
type Service struct {
	config   config.Config
	context  context.Context
	logger   *logrus.Logger
	ethereum *ethclient.Client
}

// New creates a service that deploys a bridge contract.
func New(ctx context.Context, cfg config.Config) *Service {
	return &Service{
		config:   cfg,
		context:  ctx,
		logger:   cfg.Logger(),
		ethereum: cfg.EthereumClient(),
	}
}

// Run performs deploying a bridge smart contract.
func (s *Service) Run() (err error) {
	contractAddress, err := s.deployContract()
	if err != nil {
		return errors.Wrap(err, "failed to deploy contract")
	}

	if err := s.saveBridgeAddress(*contractAddress); err != nil {
		return errors.Wrap(err, "failed to set contract address")
	}

	return nil
}

// deployContract deploys a bridge contract.
func (s *Service) deployContract() (*common.Address, error) {
	chainId, err := s.ethereum.NetworkID(s.context)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}

	address, pk := s.config.EthereumSigner()

	txOpts, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transaction options")
	}

	nonce, err := s.ethereum.PendingNonceAt(s.context, address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a nonce")
	}

	txOpts.Nonce = new(big.Int).SetUint64(nonce)

	ethConfig := s.config.EthereumConfig()
	txOpts.GasLimit = ethConfig.GasLimit.Uint64()

	gasPrice, err := s.ethereum.SuggestGasPrice(s.context)
	if err != nil {
		return nil, errors.Wrap(err, "failed to suggest gas price")
	}

	txOpts.GasPrice = gasPrice

	contractAddress, tx, _, err := generated.DeployBridge(
		txOpts,
		s.ethereum,
		s.config.DeployConfig().SupportedTokens,
		s.config.DepositCompensation(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to submit contract tx")
	}

	s.logger.WithFields(logrus.Fields{
		"tx_hash":          tx.Hash(),
		"contract_address": contractAddress.Hex(),
	}).Info("Contract deployed")

	return &contractAddress, nil
}

// SetBridgeAddress sets an address of the bridge contract to the storage.
func (s *Service) saveBridgeAddress(address common.Address) error {
	f, err := os.OpenFile(s.config.BridgeAddressStorage(), os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return errors.Wrap(err, "failed to add the address to the storage")
	}

	if _, err := f.WriteString(address.String()); err != nil {
		return errors.Wrap(err, "failed to add the address to the storage")
	}
	return nil
}
