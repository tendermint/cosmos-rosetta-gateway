package launchpad

import (
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	cosmos     CosmosAPI
	tendermint TendermintAPI

	properties rosetta.NetworkProperties
}

func NewLaunchpad(tendermint TendermintAPI, cosmos CosmosAPI, properties rosetta.NetworkProperties) rosetta.Adapter {
	return &Launchpad{
		tendermint: tendermint,
		cosmos:     cosmos,
		properties: properties,
	}
}
