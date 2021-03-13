package config

import (
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/odin/client"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"google.golang.org/grpc"
)

// Odin defines an interface for wrapped odin client.
type Odin interface {
	OdinClient() *client.Client
}

// odin defines typed wrapper for the odin client.
type odin struct {
	getter kv.Getter
	once   comfig.Once
	client *client.Client
}

// NewOdin creates a new odin client.
func NewOdin(getter kv.Getter) Odin {
	return &odin{getter: getter}
}

// OdinClient returns odin client.
func (o *odin) OdinClient() *client.Client {
	o.once.Do(func() interface{} {
		var config struct {
			Endpoint string `fig:"endpoint,required"`
			Storage  string `fig:"storage,required"`
		}

		if err := figure.Out(&config).From(kv.MustGetStringMap(o.getter, "odin")).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out rpc"))
		}

		clientConn, err := grpc.Dial(config.Endpoint, grpc.WithInsecure())
		if err != nil {
			return errors.Wrap(err, "failed to dial")
		}

		serviceClient := tx.NewServiceClient(clientConn)
		odinClient := client.New(&serviceClient, config.Storage)
		o.client = odinClient

		return nil
	})

	return o.client
}
