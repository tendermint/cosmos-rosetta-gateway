package launchpad

import (
	"context"
	"github.com/coinbase/rosetta-sdk-go/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/mocks"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"
	mocks1 "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/mocks"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	"testing"
	"time"
)

func TestLaunchpad_Block(t *testing.T) {
	var (
		mc = &mocks.SdkClient{}
		ma = &mocks1.TendermintClient{}
	)
	defer mc.AssertExpectations(t)

	ti, err := time.Parse(time.RFC3339, "2019-04-22T17:01:51Z")
	require.NoError(t, err)

	ma.
		On("Block", uint64(1)).
		Return(tendermint.BlockResponse{
			BlockId: tendermint.BlockId{
				Hash: "11",
			},
			Block: tendermint.Block{
				Header: tendermint.BlockHeader{
					Height: "2",
					Time:   ti.Format(time.RFC3339),
					LastBlockId: tendermint.BlockId{
						Hash: "12",
					},
				},
			},
		}, nil, nil).
		Once()

	ma.
		On("TxSearch", `tx.height=2`).
		Return(tendermint.TxSearchResponse{
			Txs: []tendermint.TxSearchResponseResultTxs{
				{
					Hash: "3",
				},
				{
					Hash: "4",
				},
			},
		}, nil, nil).
		Once()

	addr1 := sdk.AccAddress([]byte("8"))
	addr2 := sdk.AccAddress([]byte("9"))
	coins1 := sdk.NewCoins(sdk.NewInt64Coin("atom", 10))
	testTx := auth.StdTx{
		Msgs: []sdk.Msg{bank.NewMsgSend(addr1, addr2, coins1)},
	}
	mc.
		On("GetTx", mock.Anything, "3").
		Return(sdk.TxResponse{
			TxHash: "3",
			Tx:     testTx,
		}, nil, nil).
		Once()

	addr3 := sdk.AccAddress([]byte("8"))
	addr4 := sdk.AccAddress([]byte("9"))
	coins2 := sdk.NewCoins(sdk.NewInt64Coin("atom", 10))
	testTx = auth.StdTx{
		Msgs: []sdk.Msg{bank.NewMsgSend(addr3, addr4, coins2)},
	}
	mc.
		On("GetTx", mock.Anything, "4").
		Return(sdk.TxResponse{
			TxHash: "4",
			Tx:     testTx,
		}, nil, nil).
		Once()

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			OperationTransfer,
		},
	}

	adapter := NewLaunchpad(mc, ma, properties)

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
								Address: addr3.String(),
							},
							Amount: &types.Amount{
								Value: "-10",
								Currency: &types.Currency{
									Symbol: "atom",
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
								Address: addr4.String(),
							},
							Amount: &types.Amount{
								Value: "10",
								Currency: &types.Currency{
									Symbol: "atom",
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
								Address: addr1.String(),
							},
							Amount: &types.Amount{
								Value: "-10",
								Currency: &types.Currency{
									Symbol: "atom",
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
								Address: addr2.String(),
							},
							Amount: &types.Amount{
								Value: "10",
								Currency: &types.Currency{
									Symbol: "atom",
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
	mc := &mocks.SdkClient{}
	defer mc.AssertExpectations(t)

	addr1 := sdk.AccAddress("8")
	addr2 := sdk.AccAddress("9")
	coins1 := sdk.NewCoins(sdk.NewInt64Coin("atom", 10))
	testTx := auth.StdTx{
		Msgs: []sdk.Msg{bank.NewMsgSend(addr1, addr2, coins1)},
	}
	mc.
		On("GetTx", mock.Anything, "1").
		Return(sdk.TxResponse{
			TxHash: "1",
			Tx:     testTx,
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

	adapter := NewLaunchpad(mc, tendermint.NewClient(""), properties)

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
						Address: addr1.String(),
					}, Amount: &types.Amount{
						Value: "-10",
						Currency: &types.Currency{
							Symbol: "atom",
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
						Address: addr2.String(),
					},
					Amount: &types.Amount{
						Value: "10",
						Currency: &types.Currency{
							Symbol: "atom",
						},
					},
				},
			},
		},
	}, tx)
}

func TestLaunchpad_BlockTransactionWithError(t *testing.T) {
	mc := &mocks.SdkClient{}
	defer mc.AssertExpectations(t)

	addr1 := sdk.AccAddress("8")
	addr2 := sdk.AccAddress("9")
	coins1 := sdk.NewCoins(sdk.NewInt64Coin("atom", 10))
	testTx := auth.StdTx{
		Msgs: []sdk.Msg{bank.NewMsgSend(addr1, addr2, coins1)},
	}
	mc.
		On("GetTx", mock.Anything, "1").
		Return(sdk.TxResponse{
			TxHash: "1",
			Code:   7,
			Tx:     testTx,
		}, nil, nil).Once()

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	adapter := NewLaunchpad(mc, tendermint.NewClient(""), properties)
	tx, txErr := adapter.BlockTransaction(context.Background(), &types.BlockTransactionRequest{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: "1",
		},
	})
	require.Nil(t, txErr)
	require.Equal(t, "Reverted", tx.Transaction.Operations[0].Status)
	require.Equal(t, "Reverted", tx.Transaction.Operations[1].Status)
}

func TestLaunchpad_Block_DoesNotWorkOfflineMode(t *testing.T) {
	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			OperationTransfer,
		},
		OfflineMode: true,
	}

	adapter := NewLaunchpad(sdk.NewClient(""), tendermint.NewClient(""), properties)

	var height int64 = 1
	_, err := adapter.Block(context.Background(), &types.BlockRequest{
		BlockIdentifier: &types.PartialBlockIdentifier{
			Index: &height,
		},
	})
	require.Equal(t, err, ErrEndpointDisabledOfflineMode)
}

func TestLaunchpad_BlockTransaction_FailsOfflineMode(t *testing.T) {
	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
		OfflineMode: true,
	}

	adapter := NewLaunchpad(sdk.NewClient(""), tendermint.NewClient(""), properties)
	_, txErr := adapter.BlockTransaction(context.Background(), &types.BlockTransactionRequest{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: "1",
		},
	})

	require.Equal(t, txErr, ErrEndpointDisabledOfflineMode)
}
