package launchpad

import (
	"context"
	"fmt"
	"net/http"

	cosmoslp "github.com/tendermint/cosmos-rosetta-gateway/generated/cosmos-launchpad"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	api                API
	tendermintEndpoint string
	c                  *http.Client

	properties rosetta.NetworkProperties
}

type API struct {
	Bank       BankAPI
	Tendermint TendermintAPI
}

type BankAPI interface {
	BankBalancesAddressGet(ctx context.Context, address string) (cosmoslp.InlineResponse2004, *http.Response, error)
}

type TendermintAPI interface {
	NodeInfoGet(ctx context.Context) (cosmoslp.InlineResponse200, *http.Response, error)
}

func NewLaunchpad(c *http.Client, api API, tendermintEndpoint string, properties rosetta.NetworkProperties) rosetta.Adapter {
	return &Launchpad{
		api:                api,
		c:                  c,
		tendermintEndpoint: tendermintEndpoint,
		properties:         properties,
	}
}

func (l Launchpad) tendermint(path string) string {
	return fmt.Sprintf("%s%s", l.tendermintEndpoint, path)
}
