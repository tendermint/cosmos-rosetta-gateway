package launchpad

import (
	"net/http"

	http2 "github.com/tendermint/cosmos-rosetta-gateway/http"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	properties http2.Properties

	endpoint string
	c        *http.Client
}

func NewLaunchpad(c *http.Client, endpoint string, properties http2.Properties) rosetta.Adapter {
	return &Launchpad{
		properties: properties,
		c:          c,
		endpoint:   endpoint,
	}
}
