package deployer

import (
	"context"
	"crypto/ecdsa"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/system-contracts/generated"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/odin/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"math/big"
)

// Service defines a service that deploys a bridge contract.
type Service struct {
	config config.Config
	log    *logrus.Logger
	eth    *ethclient.Client
	odin   client.Client
}

// New creates a service that deploys a bridge contract.
func New(cfg config.Config) *Service {
	return &Service{
		config: cfg,
		log:    cfg.Logger(),
		eth:    cfg.EtherClient(),
		odin:   cfg.OdinClient(),
	}
}

// Run performs deploying a bridge smart contract.
func (s *Service) Run(ctx context.Context) (err error) {
	contractAddress, err := s.deployContract(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to deploy contract")
	}

	if err := s.odin.SetBridgeAddress(*contractAddress); err != nil {
		return errors.Wrap(err, "failed to set contract address")
	}

	s.log.WithField("contract", contractAddress.Hex()).Info("Contract deployed")

	return nil
}

// deployContract deploys a bridge contract.
func (s *Service) deployContract(ctx context.Context) (*common.Address, error) {
	cfg := s.config.DeployerConfig()

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "error casting private key to ECDSA")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.Wrap(err, "error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.eth.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a nonce")
	}

	chainId, err := s.eth.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transaction options")
	}

	gasPrice, err := s.eth.SuggestGasPrice(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to suggest gas price")
	}

	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = cfg.GasLimit
	transactOpts.Nonce = big.NewInt(int64(nonce))

	contractAddress, _, _, err := generated.DeployBridge(
		transactOpts,
		s.eth,
		cfg.SupportedTokens,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to submit contract tx")
	}

	return &contractAddress, nil
}
