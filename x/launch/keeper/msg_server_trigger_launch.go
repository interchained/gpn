package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) TriggerLaunch(goCtx context.Context, msg *types.MsgTriggerLaunch) (*types.MsgTriggerLaunchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrChainNotFound, msg.ChainID)
	}

	// Check sender is the coordinator of the chain
	coordinatorID, found := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}
	if chain.CoordinatorID != coordinatorID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the chain is %v",
			chain.CoordinatorID,
		))
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrap(types.ErrTriggeredLaunch, msg.ChainID)
	}

	if msg.RemainingTime < k.MinLaunchTime(ctx) {
		return nil, sdkerrors.Wrapf(types.ErrLaunchTimeTooLow, "%v", msg.RemainingTime)
	}
	if msg.RemainingTime > k.MaxLaunchTime(ctx) {
		return nil, sdkerrors.Wrapf(types.ErrLaunchTimeTooHigh, "%v", msg.RemainingTime)
	}

	chain.LaunchTriggered = true
	toast := ctx.BlockTime().Unix()
	chain.LaunchTimestamp = toast + int64(msg.RemainingTime)
	k.SetChain(ctx, chain)

	return &types.MsgTriggerLaunchResponse{}, nil
}
