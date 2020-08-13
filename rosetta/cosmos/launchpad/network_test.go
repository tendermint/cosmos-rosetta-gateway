package launchpad

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/require"
)

func TestLaunchpad_NetworkList(t *testing.T) {
	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
	}

	adapter := NewLaunchpad(nil, "", "http://the-url", properties)

	list, err := adapter.NetworkList(context.Background(), nil)
	require.Nil(t, err)

	require.Len(t, list.NetworkIdentifiers, 1)
	require.Equal(t, list.NetworkIdentifiers[0].Network, "TheNetwork")
	require.Equal(t, list.NetworkIdentifiers[0].Blockchain, "TheBlockchain")
}

func TestLaunchpad_NetworkOptions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/node_info", r.URL.Path)
		json.NewEncoder(w).Encode(nodeResponse{
			NodeInfo: nodeInfo{
				Version: "5",
			},
		})
	}))
	defer ts.Close()

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	adapter := NewLaunchpad(http.DefaultClient, "", ts.URL, properties)

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
	tm, err := time.Parse(time.RFC3339, "2019-04-22T17:01:51Z")
	require.NoError(t, err)

	tsCosmos := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/blocks/latest", r.URL.Path)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"block": map[string]interface{}{
				"header": map[string]interface{}{
					"time":   tm.Format(time.RFC3339),
					"height": "16",
					"last_block_id": map[string]interface{}{
						"hash": "ABC",
					},
				},
			},
		})
	}))
	defer tsCosmos.Close()

	tsTendermint := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/net_info":
			json.NewEncoder(w).Encode(map[string]interface{}{
				"result": map[string]interface{}{
					"peers": []map[string]interface{}{
						{
							"node_info": map[string]interface{}{
								"id": "YZ",
							},
						},
					},
				},
			})
		case "/block":
			require.Equal(t, "1", r.URL.Query().Get("height"))
			json.NewEncoder(w).Encode(map[string]interface{}{
				"result": map[string]interface{}{
					"block_id": map[string]interface{}{
						"hash": "DEF",
					},
				},
			})
		}
	}))
	defer tsTendermint.Close()

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	adapter := NewLaunchpad(http.DefaultClient, tsTendermint.URL, tsCosmos.URL, properties)

	status, adapterErr := adapter.NetworkStatus(context.Background(), nil)
	require.Nil(t, adapterErr)
	require.NotNil(t, status)

	require.Equal(t, &types.NetworkStatusResponse{
		CurrentBlockIdentifier: &types.BlockIdentifier{
			Index: 16,
			Hash:  "ABC",
		},
		CurrentBlockTimestamp: tm.UnixNano() / 1000000,
		GenesisBlockIdentifier: &types.BlockIdentifier{
			Hash: "DEF",
		},
		Peers: []*types.Peer{
			{
				PeerID: "YZ",
			},
		},
	}, status)
}
