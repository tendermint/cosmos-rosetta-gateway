package launchpad

import (
	"context"
	"testing"
	"time"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/alttendermint/mocks"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/alttendermint"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/altsdk"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	cosmosmocks "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/mocks"
	tendermintmocks "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/mocks"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_NetworkList(t *testing.T) {
	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
	}

	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{}, altsdk.NewClient(""), alttendermint.NewClient(""), properties)

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

	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{Tendermint: m}, altsdk.NewClient(""), alttendermint.NewClient(""), properties)

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
					Status:     StatusSuccess,
					Successful: true,
				},
				{
					Status:     StatusReverted,
					Successful: false,
				},
			},
			OperationTypes: []string{
				OperationTransfer,
			},
		},
	}, options)
}

func TestLaunchpad_NetworkStatus(t *testing.T) {
	m := &tendermintmocks.TendermintInfoAPI{}
	defer m.AssertExpectations(t)

	mt := &mocks.TendermintClient{}

	ti, err := time.Parse(time.RFC3339, "2019-04-22T17:01:51Z")
	require.NoError(t, err)

	mt.
		On("Block", uint64(0)).
		Return(alttendermint.BlockResponse{
			Block: alttendermint.Block{
				Header: alttendermint.BlockHeader{
					Time:   ti.Format(time.RFC3339),
					Height: "2",
				},
			},
			BlockId: alttendermint.BlockId{
				Hash: "3",
			},
		}, nil, nil).
		Once()

	mt.
		On("Block", uint64(1)).
		Return(alttendermint.BlockResponse{
			Block: alttendermint.Block{
				Header: alttendermint.BlockHeader{
					Height: "1",
				},
			},
			BlockId: alttendermint.BlockId{
				Hash: "4",
			},
		}, nil, nil).
		Once()

	mt.
		On("NetInfo", mock.Anything).
		Return(alttendermint.NetInfoResponse{
			Peers: []alttendermint.Peer{
				{
					NodeInfo: alttendermint.NodeInfo{
						Id: "1",
					},
				},
				{
					NodeInfo: alttendermint.NodeInfo{
						Id: "2",
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

	adapter := NewLaunchpad(
		TendermintAPI{Info: m},
		CosmosAPI{},
		altsdk.NewClient(""),
		mt,
		properties,
	)

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
