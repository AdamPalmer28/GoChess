package chess_bot

import (
	"chess/src/chess_bot/evaluate"
	"chess/src/chess_engine/gamestate"
)

// Evaluate the board
func Evaluate(gs *gamestate.GameState) float64 {

	cb := gs.Board
	
	// Pre requisite data
	EvalMoves := evaluate.GetEvalMoveRays(gs)
	//BoardActivity := evaluate.GetBoardActivity(gs, EvalMoves) // function not implemented yet
	
	// get scores
	var score float64 = 0.0
	// piece total
	score += evaluate.EvalPieceCounts(cb)
	
	// pawn eval
	PawnScore := evaluate.PawnEval(cb, &gs.MoveRays.PawnCapRays)
	score += PawnScore[0] - PawnScore[1]
	
	// knight eval 
	KnightScore := evaluate.KnightEval(cb, EvalMoves.W_kn_rays, EvalMoves.B_kn_rays)
	score += KnightScore[0] - KnightScore[1]
	
	// bishop eval
	BishopScore := evaluate.BishopEval(cb, &EvalMoves)
	score += BishopScore[0] - BishopScore[1]
	
	// rook eval
	RookScore := evaluate.RookEval(cb, &EvalMoves)
	score += RookScore[0] - RookScore[1]
	
	// queen eval
	QueenScore := evaluate.QueenEval(cb, &EvalMoves)
	score += QueenScore[0] - QueenScore[1]
	
	// king eval
	KingScore := evaluate.KingSafetyEval(cb, &EvalMoves)
	score += KingScore[0] - KingScore[1]
	
	
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
	Total float64  `json:"total"`

	PieceValue float64 `json:"Piece Value"`

	PawnEval [2]float64 `json:"Pawn"`
	KnightEval [2]float64 `json:"Knight"`
	BishopEval [2]float64 `json:"Bishop"`
	RookEval [2]float64 `json:"Rook"`
	QueenEval [2]float64 `json:"Queen"`
	KingSafety [2]float64 `json:"King"`

}

func EvalScoreData(gs *gamestate.GameState) EvalScore {
	// gets eval data manually for the website output
	var eval EvalScore

	eval.Total = Evaluate(gs)

	cb := gs.Board
	EvalMoves := evaluate.GetEvalMoveRays(gs)

	eval.PieceValue = evaluate.EvalPieceCounts(cb)

	eval.PawnEval = evaluate.PawnEval(cb, &gs.MoveRays.PawnCapRays)
	eval.KnightEval = evaluate.KnightEval(cb, EvalMoves.W_kn_rays, EvalMoves.B_kn_rays)
	eval.BishopEval = evaluate.BishopEval(cb, &EvalMoves)
	eval.RookEval = evaluate.RookEval(cb, &EvalMoves)
	eval.QueenEval = evaluate.QueenEval(cb, &EvalMoves)
	eval.KingSafety = evaluate.KingSafetyEval(cb, &EvalMoves)

	return eval
}