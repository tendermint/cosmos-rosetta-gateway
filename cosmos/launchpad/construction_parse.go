package launchpad

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"github.com/coinbase/rosetta-sdk-go/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
)

const (
	Memo = "memo"
)

// ConstructionParse implements the /construction/parse endpoint.
func (l Launchpad) ConstructionParse(ctx context.Context, request *types.ConstructionParseRequest) (*types.ConstructionParseResponse, *types.Error) {
	var stdTx cosmosclient.StdTx

	rawTx, err := hex.DecodeString(request.Transaction)
	if err != nil {
		if rawTx, err = base64.StdEncoding.DecodeString(request.Transaction); err != nil {
			return nil, ErrTxMalformed
		}
	}
	if err := json.Unmarshal(rawTx, &stdTx); err != nil {
		return nil, ErrTxUnmarshal
	}

	var signers []string
	for _, s := range stdTx.Value.Signatures {
		addr := cosmostypes.AccAddress(s.PubKey.Value).String()
		signers = append(signers, addr)
	}

	metadata := make(map[string]interface{})
	if stdTx.Value.Memo != "" {
		metadata[Memo] = stdTx.Value.Memo
	}

	return &types.ConstructionParseResponse{
		Operations: toOperations(stdTx.Value.Msg, false),
		Signers:    signers,
		Metadata:   metadata,
	}, nil
}
