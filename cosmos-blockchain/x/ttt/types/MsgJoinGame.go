package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgJoinGame{}

type MsgJoinGame struct {
	ID       string         `json:"id" yaml:"id"`
	Opponent sdk.AccAddress `json:"opponent" yaml:"opponent"`
	Status   string         `json:"status" yaml:"status"`
}

func NewMsgJoinGame(id string, opponent sdk.AccAddress, status string) MsgJoinGame {
	return MsgJoinGame{
		ID:       id,
		Opponent: opponent,
		Status:   status,
	}
}

func (msg MsgJoinGame) Route() string {
	return RouterKey
}

func (msg MsgJoinGame) Type() string {
	return "JoinGame"
}

func (msg MsgJoinGame) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Opponent)}
}

func (msg MsgJoinGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgJoinGame) ValidateBasic() error {
	if msg.Opponent.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "opponent can't be empty")
	}
	return nil
}
