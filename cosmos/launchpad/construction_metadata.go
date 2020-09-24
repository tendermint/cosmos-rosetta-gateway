package launchpad

import (
	"context"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) ConstructionMetadata(ctx context.Context, r *types.ConstructionMetadataRequest) (*types.ConstructionMetadataResponse, *types.Error) {
	if l.properties.OfflineMode {
		return nil, ErrEndpointDisabledOfflineMode
	}

	if len(r.Options) == 0 {
		return nil, ErrInvalidRequest
	}

	addr, ok := r.Options[OptionAddress]
	if !ok {
		return nil, ErrInvalidAddress
	}
	addrString := addr.(string)
	accRes, err := l.cosmos.GetAuthAccount(ctx, addrString)
	if err != nil {
		return nil, rosetta.WrapError(ErrInterpreting, err.Error())
	}

	gas, ok := r.Options[GasKey]
	if !ok {
		return nil, rosetta.WrapError(ErrInvalidAddress, "gas not set")
	}

	// TODO: Check if suggested fee can be added
	res := &types.ConstructionMetadataResponse{
		Metadata: map[string]interface{}{
			AccountNumberKey: accRes.Result.Value.AccountNumber,
			SequenceKey:      accRes.Result.Value.Sequence,
			ChainIdKey:       r.NetworkIdentifier.Network,
			GasKey:           gas,
		},
	}

	return res, nil
}
