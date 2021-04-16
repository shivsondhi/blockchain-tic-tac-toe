package ttt

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/shivsondhi/ttt/x/ttt/keeper"
	"github.com/shivsondhi/ttt/x/ttt/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case types.MsgCreateGame:
			return handleMsgCreateGame(ctx, k, msg)
		case types.MsgJoinGame:
			return handleMsgJoinGame(ctx, k, msg)
		case types.MsgPlayMove:
			return handleMsgPlayMove(ctx, k, msg)
		case types.MsgDeleteGame:
			return handleMsgDeleteGame(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
