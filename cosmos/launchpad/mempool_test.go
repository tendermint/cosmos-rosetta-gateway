package launchpad

import (
	"context"
	"testing"

	mocks3 "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/mocks"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_Mempool(t *testing.T) {
	m := &mocks3.TendermintClient{}
	defer m.AssertExpectations(t)

	m.
		On("UnconfirmedTxs").
		Return(tendermint.UnconfirmedTxsResponse{
			Txs: []string{
				"1QEoKBapCl0l5qD4CiRkNGFiMDdlYi1jZGUxLTRjZmQtOWI3OS04MzYzNjFmN2RjNTcSFKeCHRQzgA2HavcLTcf4xdScUjrtGghtYW5vbGV0ZSIRdXNlckBtYW5vbGV0ZS5jb20SBBDAmgwaagom61rphyECU9fDYFDAP5TWDimv6z0BdK6oyV\nzv3iCb9fUWAAb4AoYSQCbvAfmO+aqF5WZ1M67XLZbV7OI3Sq8sbnV58tx5gf3nW/C/89pTTNmWmBskrOzmbmNEmBPQl1biuXAsUCwyMfE=",
			},
		}, nil, nil)

	adapter := NewLaunchpad(sdk.NewClient(""), m, rosetta.NetworkProperties{})

	mempool, err := adapter.Mempool(context.Background(), &types.NetworkRequest{})
	require.Nil(t, err)

	require.Equal(t, &types.MempoolResponse{TransactionIdentifiers: []*types.TransactionIdentifier{
		{
			Hash: "99b044765216517005cf096e26111016be457454ca7f83d5498d4b1142c89631",
		},
	}}, mempool)
}

func TestLaunchpad_Mempool_FailsOfflineMode(t *testing.T) {
	properties := rosetta.NetworkProperties{
		OfflineMode: true,
	}
	adapter := NewLaunchpad(CosmosAPI{}, altsdk.NewClient(""), tendermint.NewClient(""), properties)

	_, err := adapter.Mempool(context.Background(), &types.NetworkRequest{})
	require.Equal(t, ErrEndpointDisabledOfflineMode, err)
}

func TestLaunchpad_MempoolTransaction(t *testing.T) {
	ma := &mocks3.TendermintClient{}
	defer ma.AssertExpectations(t)

	ma.
		On("Tx", "ABCTHEHASH").
		Return(tendermint.TxResponse{
			Hash: "ABCTHEHASH",
		},
			nil, nil)

	adapter := NewLaunchpad(sdk.NewClient(""), ma, rosetta.NetworkProperties{})
	res, err := adapter.MempoolTransaction(context.Background(), &types.MempoolTransactionRequest{
		TransactionIdentifier: &types.TransactionIdentifier{Hash: "ABCTHEHASH"},
	})

	require.Nil(t, err)

	require.Equal(t, "ABCTHEHASH", res.Transaction.TransactionIdentifier.Hash)
}

func TestLaunchpad_MempoolTransaction_FailsOfflineMode(t *testing.T) {
	properties := rosetta.NetworkProperties{
		OfflineMode: true,
	}
	adapter := NewLaunchpad(CosmosAPI{}, altsdk.NewClient(""), tendermint.NewClient(""), properties)
	_, err := adapter.MempoolTransaction(context.Background(), &types.MempoolTransactionRequest{
		TransactionIdentifier: &types.TransactionIdentifier{Hash: "ABCTHEHASH"},
	})

	require.Equal(t, ErrEndpointDisabledOfflineMode, err)
}
