package client

import (
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/odin/path"
	"io"
	"net/http"
	"net/url"
)

type OdinInterface interface {
	Get(endpoint string) ([]byte, error)
	Put(endpoint string, body io.Reader) ([]byte, error)
	Post(endpoint string, body io.Reader) ([]byte, error)
}

type OdinClient struct {
	client   *http.Client
	resolver path.Resolver
}

func New(client *http.Client, base *url.URL) *OdinClient {
	return &OdinClient{
		client:   client,
		resolver: path.NewResolver(base),
	}
}
