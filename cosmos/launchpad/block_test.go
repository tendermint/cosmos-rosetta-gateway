package launchpad

import (
	"context"
	"testing"
	"time"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/alttendermint/mocks"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/alttendermint"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/altsdk"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	cosmosmocks "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/mocks"
	tendermintclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
	tendermintmocks "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/mocks"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_Block(t *testing.T) {
	var (
		mt = &tendermintmocks.TendermintInfoAPI{}
		mc = &cosmosmocks.CosmosTransactionsAPI{}
		ma = &mocks.TendermintClient{}
	)
	defer mt.AssertExpectations(t)
	defer mc.AssertExpectations(t)

	ti, err := time.Parse(time.RFC3339, "2019-04-22T17:01:51Z")
	require.NoError(t, err)

	ma.
		On("Block", uint64(1)).
		Return(alttendermint.BlockResponse{
			BlockId: alttendermint.BlockId{
				Hash: "11",
			},
			Block: alttendermint.Block{
				Header: alttendermint.BlockHeader{
					Height: "2",
					Time:   ti.Format(time.RFC3339),
					LastBlockId: alttendermint.BlockId{
						Hash: "12",
					},
				},
			},
		}, nil, nil).
		Once()

	var opts *tendermintclient.TxSearchOpts
	mt.
		On("TxSearch", mock.Anything, `"tx.height=2"`, opts).
		Return(tendermintclient.TxSearchResponse{
			Result: tendermintclient.TxSearchResponseResult{
				Txs: []tendermintclient.TxSearchResponseResultTxs{
					{
						Hash: "3",
					},
					{
						Hash: "4",
					},
				},
			},
		}, nil, nil).
		Once()

	mc.
		On("TxsHashGet", mock.Anything, "3").
		Return(cosmosclient.TxQuery{
			Txhash: "3",
			Tx: cosmosclient.StdTx{
				Value: cosmosclient.StdTxValue{
					Msg: []cosmosclient.Msg{
						{
							Type: "cosmos-sdk/MsgSend",
							Value: cosmosclient.MsgValue{
								FromAddress: "6",
								ToAddress:   "15",
								Amount: []cosmosclient.Coin{
									{
										Amount: "13",
										Denom:  "14",
									},
								},
							},
						},
					},
				},
			},
		}, nil, nil).
		Once()

	mc.
		On("TxsHashGet", mock.Anything, "4").
		Return(cosmosclient.TxQuery{
			Txhash: "4",
			Tx: cosmosclient.StdTx{
				Value: cosmosclient.StdTxValue{
					Msg: []cosmosclient.Msg{
						{
							Type: "cosmos-sdk/MsgSend",
							Value: cosmosclient.MsgValue{
								FromAddress: "8",
								ToAddress:   "16",
								Amount: []cosmosclient.Coin{
									{
										Amount: "9",
										Denom:  "10",
									},
								},
							},
						},
					},
				},
			},
		}, nil, nil).
		Once()

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			OperationTransfer,
		},
	}

	adapter := NewLaunchpad(TendermintAPI{Info: mt}, CosmosAPI{Transactions: mc}, altsdk.NewClient(""), ma, properties)

	var h int64 = 1
	block, blockErr := adapter.Block(context.Background(), &types.BlockRequest{
		BlockIdentifier: &types.PartialBlockIdentifier{
			Index: &h,
		},
	})
	require.Nil(t, blockErr)
	require.NotNil(t, block)

	require.NotNil(t, block.Block)

	// a small hack to get the transactions in the same order.
	// without the hack, we need to check BlockResponse.Block props one by one, which is longer.
	require.Len(t, block.Block.Transactions, 2)
	require.NotNil(t, block.Block.Transactions[0].TransactionIdentifier)

	if block.Block.Transactions[0].TransactionIdentifier.Hash == "3" {
		block.Block.Transactions[0],
			block.Block.Transactions[1] =
			block.Block.Transactions[1],
			block.Block.Transactions[0]
	}

	// compare the full response.
	require.Equal(t, &types.BlockResponse{
		Block: &types.Block{
			BlockIdentifier: &types.BlockIdentifier{
				Index: int64(2),
				Hash:  "11",
			},
			Transactions: []*types.Transaction{
				{
					TransactionIdentifier: &types.TransactionIdentifier{
						Hash: "4",
					},
					Operations: []*types.Operation{
						{
							OperationIdentifier: &types.OperationIdentifier{},
							Type:                OperationTransfer,
							Status:              StatusSuccess,
							Account: &types.AccountIdentifier{
								Address: "8",
							},
							Amount: &types.Amount{
								Value: "-9",
								Currency: &types.Currency{
									Symbol: "10",
								},
							},
						},
						{
							OperationIdentifier: &types.OperationIdentifier{
								Index: 1,
							},
							Type:   OperationTransfer,
							Status: StatusSuccess,
							Account: &types.AccountIdentifier{
								Address: "16",
							},
							Amount: &types.Amount{
								Value: "9",
								Currency: &types.Currency{
									Symbol: "10",
								},
							},
						},
					},
				},
				{
					TransactionIdentifier: &types.TransactionIdentifier{
						Hash: "3",
					},
					Operations: []*types.Operation{
						{
							OperationIdentifier: &types.OperationIdentifier{},
							Type:                OperationTransfer,
							Status:              StatusSuccess,
							Account: &types.AccountIdentifier{
								Address: "6",
							},
							Amount: &types.Amount{
								Value: "-13",
								Currency: &types.Currency{
									Symbol: "14",
								},
							},
						},
						{
							OperationIdentifier: &types.OperationIdentifier{
								Index: 1,
							},
							Type:   OperationTransfer,
							Status: StatusSuccess,
							Account: &types.AccountIdentifier{
								Address: "15",
							},
							Amount: &types.Amount{
								Value: "13",
								Currency: &types.Currency{
									Symbol: "14",
								},
							},
						},
					},
				},
			},
			Timestamp: ti.UnixNano() / 1000000,
			ParentBlockIdentifier: &types.BlockIdentifier{
				Index: int64(1),
				Hash:  "12",
			},
		},
	}, block)

}

