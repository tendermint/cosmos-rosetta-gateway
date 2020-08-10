package rosetta

import "github.com/coinbase/rosetta-sdk-go/types"

// NewError constructs an error for the Rosetta API calls.
func NewError(code int32, msg string) *types.Error {
	return &types.Error{
		Code:    code,
		Message: msg,
	}
}
