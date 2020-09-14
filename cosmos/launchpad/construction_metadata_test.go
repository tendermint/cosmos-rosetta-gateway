package launchpad

import (
	"context"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/altsdk"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/mocks"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_ConstructionMetadata(t *testing.T) {
	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
		},
	}

	networkIdentifier := types.NetworkIdentifier{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
	}

	m := &mocks.CosmosAuthAPI{}
	m.
		On("AuthAccountsAddressGet", mock.Anything, "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na").
		Return(cosmosclient.InlineResponse2005{
			Value: cosmosclient.InlineResponse2005Value{
				AccountNumber: "0",
				Address:       "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na",
				Sequence:      "1",
			},
		}, nil, nil).Once()

	feeMultiplier := float64(200000)
	options := map[string]interface{}{
		OptionAddress: "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na",
		OptionGas:     &feeMultiplier,
	}

	expMetadata := map[string]interface{}{
		AccountNumberKey: "0",
		SequenceKey:      "1",
		ChainIdKey:       "TheNetwork",
	}
	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{Auth: m}, altsdk.NewClient(""), properties)
	metaResp, err := adapter.ConstructionMetadata(context.Background(), &types.ConstructionMetadataRequest{
		NetworkIdentifier: &networkIdentifier,
		Options:           options,
	})

	require.Nil(t, err)
	require.NotNil(t, metaResp)
	if diff := cmp.Diff(metaResp.Metadata, expMetadata); diff != "" {
		t.Errorf("Metadata mismatch %s", diff)
	}
}
