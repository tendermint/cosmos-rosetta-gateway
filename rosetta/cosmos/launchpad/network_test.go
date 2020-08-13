package launchpad

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	tsTendermint := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/net_info":
			w.Write([]byte(`{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "listening": true,
    "listeners": [
      "Listener(@)"
    ],
    "n_peers": "0",
    "peers": []
  }
}`))
		case "/block":
			callingGenesis := r.URL.Query().Get("height") == "1"
			if callingGenesis {
				w.Write([]byte(`{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "360A1DED0DEE79A8A28FBD88517EA3B6A9719460A9BE30D8E8D786D5AD79127B",
      "parts": {
        "total": "1",
        "hash": "82914D192B2C538716049AA0193DDADD8F855AAA04B596587E0B2BE4CEF27E5E"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "10",
          "app": "0"
        },
        "chain_id": "blog",
        "height": "1",
        "time": "2020-08-13T11:36:18.162487Z",
        "last_block_id": {
          "hash": "",
          "parts": {
            "total": "0",
            "hash": ""
          }
        },
        "last_commit_hash": "",
        "data_hash": "",
        "validators_hash": "6260C775FAC0092DFD7574B4943F6A180F409189BA24BF9B9A3A4C74CA512D47",
        "next_validators_hash": "6260C775FAC0092DFD7574B4943F6A180F409189BA24BF9B9A3A4C74CA512D47",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "",
        "last_results_hash": "",
        "evidence_hash": "",
        "proposer_address": "89E00B1FA0E5DACAAB55AD38C422ADB433936C69"
      },
      "data": {
        "txs": null
      },
      "evidence": {
        "evidence": null
      },
      "last_commit": {
        "height": "0",
        "round": "0",
        "block_id": {
          "hash": "",
          "parts": {
            "total": "0",
            "hash": ""
          }
        },
        "signatures": null
      }
    }
  }
}`))
			} else {
				w.Write([]byte(`{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "8FEB56E18A7B5FE53C42EEB43CD0113D24BB1B2DCEA4747004887A1464E5826C",
      "parts": {
        "total": "1",
        "hash": "0AB661CF6539F5CDAE6FD6DFE9F9B6AB87126578BF5D39CD5987888451938217"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "10",
          "app": "0"
        },
        "chain_id": "blog",
        "height": "1230",
        "time": "2020-08-13T13:32:57.228899Z",
        "last_block_id": {
          "hash": "8C11129024646574E9A6E6B861C45ABAC7AE5684EA187621AEBF14B93DD44F2D",
          "parts": {
            "total": "1",
            "hash": "CB57B80F10351B55AA76F11343E8B65F8E5CCDEF5C9C3218B4CFF01616F8C6F4"
          }
        },
        "last_commit_hash": "DE38C291699FECB9AAECEF5F083B2ED090CA0B98BB9F883E1FCD479765F73AFF",
        "data_hash": "",
        "validators_hash": "6260C775FAC0092DFD7574B4943F6A180F409189BA24BF9B9A3A4C74CA512D47",
        "next_validators_hash": "6260C775FAC0092DFD7574B4943F6A180F409189BA24BF9B9A3A4C74CA512D47",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "E906D1B22F83CB7C7B8E838D4C4B96F114780F6A09C341ED508AED87CF7C367F",
        "last_results_hash": "",
        "evidence_hash": "",
        "proposer_address": "89E00B1FA0E5DACAAB55AD38C422ADB433936C69"
      },
      "data": {
        "txs": null
      },
      "evidence": {
        "evidence": null
      },
      "last_commit": {
        "height": "1229",
        "round": "0",
        "block_id": {
          "hash": "8C11129024646574E9A6E6B861C45ABAC7AE5684EA187621AEBF14B93DD44F2D",
          "parts": {
            "total": "1",
            "hash": "CB57B80F10351B55AA76F11343E8B65F8E5CCDEF5C9C3218B4CFF01616F8C6F4"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "89E00B1FA0E5DACAAB55AD38C422ADB433936C69",
            "timestamp": "2020-08-13T13:32:57.228899Z",
            "signature": "A1TjQYHitfWrQDz9+Xvj8aJymj+HSPSHZOJHblF8dJBoUTNnrHRqVAJsMH0LWBItyn7JJdCGd9cci4VfwBnNBQ=="
          }
        ]
      }
    }
  }
}`))
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
