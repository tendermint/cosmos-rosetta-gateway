package launchpad

import (
	"context"
	"testing"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/require"
)

func TestLaunchpad_ConstructionParse(t *testing.T) {
	properties := Options{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
	}
	adapter := NewLaunchpad(sdk.NewClient(""), tendermint.NewClient(""), properties)

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
