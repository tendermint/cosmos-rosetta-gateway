package launchpad

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) ConstructionMetadata(ctx context.Context, r *types.ConstructionMetadataRequest) (*types.ConstructionMetadataResponse, *types.Error) {
	// validate network identifier

	if r.Options == nil {
		return nil, ErrInterpreting
	}

	addr := r.Options[OptionAddress]
	addrString := addr.(string)
	accRes, _, err := l.cosmos.Auth.AuthAccountsAddressGet(ctx, addrString)
	if err != nil {
		return nil, ErrInvalidAddress
	}

	res := &types.ConstructionMetadataResponse{
		Metadata: map[string]interface{}{
			OptionsAccountNumber: accRes.Value.AccountNumber,
			OptionsSequence:      accRes.Value.Sequence,
			OptionsChainId:       r.NetworkIdentifier.Network,
		},
	}

	return res, nil
}
