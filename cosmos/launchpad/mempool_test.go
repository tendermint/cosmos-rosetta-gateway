package launchpad

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coinbase/rosetta-sdk-go/types"

	"github.com/stretchr/testify/mock"

	tendermintclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/mocks"
)

func TestLaunchpad_Mempool(t *testing.T) {
	m := &mocks.TendermintInfoAPI{}
	defer m.AssertExpectations(t)

	m.
		On("UnconfirmedTxs", mock.Anything, mock.Anything).
		Return(tendermintclient.UnconfirmedTransactionsResponse{
			Result: tendermintclient.UnconfirmedTransactionsResponseResult{
				NTxs:       "1",
				Total:      "1",
				TotalBytes: "1",
				Txs: []string{
					"1QEoKBapCl0l5qD4CiRkNGFiMDdlYi1jZGUxLTRjZmQtOWI3OS04MzYzNjFmN2RjNTcSFKeCHRQzgA2HavcLTcf4xdScUjrtGghtYW5vbGV0ZSIRdXNlckBtYW5vbGV0ZS5jb20SBBDAmgwaagom61rphyECU9fDYFDAP5TWDimv6z0BdK6oyV\nzv3iCb9fUWAAb4AoYSQCbvAfmO+aqF5WZ1M67XLZbV7OI3Sq8sbnV58tx5gf3nW/C/89pTTNmWmBskrOzmbmNEmBPQl1biuXAsUCwyMfE=",
				},
			},
		}, nil, nil)

	adapter := NewLaunchpad(TendermintAPI{Info: m}, CosmosAPI{}, rosetta.NetworkProperties{})

	mempool, err := adapter.Mempool(context.Background(), &types.NetworkRequest{})
	require.Nil(t, err)

	require.Equal(t, &types.MempoolResponse{TransactionIdentifiers: []*types.TransactionIdentifier{
		{
			Hash: "99b044765216517005cf096e26111016be457454ca7f83d5498d4b1142c89631",
		},
	}}, mempool)
}
