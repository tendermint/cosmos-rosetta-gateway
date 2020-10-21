package launchpad

import (
	"context"
	"encoding/json"

	"github.com/coinbase/rosetta-sdk-go/types"
)

type BroadcastReq struct {
	Tx   json.RawMessage `json:"tx"`
	Mode string          `json:"mode"`
}

func (l launchpad) ConstructionSubmit(ctx context.Context, req *types.ConstructionSubmitRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	return nil, nil
}
