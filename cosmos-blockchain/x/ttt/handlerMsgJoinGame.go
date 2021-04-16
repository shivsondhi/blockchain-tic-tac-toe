package ttt

import (
	"crypto/sha1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/shivsondhi/ttt/x/ttt/keeper"
	"github.com/shivsondhi/ttt/x/ttt/types"
)

func handleMsgJoinGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgJoinGame) (*sdk.Result, error) {
	// Checks
	// check if game exists
	if !k.GameExists(ctx, msg.ID) {
		return nil, types.ErrInvalid
	}
	// check if game status is open
	if !k.GameOpen(ctx, msg.ID) {
		return nil, types.ErrIllegalAccess
	}

	// Hash and concatenate public keys of opponent and creator. Then set X and O.
	x := msg.Opponent
	o := msg.Opponent
	// creator's hashed pubkey
	creatorPubKey, _ := k.AccKeeper.GetPubKey(ctx, k.GetGameOwner(ctx, msg.ID))
	creatorPKStr, err := sdk.Bech32ifyPubKey("cosmospub", creatorPubKey)
	if err != nil {
		panic(err)
	}
	// opponent's hashed pubkey
	opponentPubKey, _ := k.AccKeeper.GetPubKey(ctx, msg.Opponent)
	opponentPKStr, err := sdk.Bech32ifyPubKey("cosmospub", opponentPubKey)
	if err != nil {
		panic(err)
	}
	// concatenate strings and hash
	concatStr := creatorPKStr + opponentPKStr
	hasher := sha1.New()
	hasher.Write([]byte(concatStr))
	hash := hasher.Sum(nil)
	// check first bit in byte array to set X and O
	if string(hash[0]) == "0" {
		o = k.GetGameOwner(ctx, msg.ID)
	} else {
		x = k.GetGameOwner(ctx, msg.ID)
	}


	var game = types.Game{
		Creator:  k.GetGameOwner(ctx, msg.ID),
		ID:       msg.ID,
		Opponent: msg.Opponent,
		Status:   msg.Status,
		State:    "         ",
		X:        x,
		O:        o,
		NextMove: "X",
		Winner:   "null",
	}

	k.SetGame(ctx, game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
