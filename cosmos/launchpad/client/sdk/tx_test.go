package sdk

import (
	"context"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTx(t *testing.T) {
	bz, err := ioutil.ReadFile("testdata/validtx.json")
	require.NoError(t, err)

	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(bz)
	}))
	defer s.Close()

	client := NewClient(s.URL)

	hash := "CFFE3295A82BC0104F1175C26384235B6B3DA80306597F8590684282E195EF1C"
	res, err := client.GetTx(context.Background(), hash)
	t.Log(res)

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, hash, res.TxHash)
}

func TestPostTx(t *testing.T) {
	bz, err := ioutil.ReadFile("testdata/broadcast.json")
	require.NoError(t, err)

	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(bz)
	}))
	defer s.Close()

	client := NewClient(s.URL)

	res, err := client.PostTx(context.Background(), bz)
	t.Log(res)

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, 1, len(res.Tx.GetMsgs()))
	t.Log(res)
}
