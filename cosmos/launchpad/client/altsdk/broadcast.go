package altsdk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type BroadcastRequest struct {
	Tx   types.StdTx `json:"tx"`
	Mode string      `json:"mode"`
}

// Broadcast the specified tx.
func (c client) Broadcast(req []byte) error {
	b := bytes.NewBuffer(req)
	post, err := http.Post(c.buildEndpoint(BroadcastEndpoint), "", b)
	if err != nil {
		return err
	}
	defer post.Body.Close()

	resp, err := ioutil.ReadAll(post.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s", resp)

	return nil
}
