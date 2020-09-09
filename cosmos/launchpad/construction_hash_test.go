package launchpad

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"
)

func TestLaunchpad_ConstructionHash(t *testing.T) {
	open, err := os.Open("./testdata/signed-tx.json")
	require.NoError(t, err)

	expectedHash := "6F22EA7620EBCB5078D244F06E88DD26906BA1685135BFC34F83FEFDD653198A"

	bz, err := ioutil.ReadAll(open)
	require.NoError(t, err)

	var stdTx auth.StdTx
	err = bank.ModuleCdc.UnmarshalJSON(bz, &stdTx)
	require.NoError(t, err)

	fmt.Printf("%v\n", stdTx.Signatures)
}
