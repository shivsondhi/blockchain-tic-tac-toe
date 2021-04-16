package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGame{}

type MsgCreateGame struct {
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
	Opponent sdk.AccAddress `json:"opponent" yaml:"opponent"`
	Status   string         `json:"status" yaml:"status"`
	State    string         `json:"state" yaml:"state"`
	X        sdk.AccAddress `json:"X" yaml:"X"`
	O        sdk.AccAddress `json:"O" yaml:"O"`
	NextMove string         `json:"nextMove" yaml:"nextMove"`
	Winner   string         `json:"winner" yaml:"winner"`
}

func NewMsgCreateGame(creator sdk.AccAddress, opponent sdk.AccAddress, status string, state string, X sdk.AccAddress, O sdk.AccAddress, nextMove string, winner string) MsgCreateGame {
	return MsgCreateGame{
		Creator:  creator,
		Opponent: opponent,
		Status:   status,
		State:    state,
		X:        X,
		O:        O,
		NextMove: nextMove,
		Winner:   winner,
	}
}

func (msg MsgCreateGame) Route() string {
	return RouterKey
}

func (msg MsgCreateGame) Type() string {
	return "CreateGame"
}

func (msg MsgCreateGame) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateGame) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
