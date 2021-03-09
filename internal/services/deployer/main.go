package deployer

import (
	"context"
	"crypto/ecdsa"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/system-contracts/generated"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/odin/client"
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
	odin   *client.OdinClient
}

func New(cfg config.Config) *Service {
	/*	odin, err := odin.NewConnector(cfg.Odin()).Builder()
		if err != nil {
			cfg.Log().WithError(err).Fatal("failed to make builder")
		}*/

	return &Service{
		log:    cfg.Log(),
		eth:    cfg.EtherClient(),
		config: cfg,
		// odin: odin,
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

	contractAddress, err := s.deployContract(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to deploy contract")
	}

	fields := logan.F{}
	fields["contract"] = contractAddress.Hex()
	s.log.WithFields(fields).Info("contract deployed")

	/*	if err := s.odin.SetEtherBridgeAddress(contractAddress); err != nil {
			return errors.Wrap(err, "failed to set contract address")
		}
	*/
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

	contractAddress, _, _, err := generated.DeployBridge(
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
		[]common.Address{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to submit contract tx")
	}

	return &contractAddress, nil
}
