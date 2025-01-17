package keeper

import (
	"context"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) CreateChain(goCtx context.Context, msg *types.MsgCreateChain) (*types.MsgCreateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the coordinator ID associated to the sender address
	coordID, found := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}

	// Compute the chain id
	chainNameCount, found := k.GetChainNameCount(ctx, msg.ChainName)
	if !found {
		chainNameCount = types.ChainNameCount{
			ChainName: msg.ChainName,
			Count:     0,
		}
	}
	chainID := types.ChainIDFromChainName(msg.ChainName, chainNameCount.Count)
	chainNameCount.Count++

	// chainID must always be unique by design
	// if it already exists then something is wrong in the protocol
	_, found = k.GetChain(ctx, chainID)
	if found {
		return nil, spnerrors.Criticalf("chain id %s already exists while it must be unique", chainID)
	}

	// Initialize the chain
	chain := types.Chain{
		ChainID:         chainID,
		CoordinatorID:   coordID,
		CreatedAt:       ctx.BlockTime().Unix(),
		SourceURL:       msg.SourceURL,
		SourceHash:      msg.SourceHash,
		LaunchTriggered: false,
		LaunchTimestamp: 0,
	}

	// Initialize initial genesis
	var err error
	if msg.GenesisURL == "" {
		chain.InitialGenesis, err = codec.NewAnyWithValue(&types.DefaultInitialGenesis{})

		// DefaultInitialGenesis must always be able to be encoded into Any therefore this is a critical error
		if err != nil {
			return nil, spnerrors.Criticalf("DefaultInitialGenesis can't be used as initial genesis %s", err.Error())
		}
	} else {
		chain.InitialGenesis, err = codec.NewAnyWithValue(&types.GenesisURL{
			Url:  msg.GenesisURL,
			Hash: msg.GenesisHash,
		})
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
	}

	// Store values
	k.SetChain(ctx, chain)
	k.SetChainNameCount(ctx, chainNameCount)

	return &types.MsgCreateChainResponse{
		ChainID: chainID,
	}, nil
}
