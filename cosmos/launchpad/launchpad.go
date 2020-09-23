package launchpad

import (
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	tendermint TendermintClient

	cosmos SdkClient

	properties rosetta.NetworkProperties
}

func NewLaunchpad(cosmos SdkClient, tendermint TendermintClient, properties rosetta.NetworkProperties) rosetta.Adapter {
	return &Launchpad{
		cosmos:     cosmos,
		tendermint: tendermint,
		properties: properties,
	}
}
