package launchpad

import (
	"context"
	"fmt"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/require"
)

func TestLaunchpad_AccountBalance(t *testing.T) {
	adapter := NewLaunchpad("http://localhost:1317/")

	res, err := adapter.AccountBalance(context.Background(), &types.AccountBalanceRequest{
		NetworkIdentifier: nil,
		AccountIdentifier: &types.AccountIdentifier{
			Address: "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na",
		},
		BlockIdentifier: nil,
	})
	require.Nil(t, err)

	fmt.Printf("%v\n", res.Balances[0].Value)
}
