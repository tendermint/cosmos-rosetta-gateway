package launchpad

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	client "github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad/client/generated"
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
	BankBalancesAddressGet(ctx context.Context, address string) (client.InlineResponse2004, *http.Response, error)
}

type TendermintAPI interface {
	NodeInfoGet(ctx context.Context) (client.InlineResponse200, *http.Response, error)
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
