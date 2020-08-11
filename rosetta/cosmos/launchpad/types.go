package launchpad

import (
	"net/http"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	endpoint string
	c        *http.Client
}

func NewLaunchpad(c *http.Client, endpoint string) rosetta.Adapter {
	return &Launchpad{
		c:        c,
		endpoint: endpoint,
	}
}
