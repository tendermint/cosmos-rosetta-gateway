package launchpad

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
)

func gentx(signed bool) cosmosclient.StdTx {
	tx := cosmosclient.StdTx{
		Value: cosmosclient.StdTxValue{
			Msg: []cosmosclient.Msg{
				{
					Type: "cosmos-sdk/MsgSend",
					Value: cosmosclient.MsgValue{
						FromAddress: "cosmos12qqzw4tqu32anlcx0a3hupvgdhaf4cc8j9wfyd",
						ToAddress:   "cosmos10rpmm9ur87le39hehteha37sg5awdsns6huyvy",
						Amount: []cosmosclient.Coin{
							{
								Denom:  "token",
								Amount: "10",
							},
						},
					},
				},
			},
			Fee:        cosmosclient.StdTxValueFee{},
			Signatures: []cosmosclient.StdTxValueSignatures{},
			Memo:       "m",
		},
	}
	if signed {
		tx.Value.Signatures = append(tx.Value.Signatures, cosmosclient.StdTxValueSignatures{
			PubKey: cosmosclient.StdTxValuePubKey{
				Value: "A00/bnqYY3LTTI6MYAg0P4SLSeTWgklpH3ERqFkCsbC2",
			},
		})
	}
	return tx
}

func jsonMarshal(t *testing.T, data interface{}) []byte {
	b, err := json.Marshal(data)
	require.NoError(t, err)
	return b
}
