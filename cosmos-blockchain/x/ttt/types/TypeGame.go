package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Game struct {
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
	ID       string         `json:"id" yaml:"id"`
	Opponent sdk.AccAddress `json:"opponent" yaml:"opponent"`
	Status   string         `json:"status" yaml:"status"`
	State    string         `json:"state" yaml:"state"`
	X        sdk.AccAddress `json:"X" yaml:"X"`
	O        sdk.AccAddress `json:"O" yaml:"O"`
	NextMove string         `json:"nextMove" yaml:"nextMove"`
	Winner   string         `json:"winner" yaml:"winner"`
}
