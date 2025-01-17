package launch_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch"
)

func TestGenesis(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)

	original := sample.LaunchGenesisState()
	launch.InitGenesis(ctx, *keeper, original)
	got := launch.ExportGenesis(ctx, *keeper)

	// Compare lists
	require.Len(t, got.ChainList, len(original.ChainList))
	require.Subset(t, original.ChainList, got.ChainList)

	require.Len(t, got.ChainNameCountList, len(original.ChainNameCountList))
	require.Subset(t, original.ChainNameCountList, got.ChainNameCountList)

	require.Len(t, got.GenesisAccountList, len(original.GenesisAccountList))
	require.Subset(t, original.GenesisAccountList, got.GenesisAccountList)

	require.Len(t, got.VestedAccountList, len(original.VestedAccountList))
	require.Subset(t, original.VestedAccountList, got.VestedAccountList)

	require.Len(t, got.GenesisValidatorList, len(original.GenesisValidatorList))
	require.Subset(t, original.GenesisValidatorList, got.GenesisValidatorList)

	require.Len(t, got.RequestList, len(original.RequestList))
	require.Subset(t, original.RequestList, got.RequestList)

	require.Len(t, got.RequestCountList, len(original.RequestCountList))
	require.Subset(t, original.RequestCountList, got.RequestCountList)

	require.Equal(t, original.Params, got.Params)
}
