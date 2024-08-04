package chess_bot

import (
	"chess/src/chess_bot/evaluate"
	"chess/src/chess_engine/gamestate"
)


// Evaluate the board
func Evaluate(gs *gamestate.GameState) float64 {

	cb := gs.Board
	
	//MoveRAys
	var score float64 = 0.0
	
	EvalMoves := evaluate.GetEvalMoveRays(gs)
	
	// piece total
	score += evaluate.EvalPieceCounts(cb)
	
	// pawn eval
	score += evaluate.PawnEval(cb, &gs.MoveRays.PawnCapRays)
	
	// knight eval 
	score += evaluate.KnightEval(cb, EvalMoves.W_kn_rays, EvalMoves.B_kn_rays)
	
	// bishop eval
	score += evaluate.BishopEval(&EvalMoves, cb)
	
	// rook eval
	score += evaluate.RookEval(&EvalMoves, cb)
	
	// queen eval
	score += evaluate.QueenEval(&EvalMoves, cb)
	
	// king eval
	
	
	
	var score_scalar float64 = 1.0
	if !gs.White_to_move {
		score_scalar = -1.0
	}
	// general
	// score += score_scalar * float64(len(gs.MoveList)) / 50 // number of moves
	if gs.InCheck {
		score += score_scalar * 0.2 // check
	}

	return score
}


// ============================================================================
// Eval score for analysis

type EvalScore struct {
	Total float64

	playerScore [2]float64

	PawnEval [2]float64
	KnightEval [2]float64
	BishopEval [2]float64
	RookEval [2]float64
	QueenEval [2]float64
	KingSafety [2]float64

}

func EvalScoreData(gs *gamestate.GameState) EvalScore {
	var eval EvalScore

	cb := gs.Board
	
	//MoveRAys
	var score float64 = 0.0
	
	EvalMoves := evaluate.GetEvalMoveRays(gs)