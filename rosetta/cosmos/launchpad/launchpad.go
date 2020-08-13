package launchpad

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	properties rosetta.NetworkProperties

	cosmosEndpoint     string
	tendermintEndpoint string

	c *http.Client
}

func NewLaunchpad(c *http.Client, tendermintEndpoint, cosmosEndpoint string,
	properties rosetta.NetworkProperties) rosetta.Adapter {
	return &Launchpad{
		properties:         properties,
		c:                  c,
		tendermintEndpoint: tendermintEndpoint,
		cosmosEndpoint:     cosmosEndpoint,
	}
}

func (l Launchpad) cosmos(path string) string {
	return fmt.Sprintf("%s%s", l.cosmosEndpoint, path)
}

func (l Launchpad) tendermint(path string) string {
	return fmt.Sprintf("%s%s", l.tendermintEndpoint, path)
}

func (l Launchpad) request(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	latestBlockReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, body)
	if err != nil {
		return nil, err
	}
	return l.c.Do(latestBlockReq)
}
