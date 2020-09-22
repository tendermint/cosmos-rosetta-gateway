package launchpad

import (
	"context"
	"testing"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/alttendermint"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/altsdk"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_ConstructionParse(t *testing.T) {
	var (
		properties = rosetta.NetworkProperties{
			Blockchain: "TheBlockchain",
			Network:    "TheNetwork",
			SupportedOperations: []string{
				"Transfer",
				"Reward",
			},
		}
		adapter = NewLaunchpad(TendermintAPI{}, CosmosAPI{}, altsdk.NewClient(""), alttendermint.NewClient(""), properties)
	)

	cases := []struct {
		name  string
		getTx func() string
		resp  *types.ConstructionParseResponse
		err   *types.Error
	}{
		// TODO: Add a test for unsigned tx
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
