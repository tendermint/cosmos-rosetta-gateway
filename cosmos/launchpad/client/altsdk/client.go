package altsdk

import (
	"fmt"
)

const (
	BroadcastEndpoint = "/txs"
)

type Client interface {
	Broadcast(request []byte) (*BroadcastResp, error)
}

type client struct {
	endpoint string
}

// NewClient returns the client to call Cosmos RPC.
func NewClient(endpoint string) Client {
	return &client{
		endpoint: endpoint,
	}
}

func (c client) buildEndpoint(path string) string {
	return fmt.Sprintf("%s%s", c.endpoint, path)
}
