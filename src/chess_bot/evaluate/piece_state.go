package evaluate

import (
	"chess/src/chess_engine/board"
)

// piece total
func EvalPieceCounts(cb board.ChessBoard) float64 {

	score := 0.0

	// pawns
	score += 1.0 * float64(cb.WhitePawns.Count() - cb.BlackPawns.Count())

	// knights
	score += 3.0 * float64(cb.WhiteKnights.Count() - cb.BlackKnights.Count())

	// bishops
	score += 3.0 * float64(cb.WhiteBishops.Count() - cb.BlackBishops.Count())

	// rooks
	score += 5.0 * float64(cb.WhiteRooks.Count() - cb.BlackRooks.Count())

	// queens
	score += 9.0 * float64(cb.WhiteQueens.Count() - cb.BlackQueens.Count())

	return score
}