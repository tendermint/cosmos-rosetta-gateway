package rosetta

import (
	"fmt"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// NewError constructs an error for the Rosetta API calls.
func NewError(code int32, msg string) *types.Error {
	return &types.Error{
		Code:    code,
		Message: msg,
	}
}

func WrapError(err *types.Error, msg string) *types.Error {
	return &types.Error{
		Code:    err.Code,
		Message: fmt.Sprintf("%s: %s", err.Message, msg),
	}
}
