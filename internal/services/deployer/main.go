package deployer

import (
	"context"
	"crypto/ecdsa"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/system-contracts/generated"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/big"
)

type Service struct {
	log    *logan.Entry
	config config.Config
	eth    *ethclient.Client
}

func New(cfg config.Config) *Service {
	return &Service{
		log:    cfg.Log(),
		eth:    cfg.EtherClient(),
		config: cfg,
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

	contract, err := s.deployContract(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to deploy contract")
	}

	fields := logan.F{}
	fields["contract"] = contract.Hex()
	s.log.WithFields(fields).Info("contract deployed")

	return nil
}

func (s *Service) deployContract(ctx context.Context) (*common.Address, error) {
	privateKey, err := crypto.HexToECDSA(s.config.DeployerConfig().KeyPair)
	if err != nil {
		return nil, errors.Wrap(err, "error casting private key to ECDSA")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	signer := func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	}

	nonce, err := s.eth.PendingNonceAt(ctx, crypto.PubkeyToAddress(*publicKeyECDSA))
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve account nonce")
	}

	address, _, _, err := generated.DeployOdinBridge(
		&bind.TransactOpts{
			From:     crypto.PubkeyToAddress(*publicKeyECDSA),
			Nonce:    big.NewInt(int64(nonce)),
			Signer:   signer,
			Value:    big.NewInt(0),
			GasPrice: s.config.DeployerConfig().GasPrice,
			GasLimit: s.config.DeployerConfig().GasLimit.Uint64(),
			Context:  context.TODO(),
		},
		s.eth,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to submit contract tx")
	}

	return &address, nil
}
