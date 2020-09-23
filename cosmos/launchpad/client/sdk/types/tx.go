package types

import (
	sdk "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type BroadcastReq struct {
	Tx   sdk.StdTx `json:"tx" yaml:"tx"`
	Mode string    `json:"mode" yaml:"mode"`
}
