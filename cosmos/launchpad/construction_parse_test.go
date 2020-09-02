package launchpad

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_ConstructionParse(t *testing.T) {
	cases := []struct {
		name string
		req  *types.ConstructionParseRequest
		resp *types.ConstructionParseResponse
		err  *types.Error
	}{
		{"unsigned tx",
			&types.ConstructionParseRequest{
				Transaction: base64.StdEncoding.EncodeToString(jsonMarshal(t, gentx(false))),
			},
			&types.ConstructionParseResponse{
				Operations: []*types.Operation{
					&types.Operation{
						OperationIdentifier: &types.OperationIdentifier{},
						Type:                "Transfer",
						Status:              "Success",
						Account: &types.AccountIdentifier{
							Address: "cosmos12qqzw4tqu32anlcx0a3hupvgdhaf4cc8j9wfyd",
						},
						Amount: &types.Amount{
							Value: "-10",
							Currency: &types.Currency{
								Symbol: "token",
							},
						},
					},
					&types.Operation{
						OperationIdentifier: &types.OperationIdentifier{
							Index: 1,
						},
						Type:   "Transfer",
						Status: "Success",
						Account: &types.AccountIdentifier{
							Address: "cosmos10rpmm9ur87le39hehteha37sg5awdsns6huyvy",
						},
						Amount: &types.Amount{
							Value: "10",
							Currency: &types.Currency{
								Symbol: "token",
							},
						},
					},
				},
				Metadata: map[string]interface{}{
					"memo": "m",
				},
			},
			nil,
		},
		{"signed tx",
			&types.ConstructionParseRequest{
				Transaction: base64.StdEncoding.EncodeToString(jsonMarshal(t, gentx(true))),
			},
			&types.ConstructionParseResponse{
				Operations: []*types.Operation{
					&types.Operation{
						OperationIdentifier: &types.OperationIdentifier{},
						Type:                "Transfer",
						Status:              "Success",
						Account: &types.AccountIdentifier{
							Address: "cosmos12qqzw4tqu32anlcx0a3hupvgdhaf4cc8j9wfyd",
						},
						Amount: &types.Amount{
							Value: "-10",
							Currency: &types.Currency{
								Symbol: "token",
							},
						},
					},
					&types.Operation{
						OperationIdentifier: &types.OperationIdentifier{
							Index: 1,
						},
						Type:   "Transfer",
						Status: "Success",
						Account: &types.AccountIdentifier{
							Address: "cosmos10rpmm9ur87le39hehteha37sg5awdsns6huyvy",
						},
						Amount: &types.Amount{
							Value: "10",
							Currency: &types.Currency{
								Symbol: "token",
							},
						},
					},
				},
				Metadata: map[string]interface{}{
					"memo": "m",
				},
				Signers: []string{
					gentxacc,
				},
			},
			nil,
		},
		{"hex encoded",
			&types.ConstructionParseRequest{
				Transaction: hex.EncodeToString(jsonMarshal(t, gentx(false))),
			},
			&types.ConstructionParseResponse{
				Operations: []*types.Operation{
					&types.Operation{
						OperationIdentifier: &types.OperationIdentifier{},
						Type:                "Transfer",
						Status:              "Success",
						Account: &types.AccountIdentifier{
							Address: "cosmos12qqzw4tqu32anlcx0a3hupvgdhaf4cc8j9wfyd",
						},
						Amount: &types.Amount{
							Value: "-10",
							Currency: &types.Currency{
								Symbol: "token",
							},
						},
					},
					&types.Operation{
						OperationIdentifier: &types.OperationIdentifier{
							Index: 1,
						},
						Type:   "Transfer",
						Status: "Success",
						Account: &types.AccountIdentifier{
							Address: "cosmos10rpmm9ur87le39hehteha37sg5awdsns6huyvy",
						},
						Amount: &types.Amount{
							Value: "10",
							Currency: &types.Currency{
								Symbol: "token",
							},
						},
					},
				},
				Metadata: map[string]interface{}{
					"memo": "m",
				},
			},
			nil,
		},
		{"malformed encoding",
			&types.ConstructionParseRequest{
				Transaction: base64.StdEncoding.EncodeToString(jsonMarshal(t, gentx(false))) + "a",
			},
			nil,
			ErrTxMalformed,
		},
		{"malformed json",
			&types.ConstructionParseRequest{
				Transaction: base64.StdEncoding.EncodeToString(append(jsonMarshal(t, gentx(false)), 0x4)),
			},
			nil,
			ErrTxUnmarshal,
		},
	}

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}
	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{}, properties)

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			parseResp, parseErr := adapter.ConstructionParse(context.Background(), tt.req)
			require.Equal(t, tt.err, parseErr)
			require.Equal(t, tt.resp, parseResp)
		})
	}
}
