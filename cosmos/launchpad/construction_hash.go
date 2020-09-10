package launchpad

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) ConstructionHash(ctx context.Context, req *types.ConstructionHashRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	bz, err := base64.StdEncoding.DecodeString(req.SignedTransaction)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidTransaction, "error decoding tx")
	}

	hash := sha256.Sum256(bz)
	bzHash := hash[:]

	hashString := hex.EncodeToString(bzHash)

	return &types.TransactionIdentifierResponse{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: hashString,
		},
	}, nil
}
