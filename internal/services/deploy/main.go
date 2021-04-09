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
	"io/ioutil"
	"math/big"
)

// Service defines a service that deploys a bridge contract.
type Service struct {
	config config.Config
	log    *logrus.Logger
	eth    *ethclient.Client
}

// New creates a service that deploys a bridge contract.
func New(cfg config.Config) *Service {
	return &Service{
		config: cfg,
		log:    cfg.Logger(),
		eth:    cfg.EthereumClient(),
	}
}

// Run performs deploying a bridge smart contract.
func (s *Service) Run(ctx context.Context) (err error) {
	contractAddress, err := s.deployContract(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to deploy contract")
	}

	if err := s.saveBridgeAddress(*contractAddress); err != nil {
		return errors.Wrap(err, "failed to set contract address")
	}

	return nil
}

// deployContract deploys a bridge contract.
func (s *Service) deployContract(ctx context.Context) (*common.Address, error) {
	chainId, err := s.eth.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}

	address, pk := s.config.EthereumSigner()

	txOpts, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transaction options")
	}

	nonce, err := s.eth.PendingNonceAt(ctx, address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a nonce")
	}

	txOpts.Nonce = new(big.Int).SetUint64(nonce)

	ethConfig := s.config.EthereumConfig()
	txOpts.GasLimit = ethConfig.GasLimit.Uint64()

	gasPrice, err := s.eth.SuggestGasPrice(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to suggest gas price")
	}

	txOpts.GasPrice = gasPrice

	contractAddress, tx, _, err := generated.DeployBridge(
		txOpts,
		s.eth,
		s.config.DeployConfig().SupportedTokens,
		s.config.DepositCompensation(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to submit contract tx")
	}

	s.log.WithFields(logrus.Fields{
		"tx_hash":          tx.Hash(),
		"contract_address": contractAddress.Hex(),
	}).Info("Contract deployed")

	return &contractAddress, nil
}

// SetBridgeAddress sets an address of the bridge contract to the storage.
func (s *Service) saveBridgeAddress(address common.Address) error {
	if err := ioutil.WriteFile(s.config.BridgeAddressStorage(), address.Bytes(), 0777); err != nil {
		return errors.Wrap(err, "failed to add the address to the storage")
	}
	return nil
}
