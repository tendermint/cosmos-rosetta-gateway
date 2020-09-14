package altsdk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_Broadcast(t *testing.T) {
	client := NewClient("http://localhost:1317")

	err := client.Broadcast(BroadcastRequest{
		Mode: "sync",
	})
	require.NoError(t, err)
}
