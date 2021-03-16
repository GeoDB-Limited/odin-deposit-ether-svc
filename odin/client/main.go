package client

import (
	"fmt"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/depositer"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"io/ioutil"
)

// Client defines an interface for the wrapped cosmos sdk service client.
type Client interface {
	SetBridgeAddress(common.Address) error
	GetBridgeAddress() (*common.Address, error)
	ClaimMinting(<-chan depositer.TransferDetails)
}

// client defines typed wrapper for the cosmos sdk service client.
type client struct {
	client  *tx.ServiceClient
	storage string
}

// New creates a client that uses the given cosmos sdk service client.
func New(serviceClient *tx.ServiceClient, storage string) Client {
	return &client{
		client:  serviceClient,
		storage: storage,
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
func (c *client) GetBridgeAddress() (*common.Address, error) {
	data, err := ioutil.ReadFile(c.storage)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get the address")
	}

	address := common.BytesToAddress(data)
	return &address, nil
}

// ClaimMinting claims minting from Odin
func (c *client) ClaimMinting(transferDetails <-chan depositer.TransferDetails) {
	for data := range transferDetails {
		fmt.Println(data)
	}
}
