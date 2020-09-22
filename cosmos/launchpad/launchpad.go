package launchpad

import (
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/altsdk"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	cosmos     CosmosAPI
	tendermint TendermintAPI

	altCosmos *altsdk.Client

	properties rosetta.NetworkProperties
}

func NewLaunchpad(tendermint TendermintAPI, cosmos CosmosAPI, altCosmos *altsdk.Client, properties rosetta.NetworkProperties) rosetta.Adapter {
	return &Launchpad{
		tendermint: tendermint,
		cosmos:     cosmos,
		altCosmos:  altCosmos,
		properties: properties,
	}
}
