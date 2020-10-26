package launchpad

import (
	"context"
	"encoding/hex"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"

	types2 "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/coinbase/rosetta-sdk-go/types"
)

const (
	Memo = "memo"
)

// ConstructionParse implements the /construction/parse endpoint.
func (l launchpad) ConstructionParse(ctx context.Context, request *types.ConstructionParseRequest) (*types.ConstructionParseResponse, *types.Error) {
	rawTx, err := hex.DecodeString(request.Transaction)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidTransaction, err.Error())
	}

	var stdTx auth.StdTx
	err = Codec.UnmarshalJSON(rawTx, &stdTx)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidTransaction, err.Error())
	}

	var signers []string
	for _, sig := range stdTx.Signatures {
		addr, err := types2.AccAddressFromHex(sig.PubKey.Address().String())
		if err != nil {
			return nil, rosetta.WrapError(ErrInvalidTransaction, err.Error())
		}
		signers = append(signers, addr.String())
	}

	return &types.ConstructionParseResponse{
		Operations: getOpsFromTx(stdTx, false, true),
		Signers:    signers,
	}, nil
}
