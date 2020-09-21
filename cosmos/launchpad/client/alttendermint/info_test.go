package alttendermint

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_NetInfo(t *testing.T) {
	fileData, err := ioutil.ReadFile("testdata/net_info.json")
	require.NoError(t, err)

	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(fileData))
	}))
	defer s.Close()

	client := NewClient(s.URL)

	resp, err := client.NetInfo()
	require.NoError(t, err)

	require.Equal(t, "2", resp.NPeers)
	require.Equal(t, resp.Peers[0].NodeInfo.Id, "2b1df5de9b6d8cae633ee7b13468ce8443de56ee")
	require.Equal(t, resp.Peers[1].NodeInfo.Id, "c75553feba01261ab03d931962e0ba88570f7d96")
}
