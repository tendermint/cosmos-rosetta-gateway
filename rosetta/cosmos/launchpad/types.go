package launchpad

import (
	"time"

	"github.com/tendermint/cosmos-rosetta-gateway/util/jsonutil"
)

type nodeInfo struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

type block struct {
	Header blockHeader `json:"header"`
}

type blockHeader struct {
	Height      jsonutil.Int `json:"height"`
	LastBlockID blockID      `json:"last_block_id"`
	Time        time.Time    `json:"time"`
}

type blockID struct {
	Hash string `json:"hash"`
}

type peer struct {
	NodeInfo nodeInfo `json:"node_info"`
}
