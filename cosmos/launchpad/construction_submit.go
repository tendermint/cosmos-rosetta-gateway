package launchpad

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/coinbase/rosetta-sdk-go/types"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func (l Launchpad) ConstructionSubmit(ctx context.Context, req *types.ConstructionSubmitRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	bz, err := hex.DecodeString(req.SignedTransaction)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidTransaction, "error decoding tx")
	}

	resp, err := l.altCosmos.Broadcast(bz)
	if err != nil {
		return nil, rosetta.WrapError(ErrNodeConnection, fmt.Sprintf("error broadcasting tx: %s", err))
	}

	return &types.TransactionIdentifierResponse{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: resp.TxHash,
		},
		Metadata: map[string]interface{}{
			"log": resp.RawLog,
		},
	}, nil
}