func TestLaunchpad_BlockTransaction(t *testing.T) {
	mc := &cosmosmocks.CosmosTransactionsAPI{}
	defer mc.AssertExpectations(t)

	mc.
		On("TxsHashGet", mock.Anything, "1").
		Return(cosmosclient.TxQuery{
			Txhash: "1",
			Tx: cosmosclient.StdTx{
				Value: cosmosclient.StdTxValue{
					Msg: []cosmosclient.Msg{
						{
							Type: "cosmos-sdk/MsgSend",
							Value: cosmosclient.MsgValue{
								FromAddress: "3",
								ToAddress:   "6",
								Amount: []cosmosclient.Coin{
									{
										Amount: "4",
										Denom:  "5",
									},
								},
							},
						},
					},
				},
			},
		}, nil, nil).
		Once()

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{Transactions: mc}, altsdk.NewClient(""), alttendermint.NewClient(""), properties)

	tx, txErr := adapter.BlockTransaction(context.Background(), &types.BlockTransactionRequest{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: "1",
		},
	})
	require.Nil(t, txErr)
	require.NotNil(t, tx)

	require.Equal(t, &types.BlockTransactionResponse{
		Transaction: &types.Transaction{
			TransactionIdentifier: &types.TransactionIdentifier{
				Hash: "1",
			},
			Operations: []*types.Operation{
				{
					OperationIdentifier: &types.OperationIdentifier{},
					Type:                OperationTransfer,
					Status:              StatusSuccess,
					Account: &types.AccountIdentifier{
						Address: "3",
					}, Amount: &types.Amount{
						Value: "-4",
						Currency: &types.Currency{
							Symbol: "5",
						},
					},
				},
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 1,
					},
					Type:   OperationTransfer,
					Status: StatusSuccess,
					Account: &types.AccountIdentifier{
						Address: "6",
					},
					Amount: &types.Amount{
						Value: "4",
						Currency: &types.Currency{
							Symbol: "5",
						},
					},
				},
			},
		},
	}, tx)
}

func TestLaunchpad_BlockTransactionWithError(t *testing.T) {
	mc := &cosmosmocks.CosmosTransactionsAPI{}
	defer mc.AssertExpectations(t)

	mc.
		On("TxsHashGet", mock.Anything, "1").
		Return(cosmosclient.TxQuery{
			Txhash: "1",
			Code:   7,
			Tx: cosmosclient.StdTx{
				Value: cosmosclient.StdTxValue{
					Msg: []cosmosclient.Msg{
						{
							Type: "cosmos-sdk/MsgSend",
							Value: cosmosclient.MsgValue{
								FromAddress: "3",
								ToAddress:   "6",
								Amount: []cosmosclient.Coin{
									{
										Amount: "4",
										Denom:  "5",
									},
								},
							},
						},
					},
				},
			},
		}, nil, nil).
		Once()

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{Transactions: mc}, altsdk.NewClient(""), alttendermint.NewClient(""), properties)
	tx, txErr := adapter.BlockTransaction(context.Background(), &types.BlockTransactionRequest{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: "1",
		},
	})
	require.Nil(t, txErr)
	require.Equal(t, "Reverted", tx.Transaction.Operations[0].Status)
	require.Equal(t, "Reverted", tx.Transaction.Operations[1].Status)
}
