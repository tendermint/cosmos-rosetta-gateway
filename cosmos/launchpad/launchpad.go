package launchpad

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Options struct {
	// Blockchain represents the name of the blockchain, it is used for NetworkList endpoint.
	Blockchain string

	// Network represents the name of the network, it is used for NetworkList endpoint.
	Network string

	// AddrPrefix is the prefix used for bech32 addresses.
	AddrPrefix string

	// Offline mode forces to run without querying the node. Some endpoints won't work.
	OfflineMode bool
}

type Launchpad struct {
	tendermint TendermintClient
	cosmos     SdkClient

	options Options
}

func NewLaunchpad(cosmos SdkClient, tendermint TendermintClient, options Options) rosetta.Adapter {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(
		options.AddrPrefix,
		options.AddrPrefix+sdk.PrefixPublic)

	return &Launchpad{
		cosmos:     cosmos,
		tendermint: tendermint,
		options:    options,
	}
}
