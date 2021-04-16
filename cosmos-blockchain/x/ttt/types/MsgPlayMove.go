package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPlayMove{}

type MsgPlayMove struct {
	ID       string         `json:"id" yaml:"id"`
	Move     string         `json:"move" yaml:"move"`
	Player   sdk.AccAddress `json:"player" yaml:"player"`
}

func NewMsgPlayMove(id string, move string, player sdk.AccAddress) MsgPlayMove {
	return MsgPlayMove{
		ID:       id,
		Move: move,
		Player: player,
	}
}

func (msg MsgPlayMove) Route() string {
	return RouterKey
}

func (msg MsgPlayMove) Type() string {
	return "PlayMove"
}

func (msg MsgPlayMove) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Player)}
}

func (msg MsgPlayMove) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgPlayMove) ValidateBasic() error {
	if msg.Player.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "player can't be empty")
	}
	return nil
}
