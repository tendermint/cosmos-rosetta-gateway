package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccountResponse struct {
	Height int64    `json:"height,omitempty"`
	Result Response `json:"result,omitempty"`
}

type Response struct {
	Type  string      `json:"type,omitempty"`
	Value BaseAccount `json:"value,omitempty"`
}

type BaseAccount struct {
	Address       string    `json:"address" yaml:"address"`
	Coins         sdk.Coins `json:"coins" yaml:"coins"`
	AccountNumber uint64    `json:"account_number" yaml:"account_number"`
	Sequence      uint64    `json:"sequence" yaml:"sequence"`
}
