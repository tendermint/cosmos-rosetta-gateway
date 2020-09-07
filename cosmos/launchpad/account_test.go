package launchpad

import (
	"context"
	"testing"

	"github.com/antihax/optional"
	openapi "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"

	mocks2 "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/mocks"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/mocks"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_AccountBalance(t *testing.T) {
	m := &mocks.CosmosBankAPI{}
	mt := &mocks2.TendermintInfoAPI{}
	defer m.AssertExpectations(t)

	m.
		On("BankBalancesAddressGet", mock.Anything, "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na").
		Return(cosmosclient.InlineResponse2004{
			Height: "12345",
			Result: []cosmosclient.Coin{
				{Denom: "stake", Amount: "400"},
				{Denom: "token", Amount: "600"},
			},
		}, nil, nil).
		Once()

	blockHash := "ABCDEFG"
	mt.
		On("Block", mock.Anything, &openapi.BlockOpts{
			Height: optional.NewFloat32(float32(12345)),
		}).
		Return(openapi.BlockResponse{
			Result: openapi.BlockComplete{
				BlockId: openapi.BlockId{
					Hash: blockHash,
				},
				Block: openapi.Block{},
			},
		}, nil, nil)

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
	}

	adapter := NewLaunchpad(TendermintAPI{Info: mt}, CosmosAPI{Bank: m}, properties)

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
