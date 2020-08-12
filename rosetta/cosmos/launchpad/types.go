package launchpad

import (
	"net/http"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	blockchain string
	network    string

	endpoint string
	c        *http.Client
}

func NewLaunchpad(c *http.Client, endpoint string, blockchain string, network string) rosetta.Adapter {
	return &Launchpad{
		blockchain: blockchain,
		network:    network,
		c:          c,
		endpoint:   endpoint,
	}
}
