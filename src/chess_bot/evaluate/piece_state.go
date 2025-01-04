package evaluate

import (
	"chess/src/chess_engine/board"
)

const (
	pawn_val float64 = 1.0
	knight_val float64 = 3.0
	bishop_val float64 = 3.0
	rook_val float64 = 5.0
	queen_val float64 = 9.0
)

// piece total
func EvalPieceCounts(cb board.ChessBoard) float64 {

	score := 0.0

	// pawns
	score += pawn_val * float64(cb.WhitePawns.Count() - cb.BlackPawns.Count())

	// knights
	score += knight_val * float64(cb.WhiteKnights.Count() - cb.BlackKnights.Count())

	// bishops
	score += bishop_val * float64(cb.WhiteBishops.Count() - cb.BlackBishops.Count())

	// rooks
	score += rook_val * float64(cb.WhiteRooks.Count() - cb.BlackRooks.Count())

	// queens
	score += queen_val * float64(cb.WhiteQueens.Count() - cb.BlackQueens.Count())

	return score
}