package launchpad

import (
	"context"
	"encoding/base64"
	"encoding/hex"

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
		if rawTx, err = base64.StdEncoding.DecodeString(request.Transaction); err != nil {
			return nil, ErrInvalidTransaction
		}
	}

	var stdTx auth.StdTx
	err = Codec.UnmarshalJSON(rawTx, &stdTx)

	var signers []string
	for _, sig := range stdTx.Signatures {
		addr, err := types2.AccAddressFromHex(sig.PubKey.Address().String())
		if err != nil {
			return nil, ErrInvalidTransaction
		}
		signers = append(signers, addr.String())
	}

	return &types.ConstructionParseResponse{
		Operations: toOperations(stdTx.Msgs, false, true),
		Signers:    signers,
	}, nil
}
