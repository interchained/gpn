package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgEditChain(t *testing.T) {
	k, _, srv, profileSrv, sdkCtx, _ := setupMsgServer(t)
	ctx := sdk.WrapSDKContext(sdkCtx)
	coordAddress := sample.AccAddress()
	coordAddress2 := sample.AccAddress()
	coordNoExist := sample.AccAddress()
	chainIDNoExist, _ := sample.ChainID(0)

	// Create coordinators
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	_, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	msgCreateCoordinator = sample.MsgCreateCoordinator(coordAddress2)
	_, err = profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	// Create a chain
	msgCreateChain := sample.MsgCreateChain(coordAddress, "foo", "")
	res, err := srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	chainID := res.ChainID

	for _, tc := range []struct {
		name  string
		msg   types.MsgEditChain
		valid bool
	}{
		{
			name: "edit source",
			msg: sample.MsgEditChain(coordAddress, chainID,
				true,
				false,
				false,
			),
			valid: true,
		},
		{
			name: "edit initial genesis with default genesis",
			msg: sample.MsgEditChain(coordAddress, chainID,
				false,
				true,
				false,
			),
			valid: true,
		},
		{
			name: "edit initial genesis with genesis url",
			msg: sample.MsgEditChain(coordAddress, chainID,
				false,
				true,
				true,
			),
			valid: true,
		},
		{
			name: "edit source and initial genesis",
			msg: sample.MsgEditChain(coordAddress, chainID,
				true,
				true,
				true,
			),
			valid: true,
		},
		{
			name: "non existent chain id",
			msg: sample.MsgEditChain(coordAddress, chainIDNoExist,
				true,
				false,
				false,
			),
			valid: false,
		},
		{
			name: "non existent coordinator",
			msg: sample.MsgEditChain(coordNoExist, chainID,
				true,
				false,
				false,
			),
			valid: false,
		},
		{
			name: "invalid coordinator",
			msg: sample.MsgEditChain(coordAddress2, chainID,
				true,
				false,
				false,
			),
			valid: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Fetch the previous state of the chain to perform checks
			var previousChain types.Chain
			var found bool
			if tc.valid {
				previousChain, found = k.GetChain(sdkCtx, tc.msg.ChainID)
				require.True(t, found)
			}

			// Send the message
			_, err := srv.EditChain(ctx, &tc.msg)
			if !tc.valid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			// The chain must continue to exist in the store
			chain, found := k.GetChain(sdkCtx, tc.msg.ChainID)
			require.True(t, found)

			// Unchanged values
			require.EqualValues(t, previousChain.CoordinatorID, chain.CoordinatorID)
			require.EqualValues(t, previousChain.CreatedAt, chain.CreatedAt)
			require.EqualValues(t, previousChain.LaunchTimestamp, chain.LaunchTimestamp)
			require.EqualValues(t, previousChain.LaunchTriggered, chain.LaunchTriggered)

			// Compare changed values
			if tc.msg.SourceURL != "" {
				require.EqualValues(t, tc.msg.SourceURL, chain.SourceURL)
				require.EqualValues(t, tc.msg.SourceHash, chain.SourceHash)
			} else {
				require.EqualValues(t, previousChain.SourceURL, chain.SourceURL)
				require.EqualValues(t, previousChain.SourceHash, chain.SourceHash)
			}

			if tc.msg.InitialGenesis != nil {
				require.EqualValues(t, tc.msg.InitialGenesis, chain.InitialGenesis)
			} else {
				require.EqualValues(t, previousChain.InitialGenesis, chain.InitialGenesis)
			}
		})
	}
}
