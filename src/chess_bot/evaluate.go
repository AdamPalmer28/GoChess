package chess_bot

import (
	"chess/chess_bot/evaluate"
	"chess/chess_engine/gamestate"
)

// Evaluate the board
func Evaluate(gs *gamestate.GameState) float64 {

	cb := gs.Board

	var score_scalar float64 = 1.0
	if !gs.White_to_move {
		score_scalar = -1.0
	}

	var score float64 = 0.0

	// piece total
	score += evaluate.EvalPieceCounts(cb)

	// pawn eval
	score += evaluate.PawnEval(cb, &gs.MoveRays.PawnCapRays)

	// knight eval
	score += evaluate.KnightEval(cb, &gs.MoveRays.KnightRays)

	// bishop eval

	// rook eval

	// queen eval

	// king eval



	// general
	score += score_scalar * float64(len(gs.MoveList)) / 50 // number of moves
	if gs.InCheck {
		score += score_scalar * 0.2 // check
	}

	return score
}
