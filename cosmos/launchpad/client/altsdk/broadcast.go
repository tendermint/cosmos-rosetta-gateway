package altsdk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type BroadcastResp struct {
	TxHash string `json:"txhash"`
	RawLog string `json:"raw_log"`
}

// Broadcast the specified tx.
func (c Client) Broadcast(req []byte) (*BroadcastResp, error) {
	b := bytes.NewBuffer(req)
	post, err := http.Post(c.getEndpoint(BroadcastEndpoint), "", b)
	if err != nil {
		return nil, err
	}
	defer post.Body.Close()

	resp, err := ioutil.ReadAll(post.Body)
	if err != nil {
		return nil, err
	}

	r := &BroadcastResp{}
	err = json.Unmarshal(resp, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
