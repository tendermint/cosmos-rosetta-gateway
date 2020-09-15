package launchpad

import (
	"context"
	"testing"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/altsdk"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_ConstructionParse(t *testing.T) {
	var (
		operations = []*types.Operation{
			{
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
			{
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
		}

		properties = rosetta.NetworkProperties{
			Blockchain: "TheBlockchain",
			Network:    "TheNetwork",
			SupportedOperations: []string{
				"Transfer",
				"Reward",
			},
		}
		adapter = NewLaunchpad(TendermintAPI{}, CosmosAPI{}, altsdk.NewClient(""), properties)
	)

	cases := []struct {
		name  string
		getTx func() string
		resp  *types.ConstructionParseResponse
		err   *types.Error
	}{
		{"unsigned tx",
			func() string {
				payloadsResp, payloadsErr := adapter.ConstructionPayloads(context.Background(), &types.ConstructionPayloadsRequest{
					Operations: operations,
					Metadata: map[string]interface{}{
						ChainIdKey:       "ck",
						SequenceKey:      float64(1),
						AccountNumberKey: float64(2),
					},
				})
				require.Nil(t, payloadsErr)
				return payloadsResp.UnsignedTransaction
			},
			&types.ConstructionParseResponse{
				Operations: operations,
				Metadata: map[string]interface{}{
					Memo:             "TODO memo",
					ChainIdKey:       "ck",
					SequenceKey:      uint64(1),
					AccountNumberKey: uint64(2),
				},
			},
			nil,
		},
		{"invalid tx",
			func() string {
				return "invalid"
			},
			nil,
			ErrInvalidTransaction,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			req := &types.ConstructionParseRequest{
				Transaction: tt.getTx(),
			}
			parseResp, parseErr := adapter.ConstructionParse(context.Background(), req)
			require.Equal(t, tt.err, parseErr)
			require.Equal(t, tt.resp, parseResp)
		})
	}
}
