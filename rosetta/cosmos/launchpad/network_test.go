package launchpad

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

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
		err := json.NewEncoder(w).Encode(nodeResponse{
			NodeInfo: nodeInfo{
				Version: "5",
			},
		})
		require.NoError(t, err)
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
	tsTendermint := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/net_info":
			peersContent := getContentsFromFile(t, "testdata/peers.json")
			_, err := w.Write(peersContent)
			require.NoError(t, err)
		case "/block":
			callingGenesis := r.URL.Query().Get("height") == "1"
			if callingGenesis {
				genesisContent := getContentsFromFile(t, "testdata/genesis-block.json")
				_, err := w.Write(genesisContent)
				require.NoError(t, err)
			} else {
				latestContent := getContentsFromFile(t, "testdata/latest-block.json")
				_, err := w.Write(latestContent)
				require.NoError(t, err)
			}
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

	adapter := NewLaunchpad(http.DefaultClient, tsTendermint.URL, "", properties)

	status, adapterErr := adapter.NetworkStatus(context.Background(), nil)
	require.Nil(t, adapterErr)
	require.NotNil(t, status)

	require.Equal(t, &types.NetworkStatusResponse{
		CurrentBlockIdentifier: &types.BlockIdentifier{
			Index: 1230,
			Hash:  "8FEB56E18A7B5FE53C42EEB43CD0113D24BB1B2DCEA4747004887A1464E5826C",
		},
		CurrentBlockTimestamp: 1597325577228,
		GenesisBlockIdentifier: &types.BlockIdentifier{
			Hash:  "360A1DED0DEE79A8A28FBD88517EA3B6A9719460A9BE30D8E8D786D5AD79127B",
			Index: 1,
		},
		Peers: nil,
	}, status)
}

func getContentsFromFile(t *testing.T, filename string) []byte {
	file, err := os.Open(filename)
	require.NoError(t, err)

	genesisContent, err := ioutil.ReadAll(file)
	return genesisContent
}
