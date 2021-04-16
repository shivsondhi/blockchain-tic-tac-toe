package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgCreateGame{}, "ttt/CreateGame", nil)
	cdc.RegisterConcrete(MsgJoinGame{}, "ttt/JoinGame", nil)
	cdc.RegisterConcrete(MsgPlayMove{}, "ttt/PlayMove", nil)
	cdc.RegisterConcrete(MsgDeleteGame{}, "ttt/DeleteGame", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
