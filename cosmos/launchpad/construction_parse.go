package launchpad

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/coinbase/rosetta-sdk-go/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
)

const (
	Memo = "memo"
)

// ConstructionParse implements the /construction/parse endpoint.
func (l Launchpad) ConstructionParse(ctx context.Context, request *types.ConstructionParseRequest) (*types.ConstructionParseResponse, *types.Error) {
	var stdTx struct {
		Msgs          []cosmosclient.Msg `json:"msgs"`
		Memo          string             `json:"memo"`
		ChainId       string             `json:"chain_id"`
		Sequence      string             `json:"sequence"`
		AccountNumber string             `json:"account_number"`
		Signatures    []struct {
			PubKey struct {
				Value string `json:"value"`
			} `json:"pub_key"`
		} `json:"signatures"`
	}

	rawTx, err := hex.DecodeString(request.Transaction)
	if err != nil {
		if rawTx, err = base64.StdEncoding.DecodeString(request.Transaction); err != nil {
			return nil, ErrInvalidTransaction
		}
	}
	if err := json.Unmarshal(rawTx, &stdTx); err != nil {
		return nil, ErrInvalidTransaction
	}

	sequence, err := strconv.ParseUint(stdTx.Sequence, 10, 64)
	if err != nil {
		return nil, ErrInvalidTransaction
	}
	accountNumber, err := strconv.ParseUint(stdTx.AccountNumber, 10, 64)
	if err != nil {
		return nil, ErrInvalidTransaction
	}

	var signers []string
	for _, sig := range stdTx.Signatures {
		addr := cosmostypes.AccAddress(sig.PubKey.Value)
		signers = append(signers, addr.String())
	}

	metadata := map[string]interface{}{
		ChainIdKey:       stdTx.ChainId,
		SequenceKey:      sequence,
		AccountNumberKey: accountNumber,
	}
	if stdTx.Memo != "" {
		metadata[Memo] = stdTx.Memo
	}

	return &types.ConstructionParseResponse{
		Operations: toOperations(stdTx.Msgs, false),
		Signers:    signers,
		Metadata:   metadata,
	}, nil
}
