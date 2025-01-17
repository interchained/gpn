package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	profile "github.com/tendermint/spn/x/profile/types"
)

func TestGenesisStateValidateValidator(t *testing.T) {
	addr := sample.AccAddress()
	tests := []struct {
		name     string
		genState *profile.GenesisState
		err      error
	}{
		{
			name:     "default is valid",
			genState: profile.DefaultGenesis(),
		}, {
			name: "valid custom genesis",
			genState: &profile.GenesisState{
				ValidatorList: []profile.Validator{
					{Address: sample.AccAddress()},
					{Address: sample.AccAddress()},
					{Address: sample.AccAddress()},
				},
			},
		}, {
			name: "duplicated validator by address",
			genState: &profile.GenesisState{
				ValidatorList: []profile.Validator{
					{Address: addr},
					{Address: addr},
				},
			},
			err: fmt.Errorf("duplicated index for validator: %s", addr),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				err := tt.genState.ValidateValidators()
				if tt.err != nil {
					require.Error(t, err)
					require.Equal(t, tt.err.Error(), err.Error())
					return
				}
				require.NoError(t, err)
			})
		})
	}
}

func TestGenesisStateValidateCoordinator(t *testing.T) {
	var (
		addr1 = sample.AccAddress()
		addr2 = sample.AccAddress()
		addr3 = sample.AccAddress()
		addr4 = sample.AccAddress()
	)
	tests := []struct {
		name     string
		genState *profile.GenesisState
		err      error
	}{
		{
			name:     "default is valid",
			genState: profile.DefaultGenesis(),
		}, {
			name: "valid custom genesis",
			genState: &profile.GenesisState{
				CoordinatorByAddressList: []profile.CoordinatorByAddress{
					{CoordinatorId: 0, Address: addr1},
					{CoordinatorId: 1, Address: addr2},
					{CoordinatorId: 2, Address: addr3},
					{CoordinatorId: 3, Address: addr4},
				},
				CoordinatorList: []profile.Coordinator{
					{CoordinatorId: 0, Address: addr1},
					{CoordinatorId: 1, Address: addr2},
					{CoordinatorId: 2, Address: addr3},
					{CoordinatorId: 3, Address: addr4},
				},
				CoordinatorCount: 4,
			},
		}, {
			name: "duplicated coordinator",
			genState: &profile.GenesisState{
				CoordinatorByAddressList: []profile.CoordinatorByAddress{
					{CoordinatorId: 0, Address: addr1},
					{CoordinatorId: 1, Address: addr1},
				},
				CoordinatorCount: 2,
			},
			err: fmt.Errorf("duplicated index for coordinatorByAddress: %s", addr1),
		}, {
			name: "duplicated coordinator id",
			genState: &profile.GenesisState{
				CoordinatorByAddressList: []profile.CoordinatorByAddress{
					{CoordinatorId: 0, Address: addr1},
					{CoordinatorId: 0, Address: addr2},
				},
				CoordinatorList: []profile.Coordinator{
					{CoordinatorId: 0, Address: addr1},
					{CoordinatorId: 0, Address: addr2},
				},
				CoordinatorCount: 2,
			},
			err: fmt.Errorf("duplicated id for coordinator: 0"),
		}, {
			name: "profile not associated with chain",
			genState: &profile.GenesisState{
				CoordinatorByAddressList: []profile.CoordinatorByAddress{
					{CoordinatorId: 0, Address: addr1},
				},
				CoordinatorList: []profile.Coordinator{
					{CoordinatorId: 0, Address: addr1},
					{CoordinatorId: 1, Address: addr2},
				},
				CoordinatorCount: 2,
			},
			err: fmt.Errorf("coordinator address not found for CoordinatorByAddress: %s", addr2),
		}, {
			name: "profile not associated with chain",
			genState: &profile.GenesisState{
				CoordinatorByAddressList: []profile.CoordinatorByAddress{
					{CoordinatorId: 0, Address: addr1},
					{CoordinatorId: 1, Address: addr2},
				},
				CoordinatorList: []profile.Coordinator{
					{CoordinatorId: 0, Address: addr1},
				},
				CoordinatorCount: 2,
			},
			err: fmt.Errorf("coordinator address not found for coordinatorID: 1"),
		}, {
			name: "invalid coordinator id",
			genState: &profile.GenesisState{
				CoordinatorByAddressList: []profile.CoordinatorByAddress{
					{CoordinatorId: 0, Address: addr1},
					{CoordinatorId: 133, Address: addr2},
				},
				CoordinatorList: []profile.Coordinator{
					{CoordinatorId: 0, Address: addr1},
					{CoordinatorId: 133, Address: addr2},
				},
				CoordinatorCount: 2,
			},
			err: fmt.Errorf("coordinator id 133 should be lower or equal than the last id 2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				err := tt.genState.ValidateCoordinators()
				if tt.err != nil {
					require.Error(t, err)
					require.Equal(t, tt.err.Error(), err.Error())
					return
				}
				require.NoError(t, err)
			})
		})
	}
}
