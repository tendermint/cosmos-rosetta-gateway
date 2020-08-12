package launchpad

import (
	"net/http"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	properties rosetta.NetworkProperties

	endpoint string
	c        *http.Client
}

func NewLaunchpad(c *http.Client, endpoint string, properties rosetta.NetworkProperties) rosetta.Adapter {
	return &Launchpad{
		properties: properties,
		c:          c,
		endpoint:   endpoint,
	}
}
