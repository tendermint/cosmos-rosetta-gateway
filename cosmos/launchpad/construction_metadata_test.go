package launchpad

import (
	"context"
	"testing"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk"
	sdktypes "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/types"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/mocks"
	mocks1 "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/mocks"
)

func TestLaunchpad_ConstructionMetadata(t *testing.T) {
	properties := properties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		AddrPrefix: "test",
	}

	networkIdentifier := types.NetworkIdentifier{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
	}

	var (
		m  = &mocks.SdkClient{}
		mt = &mocks1.TendermintClient{}
	)

	m.
		On("GetAuthAccount", mock.Anything, "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na", int64(0)).
		Return(sdktypes.AccountResponse{
			Height: 12,
			Result: sdktypes.Response{
				Value: sdktypes.BaseAccount{
					AccountNumber: "0",
					Address:       "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na",
					Sequence:      "1",
				},
			},
		}, nil, nil).Once()

	mt.
		On("Status", mock.Anything).
		Return(tendermint.StatusResponse{
			NodeInfo: tendermint.StatusNodeInfo{
				Network: "TheNetwork",
			},
		}, nil, nil).Once()

	feeMultiplier := float64(200000)
	options := map[string]interface{}{
		OptionAddress: "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na",
		OptionGas:     &feeMultiplier,
		OptionMemo:    "test",
		OptionFee:     "12345stake",
	}

	expMetadata := map[string]interface{}{
		AccountNumberKey: "0",
		SequenceKey:      "1",
		ChainIdKey:       "TheNetwork",
		OptionGas:        &feeMultiplier,
		OptionMemo:       "test",
		OptionFee:        "12345stake",
	}

	adapter := newAdapter(m, mt, properties)
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
	properties := properties{
		Blockchain:  "TheBlockchain",
		Network:     "TheNetwork",
		OfflineMode: true,
	}

	feeMultiplier := float64(200000)
	options := map[string]interface{}{
		OptionAddress: "cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na",
		OptionGas:     &feeMultiplier,
	}

	adapter := newAdapter(sdk.NewClient(""), tendermint.NewClient(""), properties)
	_, err := adapter.ConstructionMetadata(context.Background(), &types.ConstructionMetadataRequest{
		Options: options,
	})

	require.Equal(t, ErrEndpointDisabledOfflineMode, err)
}
