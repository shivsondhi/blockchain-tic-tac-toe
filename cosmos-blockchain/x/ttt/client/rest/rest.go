package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers ttt-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/ttt/game", createGameHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/ttt/game", listGameHandler(cliCtx, "ttt")).Methods("GET")
	r.HandleFunc("/ttt/game/{key}", getGameHandler(cliCtx, "ttt")).Methods("GET")
	r.HandleFunc("ttt/game/short", listShortHandler(cliCtx, "ttt")).Methods("GET")
	r.HandleFunc("ttt/game/{key}/board", getBoardHandler(cliCtx, "ttt")).Methods("GET")
	r.HandleFunc("/ttt/game", joinGameHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/ttt/game", playMoveHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/ttt/game", deleteGameHandler(cliCtx)).Methods("DELETE")

}
