package launchpad

import (
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	tendermint TendermintAPI

	altCosmos     SdkClient
	altTendermint TendermintClient

	properties rosetta.NetworkProperties
}

func NewLaunchpad(tendermint TendermintAPI, altCosmos SdkClient, altTender TendermintClient, properties rosetta.NetworkProperties) rosetta.Adapter {
	return &Launchpad{
		tendermint:    tendermint,
		altCosmos:     altCosmos,
		altTendermint: altTender,
		properties:    properties,
	}
}
