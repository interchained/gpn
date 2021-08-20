package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) RequestAddValidator(
	goCtx context.Context,
	msg *types.MsgRequestAddValidator,
) (*types.MsgRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrChainNotFound, msg.ChainID)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrap(types.ErrTriggeredLaunch, msg.ChainID)
	}

	coordAddress, found := k.profileKeeper.GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainInactive,
			"the chain %s coordinator has been deleted", chain.ChainID)
	}

	content := types.NewGenesisValidator(
		msg.ChainID,
		msg.ValAddress,
		msg.GenTx,
		msg.ConsPubKey,
		msg.SelfDelegation,
		msg.Peer,
		)
	request := types.Request{
		ChainID:   msg.ChainID,
		Creator:   msg.ValAddress,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   content,
	}

	var requestID uint64
	approved := false
	if msg.ValAddress == coordAddress {
		err := ApplyRequest(ctx, k.Keeper, msg.ChainID, request)
		if err != nil {
			return nil, err
		}
		approved = true
	} else {
		requestID = k.AppendRequest(ctx, request)
	}

	return &types.MsgRequestResponse{
		RequestID:    requestID,
		AutoApproved: approved,
	}, nil
}
