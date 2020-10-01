package launchpad

import (
	"context"
	"encoding/hex"
	"github.com/coinbase/rosetta-sdk-go/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func (l launchpad) ConstructionPayloads(ctx context.Context, req *types.ConstructionPayloadsRequest) (*types.ConstructionPayloadsResponse, *types.Error) {
	if len(req.Operations) != 2 {
		return nil, ErrInvalidOperation
	}

	if req.Operations[0].Type != req.Operations[1].Type {
		return nil, rosetta.WrapError(ErrInvalidOperation, "operation type mismatch")
	}

	msg, err := getMsgDataFromOperations(req.Operations)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidOperation, err.Error())
	}

	metadata, err := GetMetadataFromPayloadReq(req)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidRequest, err.Error())
	}

	tx := auth.NewStdTx([]cosmostypes.Msg{msg}, auth.StdFee{
		Gas: metadata.Gas,
	}, nil, "TODO memo") // TODO fees and memo(https://github.com/tendermint/cosmos-rosetta-gateway/issues/122).
	signBytes := auth.StdSignBytes(
		metadata.ChainId, metadata.AccountNumber, metadata.Sequence, tx.Fee, tx.Msgs, tx.Memo,
	)
	txBytes, err := Codec.MarshalJSON(tx)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidRequest, err.Error())
	}

	signers := msg.GetSigners()
	if len(signers) != 1 {
		return nil, rosetta.WrapError(ErrInvalidRequest, "invalid number of signers")
	}

	return &types.ConstructionPayloadsResponse{
		UnsignedTransaction: hex.EncodeToString(txBytes),
		Payloads: []*types.SigningPayload{
			{
				Address:       signers[0].String(),
				Bytes:         signBytes,
				SignatureType: "secp256k1",
			},
		},
	}, nil
}
