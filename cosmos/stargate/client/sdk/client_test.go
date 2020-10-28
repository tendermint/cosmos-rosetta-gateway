package sdk

import (
	"context"
	"testing"
)

func TestNewClient(t *testing.T) {

	x, err := NewClient("127.0.0.1:9090")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := x.GetAuthAccount(context.Background(), "cosmos1uwwucsjpfus9ur29wwtal99cn37e73ax6dtte6", 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
