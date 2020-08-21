package launchpad

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/antihax/optional"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	cosmosmocks "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/mocks"
	tendermintclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
	tendermintmocks "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/mocks"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_NetworkList(t *testing.T) {
	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
	}

	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{}, properties)

	list, err := adapter.NetworkList(context.Background(), nil)
	require.Nil(t, err)

	require.Len(t, list.NetworkIdentifiers, 1)
	require.Equal(t, list.NetworkIdentifiers[0].Network, "TheNetwork")
	require.Equal(t, list.NetworkIdentifiers[0].Blockchain, "TheBlockchain")
}

func TestLaunchpad_NetworkOptions(t *testing.T) {
	m := &cosmosmocks.CosmosTendermintAPI{}
	defer m.AssertExpectations(t)

	m.
		On("NodeInfoGet", mock.Anything).
		Return(cosmosclient.InlineResponse200{
			NodeInfo: cosmosclient.InlineResponse200NodeInfo{
				Version: "5",
			},
		}, nil, nil).
		Once()

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{Tendermint: m}, properties)

	options, err := adapter.NetworkOptions(context.Background(), nil)
	require.Nil(t, err)
	require.NotNil(t, options)

	require.Equal(t, &types.NetworkOptionsResponse{
		Version: &types.Version{
			RosettaVersion: "1.2.5",
			NodeVersion:    "5",
		},
		Allow: &types.Allow{
			OperationStatuses: []*types.OperationStatus{
				{
					Status:     "SUCCESS",
					Successful: true,
				},
			},
			OperationTypes: properties.SupportedOperations,
		},
	}, options)
}

func TestLaunchpad_NetworkStatus(t *testing.T) {
	m := &tendermintmocks.TendermintInfoAPI{}
	defer m.AssertExpectations(t)

	var blockOpts *tendermintclient.BlockOpts
	ti, err := time.Parse(time.RFC3339, "2019-04-22T17:01:51Z")
	require.NoError(t, err)

	m.
		On("Block", mock.Anything, blockOpts).
		Return(tendermintclient.BlockResponse{
			Result: tendermintclient.BlockComplete{
				Block: tendermintclient.Block{
					Header: tendermintclient.BlockHeader{
						Time:   ti.Format(time.RFC3339),
						Height: "2",
					},
				},
				BlockId: tendermintclient.BlockId{
					Hash: "3",
				},
			},
		}, nil, nil).
		Once()

	m.
		On("Block", mock.Anything, &tendermintclient.BlockOpts{
			Height: optional.NewFloat32(1),
		}).
		Return(tendermintclient.BlockResponse{
			Result: tendermintclient.BlockComplete{
				Block: tendermintclient.Block{
					Header: tendermintclient.BlockHeader{
						Height: "1",
					},
				},
				BlockId: tendermintclient.BlockId{
					Hash: "4",
				},
			},
		}, nil, nil).
		Once()

	m.
		On("NetInfo", mock.Anything).
		Return(tendermintclient.NetInfoResponse{
			Result: tendermintclient.NetInfo{
				Peers: []tendermintclient.Peer{
					{
						NodeInfo: tendermintclient.NodeInfo{
							Id: "1",
						},
					},
					{
						NodeInfo: tendermintclient.NodeInfo{
							Id: "2",
						},
					},
				},
			},
		}, nil, nil).
		Once()

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	adapter := NewLaunchpad(TendermintAPI{Info: m}, CosmosAPI{}, properties)

	status, adapterErr := adapter.NetworkStatus(context.Background(), nil)
	require.Nil(t, adapterErr)
	require.NotNil(t, status)

	require.Equal(t, &types.NetworkStatusResponse{
		CurrentBlockIdentifier: &types.BlockIdentifier{
			Index: 2,
			Hash:  "3",
		},
		CurrentBlockTimestamp: ti.UnixNano() / 1000000,
		GenesisBlockIdentifier: &types.BlockIdentifier{
			Index: 1,
			Hash:  "4",
		},
		Peers: []*types.Peer{
			{
				PeerID: "1",
			},
			{
				PeerID: "2",
			},
		},
	}, status)
}

func getContentsFromFile(t *testing.T, filename string) []byte {
	file, err := os.Open(filename)
	require.NoError(t, err)

	genesisContent, err := ioutil.ReadAll(file)
	return genesisContent
}
