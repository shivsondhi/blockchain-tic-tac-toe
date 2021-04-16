package ttt

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/shivsondhi/ttt/x/ttt/keeper"
	"github.com/shivsondhi/ttt/x/ttt/types"
)

// Handle a message to delete name
func handleMsgDeleteGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteGame) (*sdk.Result, error) {
	if !k.GameExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetGameOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteGame(ctx, msg.ID)
	return &sdk.Result{}, nil
}
