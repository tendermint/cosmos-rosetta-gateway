package launchpad

import (
	"net/http"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	blockchain string
	network    string
	operations []string

	endpoint string
	c        *http.Client
}

func NewLaunchpad(c *http.Client, endpoint string, blockchain string, network string,
	operations []string) rosetta.Adapter {
	return &Launchpad{
		blockchain: blockchain,
		network:    network,
		operations: operations,
		c:          c,
		endpoint:   endpoint,
	}
}
