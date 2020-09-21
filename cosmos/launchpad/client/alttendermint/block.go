package alttendermint

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type BlockResponse struct {
	BlockId BlockId `json:"block_id,omitempty"`
	Block   Block   `json:"block,omitempty"`
}

type BlockId struct {
	Hash string `json:"hash"`
}

type Block struct {
	Header BlockHeader `json:"header,omitempty"`
}

type BlockHeader struct {
	Height string `json:"height"`
	Time   string `json:"time"`
}

func (c Client) Block(height uint64) (BlockResponse, error) {
	resp, err := http.Get(c.getEndpoint("block"))
	if err != nil {
		return BlockResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return BlockResponse{}, err
	}

	var jsonResp map[string]json.RawMessage
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		return BlockResponse{}, err
	}

	var blockResponse BlockResponse
	err = json.Unmarshal(jsonResp["result"], &blockResponse)
	if err != nil {
		return BlockResponse{}, err
	}

	return blockResponse, nil
}
