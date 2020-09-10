package launchpad

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/cosmos/cosmos-sdk/simapp"
	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"

	"github.com/coinbase/rosetta-sdk-go/types"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func (l Launchpad) ConstructionSubmit(ctx context.Context, req *types.ConstructionSubmitRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	bz, err := base64.StdEncoding.DecodeString(req.SignedTransaction)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidTransaction, "error decoding tx")
	}

	var stdTx cosmosclient.StdTxValue
	cdc := simapp.MakeCodec()
	err = cdc.UnmarshalJSON(bz, &stdTx)
	txBroadcast := cosmosclient.InlineObject{
		Tx: cosmosclient.StdTx{
			Value: stdTx,
		},
		Mode: "async",
	}
	resp, _, err := l.cosmos.Transactions.TxsPost(ctx, txBroadcast)
	if err != nil {
		return nil, rosetta.WrapError(ErrNodeConnection, fmt.Sprintf("error broadcasting tx: %s", err))
	}

	return &types.TransactionIdentifierResponse{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: resp.Hash,
		},
	}, nil
}
