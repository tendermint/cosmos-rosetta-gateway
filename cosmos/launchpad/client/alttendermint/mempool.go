package alttendermint

type UnconfirmedTxsResponse struct {
	Txs []string `json:"txs"`
}

func (c Client) UnconfirmedTxs() (UnconfirmedTxsResponse, error) {
	panic("implement me")
}
