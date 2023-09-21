package test_gamestate

import (
	"chess/chess_engine/gamestate"
	"testing"
)

/*
Moves and undo to test:
	- simple quiet move
	- simple capture
	- enpassant
	- promotion
	- castle
*/

func TestUndo(t *testing.T) {


	// ------------------------------------------------------------------------
	// standard move

	gs := gamestate.MakeGameState()
	gs.Init()

	cb := gs.Board.Copy()

	// make a move
	move := gs.MoveList[0]
	gs.Make_move(move)

	// undo the move
	gs.Undo()

	// check if the board is the same
	if !cb.Identical(gs.Board) {
		t.Errorf("1. Board not the same after undo")
		cb.Print()
		gs.Board.Print()
	}

	// ------------------------------------------------------------------------
	// Capture move


	// ------------------------------------------------------------------------
	// Enpassant Capture


	// Move 

}