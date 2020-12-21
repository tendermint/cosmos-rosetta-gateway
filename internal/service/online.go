package service

import (
	"context"
	"time"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/tendermint/cosmos-rosetta-gateway/errors"
	crgtypes "github.com/tendermint/cosmos-rosetta-gateway/types"
)

// genesisBlockFetchTimeout defines a timeout to fetch the genesis block
const genesisBlockFetchTimeout = 15 * time.Second

// NewOnlineNetwork builds a single network adapter.
// It will get the Genesis block on the beginning to avoid calling it everytime.
func NewOnlineNetwork(network *types.NetworkIdentifier, client crgtypes.Client) (crgtypes.API, error) {
	ctx, cancel := context.WithTimeout(context.Background(), genesisBlockFetchTimeout)
	defer cancel()

	var genesisHeight int64 = 1
	block, err := client.BlockByHeight(ctx, &genesisHeight)
	if err != nil {
		return OnlineNetwork{}, err
	}

	return OnlineNetwork{
		client:  client,
		network: network,
		networkOptions: &types.NetworkOptionsResponse{
			Version: &types.Version{
				RosettaVersion: crgtypes.SpecVersion,
				NodeVersion:    client.Version(),
			},
			Allow: &types.Allow{
				OperationStatuses:       client.OperationStatuses(),
				OperationTypes:          client.SupportedOperations(),
				Errors:                  errors.SealAndListErrors(),
				HistoricalBalanceLookup: true,
			},
		},
		genesisBlockIdentifier: block.Block,
	}, nil
}

// OnlineNetwork groups together all the components required for the full rosetta implementation
type OnlineNetwork struct {
	client crgtypes.Client // used to query cosmos app + tendermint

	network        *types.NetworkIdentifier      // identifies the network, it's static
	networkOptions *types.NetworkOptionsResponse // identifies the network options, it's static

	genesisBlockIdentifier *types.BlockIdentifier // identifies genesis block, it's static
}
