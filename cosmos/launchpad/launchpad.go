package launchpad

import (
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/altsdk"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	cosmos     CosmosAPI
	tendermint TendermintClient

	altCosmos altsdk.Client

	properties rosetta.NetworkProperties
}

func NewLaunchpad(cosmos CosmosAPI, altCosmos altsdk.Client, tendermint TendermintClient, properties rosetta.NetworkProperties) rosetta.Adapter {
	return &Launchpad{
		cosmos:     cosmos,
		altCosmos:  altCosmos,
		tendermint: tendermint,
		properties: properties,
	}
}
