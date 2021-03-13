package client

import (
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"io/ioutil"
)

// Client defines typed wrapper for the cosmos sdk service client.
type Client struct {
	client  *tx.ServiceClient
	storage string
}

// New creates a client that uses the given cosmos sdk service client.
func New(client *tx.ServiceClient, storage string) *Client {
	return &Client{
		client:  client,
		storage: storage,
	}
}

// SetBridgeAddress sets an address of the bridge contract to the storage.
func (c *Client) SetBridgeAddress(address common.Address) error {
	if err := ioutil.WriteFile(c.storage, address.Bytes(), 0777); err != nil {
		return errors.Wrap(err, "failed to add the address to the storage")
	}
	return nil
}

// GetBridgeAddress returns an address of the bridge contract.
func (c *Client) GetBridgeAddress() (*common.Address, error) {
	data, err := ioutil.ReadFile(c.storage)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get the address")
	}

	address := common.BytesToAddress(data)
	return &address, nil
}
