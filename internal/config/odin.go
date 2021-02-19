package config

import (
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/odin/client"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"net/url"
)

type odiner struct {
	getter kv.Getter
	once   comfig.Once
	value  *client.OdinClient
}

type Odiner interface {
	OdinClient() *client.OdinClient
}

func NewOdiner(getter kv.Getter) Odiner {
	return &odiner{
		getter: getter,
	}
}

func (o *odiner) OdinClient() *client.OdinClient {
	o.once.Do(func() interface{} {
		var config struct {
			Endpoint *url.URL `fig:"endpoint,required"`
		}

		if err := figure.Out(&config).From(kv.MustGetStringMap(o.getter, "odin")).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out odin client"))
		}

		o.value = client.New(http.DefaultClient, config.Endpoint)
		return nil
	})

	return o.value
}
