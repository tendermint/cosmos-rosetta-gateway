package launchpad

import (
	"context"
	"encoding/base64"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	"github.com/tendermint/tendermint/crypto"
)

// ConstructionCombine implements the /construction/combine endpoint.
func (l Launchpad) ConstructionCombine(ctx context.Context, r *types.ConstructionCombineRequest,
) (*types.ConstructionCombineResponse, *types.Error) {
	bz, err := base64.StdEncoding.DecodeString(r.UnsignedTransaction)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidTransaction, "error decoding tx")
	}

	codec := simapp.MakeCodec()
	var stdTx auth.StdTx
	err = codec.UnmarshalBinaryLengthPrefixed(bz, &stdTx)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidTransaction, "error unmarshalling tx")
	}
	var pk crypto.PubKey
	var sigs []auth.StdSignature

	for _, signature := range r.Signatures {
		err = codec.UnmarshalBinaryBare(signature.PublicKey.Bytes, &pk)
		sign := auth.StdSignature{
			PubKey:    pk,
			Signature: signature.Bytes,
		}
		sigs = append(sigs, sign)
	}

	stdTx.Signatures = sigs
	txBytes, err := codec.MarshalBinaryLengthPrefixed(stdTx)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidTransaction, "error marshaling tx")
	}
	txBase64 := base64.StdEncoding.EncodeToString(txBytes)

	return &types.ConstructionCombineResponse{
		SignedTransaction: txBase64,
	}, nil
}
