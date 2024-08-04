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
	PawnScore := evaluate.PawnEval(cb, &gs.MoveRays.PawnCapRays)
	score += PawnScore[0] - PawnScore[1]
	
	// knight eval 
	KnightScore := evaluate.KnightEval(cb, EvalMoves.W_kn_rays, EvalMoves.B_kn_rays)
	score += KnightScore[0] - KnightScore[1]
	
	// bishop eval
	BishopScore := evaluate.BishopEval(&EvalMoves, cb)
	score += BishopScore[0] - BishopScore[1]
	
	// rook eval
	RookScore := evaluate.RookEval(&EvalMoves, cb)
	score += RookScore[0] - RookScore[1]
	
	// queen eval
	QueenScore := evaluate.QueenEval(&EvalMoves, cb)
	score += QueenScore[0] - QueenScore[1]
	
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


	PawnEval [2]float64
	KnightEval [2]float64
	BishopEval [2]float64
	RookEval [2]float64
	QueenEval [2]float64
	KingSafety [2]float64

}

func EvalScoreData(gs *gamestate.GameState) EvalScore {
	var eval EvalScore

	eval.Total = Evaluate(gs)

	cb := gs.Board
	EvalMoves := evaluate.GetEvalMoveRays(gs)

	eval.PawnEval = evaluate.PawnEval(cb, &gs.MoveRays.PawnCapRays)
	eval.KnightEval = evaluate.KnightEval(cb, EvalMoves.W_kn_rays, EvalMoves.B_kn_rays)
	eval.BishopEval = evaluate.BishopEval(&EvalMoves, cb)
	eval.RookEval = evaluate.RookEval(&EvalMoves, cb)
	eval.QueenEval = evaluate.QueenEval(&EvalMoves, cb)

	
	return eval
}