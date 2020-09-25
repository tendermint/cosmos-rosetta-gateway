package launchpad

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cosmos "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type properties struct {
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

	properties properties
}

type Options struct {
	// CosmosEndpoint is the endpoint that exposes the cosmos rpc in a cosmos app.
	CosmosEndpoint string

	// CosmosEndpoint is the endpoint that exposes the tendermint rpc in a cosmos app.
	TendermintEndpoint string

	// Blockchain represents the name of the blockchain, it is used for NetworkList endpoint.
	Blockchain string

	// Network represents the name of the network, it is used for NetworkList endpoint.
	Network string

	// AddrPrefix is the prefix used for bech32 addresses.
	AddrPrefix string

	// Offline mode forces to run without querying the node. Some endpoints won't work.
	OfflineMode bool
}

func NewLaunchpad(options Options) rosetta.Adapter {
	cosmosClient := cosmos.NewClient(fmt.Sprintf("http://%s", options.CosmosEndpoint))
	tendermintClient := tendermint.NewClient(fmt.Sprintf("http://%s", options.TendermintEndpoint))

	return newLaunchpad(
		cosmosClient,
		tendermintClient,
		properties{
			Blockchain:  options.Blockchain,
			Network:     options.Network,
			AddrPrefix:  options.AddrPrefix,
			OfflineMode: options.OfflineMode,
		},
	)
}

func newLaunchpad(cosmos SdkClient, tendermint TendermintClient, options properties) rosetta.Adapter {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(
		options.AddrPrefix,
		options.AddrPrefix+sdk.PrefixPublic)

	return &Launchpad{
		cosmos:     cosmos,
		tendermint: tendermint,
		properties: options,
	}
}
