package launchpad

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

const (
	Memo = "memo"
)

// ConstructionParse implements the /construction/parse endpoint.
func (l launchpad) ConstructionParse(ctx context.Context, request *types.ConstructionParseRequest) (*types.ConstructionParseResponse, *types.Error) {
	return nil, nil
}
