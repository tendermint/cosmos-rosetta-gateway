package launchpad

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Launchpad struct {
	tendermint TendermintClient

	cosmos SdkClient

	properties rosetta.NetworkProperties
}

func NewLaunchpad(cosmos SdkClient, tendermint TendermintClient, properties rosetta.NetworkProperties) rosetta.Adapter {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(
		properties.AddrPrefix,
		properties.AddrPrefix+sdk.PrefixPublic)

	return &Launchpad{
		cosmos:     cosmos,
		tendermint: tendermint,
		properties: properties,
	}
}
