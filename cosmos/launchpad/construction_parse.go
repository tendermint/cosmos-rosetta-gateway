package launchpad

import (
	"context"
	"encoding/base64"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

const (
	Memo = "memo"
)

// ConstructionParse implements the /construction/parse endpoint.
func (l Launchpad) ConstructionParse(ctx context.Context, request *types.ConstructionParseRequest) (*types.ConstructionParseResponse, *types.Error) {
	var stdTx auth.StdTx

	// TODO: handle both base64 and hex
	rawTx, err := base64.StdEncoding.DecodeString(request.Transaction)
	if err != nil {
		return nil, rosetta.NewError(5, err.Error())
	}

	codec := simapp.MakeCodec()
	err = codec.UnmarshalBinaryLengthPrefixed(rawTx, &stdTx)
	if err != nil {
		return nil, rosetta.NewError(5, err.Error())
	}

	var signers []string
	addrs := stdTx.GetSigners()
	for i := range addrs {
		signers = append(signers, addrs[i].String())
	}

	// TODO: Convert msgs to operations
	res := &types.ConstructionParseResponse{
		Operations: nil,
		Signers:    signers,
		Metadata: map[string]interface{}{
			Memo: stdTx.Memo,
		},
	}
	return res, nil
}
