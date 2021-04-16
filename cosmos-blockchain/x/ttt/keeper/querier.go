package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/shivsondhi/ttt/x/ttt/types"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for ttt clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryListGame:
			return listGame(ctx, k)
		case types.QueryListShort:
			return listShort(ctx, k)
		case types.QueryGetGame:
			return getGame(ctx, path[1:], k)
		case types.QueryGetBoard:
			return getBoard(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown ttt query endpoint")
		}
	}
}
