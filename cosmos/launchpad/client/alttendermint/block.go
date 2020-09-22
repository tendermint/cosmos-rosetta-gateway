package alttendermint

import (
	"encoding/json"
	"fmt"
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
	LastBlockId BlockId `json:"last_block_id"`
	Height string `json:"height"`
	Time   string `json:"time"`
}

func (c Client) Block(height uint64) (BlockResponse, error) {
	resp, err := http.Get(c.getEndpoint(fmt.Sprintf("block?height=%d", height)))
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

func (c Client) BlockByHash(hash string) (BlockResponse, error) {
	resp, err := http.Get(c.getEndpoint(fmt.Sprintf("block_by_hash?hash=%s", hash)))
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
