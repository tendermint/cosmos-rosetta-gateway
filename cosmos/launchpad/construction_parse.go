package launchpad

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/coinbase/rosetta-sdk-go/types"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

const (
	Memo = "memo"
)

// ConstructionParse implements the /construction/parse endpoint.
func (l Launchpad) ConstructionParse(ctx context.Context, request *types.ConstructionParseRequest) (*types.ConstructionParseResponse, *types.Error) {
	var stdTx cosmosclient.StdTx

	// TODO: handle both base64 and hex
	rawTx, err := base64.StdEncoding.DecodeString(request.Transaction)
	if err != nil {
		return nil, rosetta.NewError(5, err.Error())
	}
	if err := json.Unmarshal(rawTx, &stdTx); err != nil {
		return nil, rosetta.NewError(5, err.Error())
	}

	res := &types.ConstructionParseResponse{
		Operations: toOperations(stdTx.Value.Msg),
		//Signers:    signers,
		Metadata: map[string]interface{}{
			//Memo: stdTx.Memo,
		},
	}
	return res, nil
}
