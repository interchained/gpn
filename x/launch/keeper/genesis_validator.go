package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

// SetGenesisValidator set a specific genesisValidator in the store from its index
func (k Keeper) SetGenesisValidator(ctx sdk.Context, genesisValidator types.GenesisValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisValidatorKeyPrefix))
	b := k.cdc.MustMarshalBinaryBare(&genesisValidator)
	store.Set(types.GenesisValidatorKey(
		genesisValidator.ChainID,
		genesisValidator.Address,
	), b)
}

// GetGenesisValidator returns a genesisValidator from its index
func (k Keeper) GetGenesisValidator(
	ctx sdk.Context,
	chainID string,
	address string,

) (val types.GenesisValidator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisValidatorKeyPrefix))

	b := store.Get(types.GenesisValidatorKey(
		chainID,
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveGenesisValidator removes a genesisValidator from the store
func (k Keeper) RemoveGenesisValidator(
	ctx sdk.Context,
	chainID string,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisValidatorKeyPrefix))
	store.Delete(types.GenesisValidatorKey(
		chainID,
		address,
	))
}

// GetAllGenesisValidator returns all genesisValidator
func (k Keeper) GetAllGenesisValidator(ctx sdk.Context) (list []types.GenesisValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisValidatorKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GenesisValidator
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
