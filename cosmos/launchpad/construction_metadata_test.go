package launchpad

import (
	"context"
	sdktypes "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/types"
	"testing"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

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

	m := &mocks.SdkClient{}
	m.
		On("GetAuthAccount", mock.Anything, "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na").
		Return(sdktypes.AccountResponse{
			Height: 12,
			Result: sdktypes.Response{
				Value: sdktypes.BaseAccount{
					AccountNumber: 0,
					Address:       "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na",
					Sequence:      1,
				},
			},
		}, nil, nil).Once()

	feeMultiplier := float64(200000)
	options := map[string]interface{}{
		OptionAddress: "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na",
		OptionGas:     &feeMultiplier,
	}

	expMetadata := map[string]interface{}{
		AccountNumberKey: uint64(0),
		SequenceKey:      uint64(1),
		ChainIdKey:       "TheNetwork",
		OptionGas:        &feeMultiplier,
	}
	adapter := NewLaunchpad(m, tendermint.NewClient(""), properties)
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

func TestLaunchpad_ConstructionMetadata_FailsOfflineMode(t *testing.T) {
	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
		},
		OfflineMode: true,
	}

	feeMultiplier := float64(200000)
	options := map[string]interface{}{
		OptionAddress: "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na",
		OptionGas:     &feeMultiplier,
	}

	adapter := NewLaunchpad(CosmosAPI{}, altsdk.NewClient(""), tendermint.NewClient(""), properties)
	_, err := adapter.ConstructionMetadata(context.Background(), &types.ConstructionMetadataRequest{
		Options: options,
	})

	require.Equal(t, ErrEndpointDisabledOfflineMode, err)
}
