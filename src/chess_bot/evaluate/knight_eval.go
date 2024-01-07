package evaluate

import (
	"chess/src/chess_engine/board"
)

const (

	// optimized knight board
	mid_sq_val float64 = 0.02
	mid4_sq_val float64 = 0.05
	edge_sq_val float64 = -0.1



)

func KnightEval(cb board.ChessBoard, white_attacks []board.Bitboard, black_attacks []board.Bitboard) float64 {

	var score float64 = 0.0

	// knight positions

		// white
	score += pos_capture_eval(*cb.WhiteKnights, white_attacks)

		// black
	score -= pos_capture_eval(*cb.BlackKnights, black_attacks)
	
	return score
}


func pos_capture_eval(pieces board.Bitboard, attack_rays []board.Bitboard) float64 {

	var score float64 = 0.0
	var n int

	// piece on middle squares
	piece_on_mid4 := pieces & Mid4Centre
	if piece_on_mid4 != 0 {

		n = piece_on_mid4.Count()
		score += float64(n) * mid4_sq_val

		piece_on_mid := pieces & Mid4Centre
		n = piece_on_mid.Count()
		score += float64(n) * mid_sq_val

	} else {
		// piece on edge
		piece_on_edge := pieces & Edge
		n = piece_on_edge.Count()
		score += float64(n) * edge_sq_val
	}

	// attack rays on middle squares
	for _, ray := range attack_rays {
		attack_on_mid4 := ray & Mid4Centre
		if attack_on_mid4 != 0 {

			n = attack_on_mid4.Count()
			score += float64(n) * mid4_sq_val / 2.5

			attack_on_mid := ray & Mid4Centre
			n = attack_on_mid.Count()
			score += float64(n) * mid_sq_val / 2.5

		}
	}

	return score
}