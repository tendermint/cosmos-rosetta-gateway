package launchpad

import (
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	cosmos     CosmosAPI
	tendermint TendermintAPI

	altCosmos     SdkClient
	altTendermint TendermintClient

	properties rosetta.NetworkProperties
}

func NewLaunchpad(tendermint TendermintAPI, cosmos CosmosAPI, altCosmos SdkClient, altTender TendermintClient, properties rosetta.NetworkProperties) rosetta.Adapter {
	return &Launchpad{
		tendermint:    tendermint,
		cosmos:        cosmos,
		altCosmos:     altCosmos,
		altTendermint: altTender,
		properties:    properties,
	}
}
