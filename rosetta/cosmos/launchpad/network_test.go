package launchpad

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/require"
)

func TestLaunchpad_NetworkList(t *testing.T) {
	blockchain := "TheBlockchain"
	network := "TheNetwork"

	adapter := NewLaunchpad(nil, "http://the-url", blockchain, network)

	list, err := adapter.NetworkList(context.Background(), nil)
	require.Nil(t, err)

	require.Len(t, list.NetworkIdentifiers, 1)
	require.Equal(t, list.NetworkIdentifiers[0].Network, network)
	require.Equal(t, list.NetworkIdentifiers[0].Blockchain, blockchain)
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

	blockchain := "TheBlockchain"
	network := "TheNetwork"

	adapter := NewLaunchpad(http.DefaultClient, ts.URL, blockchain, network)

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
			OperationTypes: []string{
				"Transfer",
				"Reward",
			},
		},
	}, options)
}
