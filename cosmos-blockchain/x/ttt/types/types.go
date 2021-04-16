package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GameBoard struct {
	Row123    string `json:"1 2 3" yaml:"1 2 3"`
	Row456    string `json:"4 5 6" yaml:"4 5 6"`
	Row789    string `json:"7 8 9" yaml:"7 8 9"`
	NextMove  string   `json:"nextMove" yaml:"nextMove"`
}

type ShortGame struct {
	ID      string          `json:"id" yaml:"id"`
	Creator sdk.AccAddress  `json:"creator" yaml:"creator"`
	Status  string          `json:"status" yaml:"status"`
	Winner  string          `json:"winner" yaml:"winner"`
}
