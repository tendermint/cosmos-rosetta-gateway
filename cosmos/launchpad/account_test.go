package launchpad

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/alttendermint"

	mocks3 "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/alttendermint/mocks"

	mocks2 "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/mocks"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/mocks"
	sdktypes "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/types"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_AccountBalance(t *testing.T) {
	m := &mocks.SdkClient{}
	mt := &mocks2.TendermintInfoAPI{}
	ma := &mocks3.TendermintClient{}
	defer m.AssertExpectations(t)
	m.
		On("GetAuthAccount", mock.Anything, "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na").
		Return(sdktypes.AccountResponse{
			Height: 12,
			Result: sdktypes.Response{
				Value: sdktypes.BaseAccount{
					AccountNumber: 0,
					Coins: []sdk.Coin{
						{Denom: "stake", Amount: sdk.NewInt(400)},
						{Denom: "token", Amount: sdk.NewInt(600)},
					},
					Address:  "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na",
					Sequence: 1,
				},
			},
		}, nil, nil).Once()

	blockHash := "ABCDEFG"
	ma.
		On("Block", uint64(12345)).
		Return(alttendermint.BlockResponse{
			BlockId: alttendermint.BlockId{
				Hash: blockHash,
			},
			Block: alttendermint.Block{},
		}, nil, nil)

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
	}

	adapter := NewLaunchpad(TendermintAPI{Info: mt}, m, ma, properties)

	res, err := adapter.AccountBalance(context.Background(), &types.AccountBalanceRequest{
		AccountIdentifier: &types.AccountIdentifier{
			Address: "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na",
		},
	})
	require.Nil(t, err)
	require.Len(t, res.Balances, 2)
	require.Equal(t, res.BlockIdentifier.Hash, blockHash)
	require.Equal(t, res.BlockIdentifier.Index, int64(12345))

	// NewCoins sorts the coins by name.
	require.Equal(t, "400", res.Balances[0].Value)
	require.Equal(t, "stake", res.Balances[0].Currency.Symbol)
	require.Equal(t, "600", res.Balances[1].Value)
	require.Equal(t, "token", res.Balances[1].Currency.Symbol)
}
