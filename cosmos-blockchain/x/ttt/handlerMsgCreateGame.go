package ttt

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/shivsondhi/ttt/x/ttt/keeper"
	"github.com/shivsondhi/ttt/x/ttt/types"
)

func handleMsgCreateGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateGame) (*sdk.Result, error) {
	k.CreateGame(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
