package altsdk

import (
	"context"

	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthAccountClient(t *testing.T) {
	bz, err := ioutil.ReadFile("testdata/account.json")
	require.NoError(t, err)

	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(bz)
	}))
	defer s.Close()

	client := NewClient(s.URL)

	addr := "cosmos15lc6l4nm3s9ya5an5vnv9r6na437ajpznkplhx"
	res, err := client.GetAuthAccount(context.Background(), addr)
	require.NoError(t, err)
	require.NotNil(t, res)
	t.Log(res.Result)

	require.Equal(t, int64(9694), res.Height)
	require.Equal(t, addr, res.Result.Value.Address)
	require.Equal(t, 2, int(res.Result.Value.AccountNumber))
	require.Equal(t, 4, int(res.Result.Value.Sequence))
	require.Equal(t, int64(1000), res.Result.Value.Coins[0].Amount.Int64())
}
