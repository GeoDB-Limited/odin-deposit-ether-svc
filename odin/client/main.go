package client

import (
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
)

// Client defines an interface for the wrapped cosmos sdk service client.
type Client interface {
	SetBridgeAddress(common.Address) error
	GetBridgeAddress() (common.Address, error)
	ClaimWithdrawal(string, *big.Int) error
	GetExchangeRate(string) (*big.Int, error)
}

// client defines typed wrapper for the cosmos sdk service client.
type client struct {
	client  *tx.ServiceClient
	storage string
	log     *logrus.Logger
}

// New creates a client that uses the given cosmos sdk service client.
func New(serviceClient *tx.ServiceClient, storage string, log *logrus.Logger) Client {
	return &client{
		client:  serviceClient,
		storage: storage,
		log:     log,
	}
}

// SetBridgeAddress sets an address of the bridge contract to the storage.
func (c *client) SetBridgeAddress(address common.Address) error {
	if err := ioutil.WriteFile(c.storage, address.Bytes(), 0777); err != nil {
		return errors.Wrap(err, "failed to add the address to the storage")
	}
	return nil
}

// GetBridgeAddress returns an address of the bridge contract.
func (c *client) GetBridgeAddress() (common.Address, error) {
	data, err := ioutil.ReadFile(c.storage)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "failed to get the address")
	}

	address := common.BytesToAddress(data)
	return address, nil
}

// ClaimWithdrawal claims minting from Odin
func (c *client) ClaimWithdrawal(odinAddr string, amount *big.Int) error {
	return nil
}

func (c *client) GetExchangeRate(key string) (*big.Int, error) {
	return big.NewInt(1), nil
}
