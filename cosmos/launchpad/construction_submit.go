package launchpad

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/coinbase/rosetta-sdk-go/types"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func (l Launchpad) ConstructionSubmit(ctx context.Context, req *types.ConstructionSubmitRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	bz, err := base64.StdEncoding.DecodeString(req.SignedTransaction)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidTransaction, "error decoding tx")
	}

	bzHexString := hex.EncodeToString(bz)
	fmt.Printf("%s", bzHexString)

	resp, _, err := l.tendermint.Tx.BroadcastTxAsync(ctx, bzHexString)
	if err != nil {
		return nil, rosetta.WrapError(ErrNodeConnection, fmt.Sprintf("error broadcasting tx: %s", err))
	}

	return &types.TransactionIdentifierResponse{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: resp.Result.Hash,
		},
	}, nil
}
