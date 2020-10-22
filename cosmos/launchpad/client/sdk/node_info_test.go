package sdk

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/clienttest"
)

func TestGetNodeInfo(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx, cancel := clienttest.Ctx()
	t.Cleanup(cancel)
	e, err := clienttest.NewLaunchpad(ctx, "crgapp")
	require.NoError(t, err)
	t.Cleanup(e.Cleanup)

	client := NewClient(e.SDKAddr)

	res, err := client.GetNodeInfo(ctx)
	require.NoError(t, err)
	require.NotNil(t, res)

	require.Equal(t, "mynode", res.Moniker)
	require.Equal(t, "0.33.7", res.Version)
}
