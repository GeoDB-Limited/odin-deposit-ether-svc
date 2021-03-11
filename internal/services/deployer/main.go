package deployer

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/system-contracts/generated"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/odin/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/big"
)

type Service struct {
	config config.Config
	log    *logan.Entry
	eth    *ethclient.Client
	odin   *client.Client
}

func New(cfg config.Config) *Service {
	return &Service{
		config: cfg,
		log:    cfg.Log(),
		eth:    cfg.EtherClient(),
		odin:   cfg.OdinClient(),
	}
}

func (s *Service) Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		if rvr := recover(); rvr != nil {
			cancel()
			err = errors.Wrap(errors.FromPanic(rvr), "service panicked")
		}
	}()

	contractAddress, err := s.deployContract()
	if err != nil {
		return errors.Wrap(err, "failed to deploy contract")
	}

	if err := s.odin.SetBridgeAddress(*contractAddress); err != nil {
		return errors.Wrap(err, "failed to set contract address")
	}

	s.log.WithField("contract", contractAddress.Hex()).Info("Contract deployed")

	return nil
}

func (s *Service) deployContract() (*common.Address, error) {
	privateKey, err := crypto.HexToECDSA(s.config.DeployerConfig().KeyPair)
	if err != nil {
		return nil, errors.Wrap(err, "error casting private key to ECDSA")
	}

	chainId, err := s.eth.NetworkID(context.Background())
	if err != nil {
		return nil, errors.New("failed to get chain id")
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, errors.New("failed to create transaction options")
	}

	transactOpts.GasPrice = s.config.DeployerConfig().GasPrice
	transactOpts.GasLimit = s.config.DeployerConfig().GasLimit.Uint64()
	transactOpts.Value = big.NewInt(0)

	contractAddress, _, _, err := generated.DeployBridge(
		transactOpts,
		s.eth,
		[]common.Address{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to submit contract tx")
	}

	return &contractAddress, nil
}
