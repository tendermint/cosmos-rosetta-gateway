package launchpad

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestLaunchpad_AccountBalance(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/bank/balances/cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na", r.URL.Path)
		json.NewEncoder(w).Encode(balanceResp{
			Result: sdk.NewCoins(
				sdk.NewCoin("token", sdk.NewInt(600)),
				sdk.NewCoin("stake", sdk.NewInt(400)),
			),
		})
	}))
	defer ts.Close()

	adapter := NewLaunchpad(http.DefaultClient, ts.URL, "TheBlockchain", "TheNetwork")

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
