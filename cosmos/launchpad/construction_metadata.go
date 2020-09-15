package launchpad

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) ConstructionMetadata(ctx context.Context, r *types.ConstructionMetadataRequest) (*types.ConstructionMetadataResponse, *types.Error) {
	if len(r.Options) == 0 {
		return nil, ErrInvalidRequest
	}

	addr, ok := r.Options[OptionAddress]
	if !ok {
		return nil, ErrInvalidAddress
	}
	addrString := addr.(string)
	accRes, _, err := l.cosmos.Auth.AuthAccountsAddressGet(ctx, addrString)
	if err != nil {
		return nil, ErrInterpreting
	}

	// TODO: Check if suggested fee can be added
	res := &types.ConstructionMetadataResponse{
		Metadata: map[string]interface{}{
			AccountNumberKey: accRes.Result.Value.AccountNumber,
			SequenceKey:      accRes.Result.Value.Sequence,
			ChainIdKey:       r.NetworkIdentifier.Network,
		},
	}

	return res, nil
}
