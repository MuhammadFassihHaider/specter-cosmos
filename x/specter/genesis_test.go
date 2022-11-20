package specter_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "specter/testutil/keeper"
	"specter/testutil/nullify"
	"specter/x/specter"
	"specter/x/specter/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LoanList: []types.Loan{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LoanCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SpecterKeeper(t)
	specter.InitGenesis(ctx, *k, genesisState)
	got := specter.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LoanList, got.LoanList)
	require.Equal(t, genesisState.LoanCount, got.LoanCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
