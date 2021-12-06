package ttt

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/shivsondhi/ttt/x/ttt/keeper"
	"github.com/shivsondhi/ttt/x/ttt/types"
)

func handleMsgPlayMove(ctx sdk.Context, k keeper.Keeper, msg types.MsgPlayMove) (*sdk.Result, error) {
	currMove := ""
	nextMove := ""
	newState := ""
	currPlayer := msg.Player //temporarily

	// Check if game exists
	if !k.GameExists(ctx, msg.ID) {
		return nil, types.ErrInvalid
	}

	// Get game state
	gameState, err := k.GetGame(ctx, msg.ID)
	if err != nil {
		return nil, err
	}
	status := gameState.Status

	// Checks
	// check if game is inprogress
	if gameState.Status != "inprogress" {
		return nil, types.ErrIncorrectAccess
	}

	// get player to play
	if gameState.NextMove == "X" {
		currPlayer = gameState.X
		currMove = "X"
		nextMove = "O"
	} else {
		currPlayer = gameState.O
		currMove = "O"
		nextMove = "X"
	}
	// check player out of turn
	if !msg.Player.Equals(currPlayer) {
		return nil, types.ErrPlayerOutOfTurn
	}

	// check move format
	pos, err := strconv.Atoi(msg.Move)
	if err != nil {
		return nil, err
	}
	if (pos < 1) || (pos > 9) {
		return nil, types.ErrMoveFormat
	}

	// check move validity and update
	if string(gameState.State[pos-1]) != " " {
		return nil, types.ErrIllegalMove
	} else {
		newState = gameState.State[:pos-1] + string(currMove) + gameState.State[pos:]
	}

	// check if game is over
	full := true
	finished := false
	winner := "null"
	for i, letter := range newState {
		if string(letter) == " " {
			full = false
			continue
		}
		switch i {
		case 0:
			// if i==0 : check horizontal, check vertical and check diagonal
			if ((string(letter) == string(newState[i+1])) && (string(letter) == string(newState[i+2]))) || ((string(letter) == string(newState[i+3])) && (string(letter) == string(newState[i+6]))) || ((string(letter) == string(newState[i+4])) && (string(letter) == string(newState[i+8]))) {
				finished = true
				winner = string(letter)
			}
		case 1:
			// if i==1 : check vertical
			if (string(letter) == string(newState[i+3])) && (string(letter) == string(newState[i+6])) {
				finished = true
				winner = string(letter)
			}
		case 2:
			// if i==2 : check vertical and diagonal
			if ((string(letter) == string(newState[i+3])) && (string(letter) == string(newState[i+6]))) || ((string(letter) == string(newState[i+2])) && (string(letter) == string(newState[i+4]))) {
				finished = true
				winner = string(letter)
			}
		case 3:
			// if i==3 : check horizontal
			if (string(letter) == string(newState[i+1])) && (string(letter) == string(newState[i+2])) {
				finished = true
				winner = string(letter)
			}
		case 6:
			// if i==6 : check horizontal
			if (string(letter) == string(newState[i+1])) && (string(letter) == string(newState[i+2])) {
				finished = true
				winner = string(letter)
			}
		}
	}
	if full {
		finished = true
		if string(winner) == "" {
			winner = "draw"
		}
	}

	if finished {
		status = "finished"
	}

	var game = types.Game{
		Creator:  gameState.Creator,
		ID:       msg.ID,
		Opponent: gameState.Opponent,
		Status:   status,
		State:    newState,
		X:        gameState.X,
		O:        gameState.O,
		NextMove: nextMove,
		Winner:   winner,
	}

	k.SetGame(ctx, game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
