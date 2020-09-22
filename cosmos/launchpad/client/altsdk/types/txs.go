package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

// sdk type
//type TxResponse struct {
//	Height    int64           `json:"height"`
//	TxHash    string          `json:"txhash"`
//	Codespace string          `json:"codespace,omitempty"`
//	Code      uint32          `json:"code,omitempty"`
//	Data      string          `json:"data,omitempty"`
//	RawLog    string          `json:"raw_log,omitempty"`
//	Info      string          `json:"info,omitempty"`
//	GasWanted int64           `json:"gas_wanted,omitempty"`
//	GasUsed   int64           `json:"gas_used,omitempty"`
//	Tx        auth.StdTx      `json:"tx,omitempty"`
//	Timestamp string          `json:"timestamp,omitempty"`
//}

type TxQuery struct {
	Code   int32         `json:"code,omitempty"`
	Txhash string        `json:"txhash,omitempty"`
	Height string        `json:"height,omitempty"`
	Tx     auth.StdTx    `json:"tx,omitempty"`
	Result TxQueryResult `json:"result,omitempty"`
}

type TxQueryResult struct {
	Log       string   `json:"log,omitempty"`
	GasWanted string   `json:"gas_wanted,omitempty"`
	GasUsed   string   `json:"gas_used,omitempty"`
	Tags      []KvPair `json:"tags,omitempty"`
}

type KvPair struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type StdTx struct {
	Value StdTxValue `json:"value,omitempty"`
}

// StdTxValue struct for StdTxValue
type StdTxValue struct {
	Msg        []Msg                  `json:"msg,omitempty"`
	Fee        StdTxValueFee          `json:"fee,omitempty"`
	Memo       string                 `json:"memo,omitempty"`
	Signatures []StdTxValueSignatures `json:"signatures,omitempty"`
}

type StdTxValueFee struct {
	Gas    string     `json:"gas,omitempty"`
	Amount []sdk.Coin `json:"amount,omitempty"`
}

type Msg struct {
	Type  string   `json:"type,omitempty"`
	Value MsgValue `json:"value,omitempty"`
}
type MsgValue struct {
	FromAddress string   `json:"from_address,omitempty"`
	ToAddress   string   `json:"to_address,omitempty"`
	Amount      sdk.Coin `json:"amount,omitempty"`
}

type StdTxValueSignatures struct {
	Signature     string    `json:"signature,omitempty"`
	PubKey        PublicKey `json:"pub_key,omitempty"`
	AccountNumber string    `json:"account_number,omitempty"`
	Sequence      string    `json:"sequence,omitempty"`
}
