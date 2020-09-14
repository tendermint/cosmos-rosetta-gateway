package launchpad

import (
	"context"
	"encoding/hex"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

// ConstructionCombine implements the /construction/combine endpoint.
func (l Launchpad) ConstructionCombine(ctx context.Context, r *types.ConstructionCombineRequest,
) (*types.ConstructionCombineResponse, *types.Error) {
	bz, err := hex.DecodeString(r.UnsignedTransaction)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidTransaction, "error decoding unsigned tx")
	}

	codec := simapp.MakeCodec()
	var stdTx auth.StdTx
	err = codec.UnmarshalBinaryLengthPrefixed(bz, &stdTx)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidTransaction, "unable to unmarshal tx")
	}
	var pk crypto.PubKey
	var sigs []auth.StdSignature

	for _, signature := range r.Signatures {
		if signature.PublicKey.CurveType != "secp256k1" {
			return nil, ErrUnsupportedCurve
		}

		err = codec.UnmarshalBinaryBare(signature.PublicKey.Bytes, &pk)
		if err != nil {
			return nil, rosetta.WrapError(ErrInvalidPubkey, "unable to unmarshal pubkey")
		}

		sign := auth.StdSignature{
			PubKey:    pk,
			Signature: signature.Bytes,
		}
		sigs = append(sigs, sign)
	}

	stdTx.Signatures = sigs
	txBytes, err := codec.MarshalBinaryLengthPrefixed(stdTx)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidTransaction, "unable to marshal signed tx")
	}
	txHex := hex.EncodeToString(txBytes)

	return &types.ConstructionCombineResponse{
		SignedTransaction: txHex,
	}, nil
}
