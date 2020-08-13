package jsonutil

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	type data struct {
		No Int
	}
	cases := []struct {
		expected int64
		given    string
		canErr   bool
	}{
		{4000, `{"No": "4000"}`, false},
		{0, `{"No": "4.000"}`, true},
	}
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%d", tt.expected), func(t *testing.T) {
			var o data
			err := json.Unmarshal([]byte(tt.given), &o)
			if tt.canErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.expected, o.No.Int64())
		})
	}
}
