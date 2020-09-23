package launchpad

import (
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	tendermint TendermintClient

	altCosmos     altsdk.Client

	properties rosetta.NetworkProperties
}

func NewLaunchpad(tendermint TendermintAPI, cosmos CosmosAPI, altCosmos altsdk.Client, altTender TendermintClient, properties rosetta.NetworkProperties) rosetta.Adapter {
	return &Launchpad{
		altCosmos:     altCosmos,
		tendermint: tendermint,
		properties:    properties,
	}
}
