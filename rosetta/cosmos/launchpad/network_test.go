package launchpad

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLaunchpad_NetworkList(t *testing.T) {
	blockchain := "TheBlockchain"
	network := "TheNetwork"

	adapter := NewLaunchpad(nil, "http://the-url", blockchain, network)

	list, err := adapter.NetworkList(context.Background(), nil)
	require.Nil(t, err)

	require.Len(t, list.NetworkIdentifiers, 1)
	require.Equal(t, list.NetworkIdentifiers[0].Network, network)
	require.Equal(t, list.NetworkIdentifiers[0].Blockchain, blockchain)
}
