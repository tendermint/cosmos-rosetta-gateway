package launchpad

import (
	"context"
	"github.com/coinbase/rosetta-sdk-go/types"
	"testing"
)

func TestLaunchpad_AccountBalance(t *testing.T) {
	adapter := Launchpad{}

	adapter.AccountBalance(context.Background(), &types.AccountBalanceRequest{
		NetworkIdentifier: nil,
		AccountIdentifier: nil,
		BlockIdentifier:   nil,
	})
}
