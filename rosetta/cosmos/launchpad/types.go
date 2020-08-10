package launchpad

import (
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

var _ rosetta.Adapter = Launchpad{}

type Launchpad struct {
	endpoint string
}

func NewLaunchpad(endpoint string) *Launchpad {
	return &Launchpad{endpoint: endpoint}
}
