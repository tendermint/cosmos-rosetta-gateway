package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/irisnet/irishub/modules/auth"

	"github.com/irisnet/irishub/types"

	"github.com/irisnet/irishub/app"
)

type TxResponse struct {
	Code   int64
	TxHash string
	Tx     types.Tx
	Height int64
}

func (c Client) GetTx(ctx context.Context, hash string) (TxResponse, error) {
	path := fmt.Sprintf("/txs/%s", hash)

	r, err := http.Get(c.getEndpoint(path))
	if err != nil {
		return TxResponse{}, err
	}
	if r == nil {
		return TxResponse{}, fmt.Errorf("unable to fetch data from endpoint %s", c.getEndpoint(path))
	}
	btes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return TxResponse{}, err
	}

	var jsonTxRes map[string]json.RawMessage
	err = json.Unmarshal(btes, &jsonTxRes)
	if err != nil {
		return TxResponse{}, err
	}

	var txRes TxResponse
	var stdTx auth.StdTx
	codec := app.MakeLatestCodec()
	if err = codec.UnmarshalJSON(jsonTxRes["tx"], &stdTx); err != nil {
		return TxResponse{}, err
	}

	txRes.Tx = stdTx
	txRes.TxHash = hash

	return txRes, nil
}

func (c Client) PostTx(ctx context.Context, bz []byte) (TxResponse, error) {
	r, err := http.Post(c.getEndpoint("/txs"), "application/json", bytes.NewBuffer(bz))
	if err != nil {
		return TxResponse{}, err
	}
	if r == nil {
		return TxResponse{}, fmt.Errorf("unable to get response from endpoint %s", c.getEndpoint("/txs"))
	}
	btes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return TxResponse{}, err
	}

	var txRes TxResponse
	codec := app.MakeLatestCodec()
	if err = codec.UnmarshalJSON(btes, &txRes); err != nil {
		return TxResponse{}, err
	}

	return txRes, nil
}
