package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalid = sdkerrors.Register(ModuleName, 1, "Game does not exist")
	ErrIllegalAccess = sdkerrors.Register(ModuleName, 2, "Cannot join game that is not open")
	ErrIncorrectAccess = sdkerrors.Register(ModuleName, 3, "Cannot play in open or finished games")
	ErrPlayerOutOfTurn = sdkerrors.Register(ModuleName, 4, "Player out of turn")
	ErrMoveFormat = sdkerrors.Register(ModuleName, 5, "Incorrect move format: must be between 0 and 9")
	ErrIllegalMove = sdkerrors.Register(ModuleName, 6, "Illegal move: position already taken")
)
