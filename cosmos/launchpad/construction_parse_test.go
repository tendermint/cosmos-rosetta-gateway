package launchpad

import (
	"context"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	"testing"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/require"
)

func TestLaunchpad_ConstructionParse(t *testing.T) {
	properties := properties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
	}
	adapter := newAdapter(sdk.NewClient(""), tendermint.NewClient(""), properties)

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
			rosetta.WrapError(ErrInvalidTransaction, "encoding/hex: invalid byte: U+0069 'i'"),
		},
		// TODO: Update the transaction hash and uncomment
		//{
		//	"valid unsigned tx",
		//	func() string {
		//		return "47282816a90a3ba8a3619a0a1400d87456ee5d45a6c0fbcb677ba0c57f9dc415ba1214f880ae487e47b891ac9b35162bd3c904962afcbb1a090a0461746f6d120131120410c09a0c"
		//	},
		//	nil,
		//	nil,
		//},
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
