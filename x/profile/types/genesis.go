package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		ValidatorList:            []Validator{},
		CoordinatorByAddressList: []CoordinatorByAddress{},
		CoordinatorList:          []Coordinator{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate

	if err := gs.ValidateValidators(); err != nil {
		return err
	}

	return gs.ValidateCoordinators()
}

func (gs GenesisState) ValidateValidators() error {
	// Check for duplicated index in validator
	validatorIndexMap := make(map[string]struct{})
	for _, elem := range gs.ValidatorList {
		index := string(ValidatorKey(elem.Address))
		if _, ok := validatorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for validator: %s", elem.Address)
		}
		validatorIndexMap[index] = struct{}{}
	}
	return nil
}

func (gs GenesisState) ValidateCoordinators() error {
	// Check for duplicated index in coordinatorByAddress
	coordinatorByAddressIndexMap := make(map[string]uint64)
	for _, elem := range gs.CoordinatorByAddressList {
		index := string(CoordinatorByAddressKey(elem.Address))
		if _, ok := coordinatorByAddressIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for coordinatorByAddress: %s", elem.Address)
		}
		coordinatorByAddressIndexMap[index] = elem.CoordinatorId
	}

	// Check for duplicated ID in coordinator
	coordinatorIDMap := make(map[uint64]bool)
	count := gs.GetCoordinatorCount()
	for _, elem := range gs.CoordinatorList {
		if _, ok := coordinatorIDMap[elem.CoordinatorId]; ok {
			return fmt.Errorf("duplicated id for coordinator: %d", elem.CoordinatorId)
		}
		if elem.CoordinatorId >= count {
			return fmt.Errorf("coordinator id %d should be lower or equal than the last id %d",
				elem.CoordinatorId, count)
		}

		index := string(CoordinatorByAddressKey(elem.Address))
		if _, ok := coordinatorByAddressIndexMap[index]; !ok {
			return fmt.Errorf("coordinator address not found for CoordinatorByAddress: %s", elem.Address)
		}
		coordinatorIDMap[elem.CoordinatorId] = true

		// Remove to check if all coordinator by address exist
		delete(coordinatorByAddressIndexMap, index)
	}
	// Check if all coordinator by address exist
	for _, coordinatorID := range coordinatorByAddressIndexMap {
		return fmt.Errorf("coordinator address not found for coordinatorID: %d", coordinatorID)
	}
	return nil
}
