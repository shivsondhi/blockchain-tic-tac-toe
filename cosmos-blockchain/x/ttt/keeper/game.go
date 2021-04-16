package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/shivsondhi/ttt/x/ttt/types"
)

// GetGameCount get the total number of game
func (k Keeper) GetGameCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.GameCountPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetGameCount set the total number of game
func (k Keeper) SetGameCount(ctx sdk.Context, count int64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.GameCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateGame creates a game
func (k Keeper) CreateGame(ctx sdk.Context, msg types.MsgCreateGame) {
	// Create the game
	count := k.GetGameCount(ctx)
	var game = types.Game{
		Creator:  msg.Creator,
		ID:       strconv.FormatInt(count, 10),
		Opponent: msg.Opponent,
		Status:   msg.Status,
		State:    msg.State,
		X:        msg.X,
		O:        msg.O,
		NextMove: msg.NextMove,
		Winner:   msg.Winner,
	}

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.GamePrefix + game.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(game)
	store.Set(key, value)

	// Update game count
	k.SetGameCount(ctx, count+1)
}

// GetGame returns the game information
func (k Keeper) GetGame(ctx sdk.Context, key string) (types.Game, error) {
	store := ctx.KVStore(k.storeKey)
	var game types.Game
	byteKey := []byte(types.GamePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &game)
	if err != nil {
		return game, err
	}
	return game, nil
}

// SetGame sets the game state
func (k Keeper) SetGame(ctx sdk.Context, game types.Game) {
	gameKey := game.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(game)
	key := []byte(types.GamePrefix + gameKey)
	store.Set(key, bz)
}


// DeleteGame deletes a game
func (k Keeper) DeleteGame(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.GamePrefix + key))
}

//
// Functions used by querier
//

func listGame(ctx sdk.Context, k Keeper) ([]byte, error) {
	var gameList []types.Game
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.GamePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var game types.Game
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &game)
		gameList = append(gameList, game)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, gameList)
	return res, nil
}

func listShort(ctx sdk.Context, k Keeper) ([]byte, error) {
	var gameList []types.ShortGame
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.GamePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var game types.Game
		var shortGame types.ShortGame
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &game)
		shortGame.Creator = game.Creator
		shortGame.ID = game.ID
		shortGame.Status = game.Status
		shortGame.Winner = game.Winner
		gameList = append(gameList, shortGame)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, gameList)
	return res, nil
}

func getGame(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	game, err := k.GetGame(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, game)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func getBoard(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	game, err := k.GetGame(ctx, key)
	if err != nil {
		return nil, err
	}
	var board types.GameBoard
	board.Row123 = string(game.State[:3])
	board.Row456 = string(game.State[3:6])
	board.Row789 = string(game.State[6:])
	board.NextMove = game.NextMove

	res, err = codec.MarshalJSONIndent(k.cdc, board)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetGameOwner(ctx sdk.Context, key string) sdk.AccAddress {
	game, err := k.GetGame(ctx, key)
	if err != nil {
		return nil
	}
	return game.Creator
}

// Get if status of game is open
func (k Keeper) GameOpen(ctx sdk.Context, key string) bool {
	game, err := k.GetGame(ctx, key)
	if err != nil {
		return false
	}
	return (game.Status == "open")
}

// Check if the key exists in the store
func (k Keeper) GameExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.GamePrefix + key))
}
