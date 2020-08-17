package launchpad

import (
	"context"
	"net/http"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	client "github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad/client/cosmos/generated"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad/client/cosmos/mocks"
)

func TestLaunchpad_AccountBalance(t *testing.T) {
	m := &mocks.BankAPI{}
	defer m.AssertExpectations(t)

	m.
		On("BankBalancesAddressGet", mock.Anything, "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na").
		Return(client.InlineResponse2004{
			Result: []client.Coin{
				{"stake", "400"},
				{"token", "600"},
			},
		}, nil, nil).
		Once()

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
	}

	adapter := NewLaunchpad(http.DefaultClient, API{Bank: m}, "", properties)

	res, err := adapter.AccountBalance(context.Background(), &types.AccountBalanceRequest{
		AccountIdentifier: &types.AccountIdentifier{
			Address: "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na",
		},
	})
	require.Nil(t, err)
	require.Len(t, res.Balances, 2)

	// NewCoins sorts the coins by name.
	require.Equal(t, "400", res.Balances[0].Value)
	require.Equal(t, "stake", res.Balances[0].Currency.Symbol)
	require.Equal(t, "600", res.Balances[1].Value)
	require.Equal(t, "token", res.Balances[1].Currency.Symbol)
}
