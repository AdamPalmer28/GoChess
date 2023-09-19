package chess_bot

import (
	"chess/chess_bot/evaluate"
	"chess/chess_engine/gamestate"
)

// Evaluate the board
func Evaluate(gs *gamestate.GameState) float64 {

	cb := gs.Board

	score := 0.0

	// piece total
	score += evaluate.EvalPieceCounts(cb)

	// pawn eval

	// knight eval

	// bishop eval

	// rook eval

	// queen eval

	// king eval

	return score
}
