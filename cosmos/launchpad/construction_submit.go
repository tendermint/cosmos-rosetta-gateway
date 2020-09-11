package launchpad

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/x/auth/types"
	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"

	"github.com/coinbase/rosetta-sdk-go/types"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func (l Launchpad) ConstructionSubmit(ctx context.Context, req *types.ConstructionSubmitRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	bz, err := hex.DecodeString(req.SignedTransaction)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidTransaction, "error decoding tx")
	}

	cdc := simapp.MakeCodec()

	var stdTx sdk.StdTx
	err = cdc.UnmarshalBinaryLengthPrefixed(bz, &stdTx)

	txBroadcast := cosmosclient.InlineObject{
		Tx:   mapStdTxToApiStdTx(stdTx),
		Mode: "block",
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

func mapStdTxToApiStdTx(stdTx sdk.StdTx) cosmosclient.StdTx {
	clientStdTx := cosmosclient.StdTx{Value: cosmosclient.StdTxValue{
		Memo: stdTx.Memo,
	}}

	for _, msg := range stdTx.Msgs {
		sendTxMsg := msg.(bank.MsgSend)

		var amount []cosmosclient.Coin
		for _, c := range sendTxMsg.Amount {
			amount = append(amount, cosmosclient.Coin{
				Denom:  c.Denom,
				Amount: c.Amount.String(),
			})
		}

		m := cosmosclient.Msg{
			Type: msg.Type(),
			Value: cosmosclient.MsgValue{
				FromAddress: sendTxMsg.FromAddress.String(),
				ToAddress:   sendTxMsg.ToAddress.String(),
				Amount:      amount,
			},
		}

		clientStdTx.Value.Msg = append(clientStdTx.Value.Msg, m)
	}

	for _, s := range stdTx.Signatures {
		clientStdTx.Value.Signatures = append(clientStdTx.Value.Signatures, cosmosclient.StdTxValueSignatures{
			Signature: hex.EncodeToString(s.Signature),
		})
	}

	var fees []cosmosclient.Coin
	for _, c := range stdTx.Fee.Amount {
		clientStdTx.Value.Fee.Amount = append(clientStdTx.Value.Fee.Amount, cosmosclient.Coin{
			Denom:  c.Denom,
			Amount: c.Amount.String(),
		})
	}
	clientStdTx.Value.Fee = cosmosclient.StdTxValueFee{
		Gas:    strconv.Itoa(int(stdTx.GetGas())),
		Amount: fees,
	}

	return clientStdTx
}
