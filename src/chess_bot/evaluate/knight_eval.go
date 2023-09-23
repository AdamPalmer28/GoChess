package evaluate

import (
	"chess/chess_engine/board"
)

const (

	// optimized knight board
	mid_sq_val float64 = 0.07
	mid4_sq_val float64 = 0.05
	edge_sq_val float64 = -0.1



)

func KnightEval(cb board.ChessBoard, knight_rays *[64]board.Bitboard) float64 {

	var score float64 = 0.0

	// knight positions

		// white
	var white_attacks board.Bitboard
	for _, ind := range cb.WhiteKnights.Index() {
		white_attacks |= knight_rays[ind]
	}
	score += pos_capture_eval(*cb.WhiteKnights, white_attacks)

		// black
	var black_attacks board.Bitboard
	for _, ind := range cb.BlackKnights.Index() {
		black_attacks |= knight_rays[ind]
	}
	score -= pos_capture_eval(*cb.BlackKnights, black_attacks)
	
	return score
}


func pos_capture_eval(pieces board.Bitboard, attack_rays board.Bitboard) float64 {

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
	attack_on_mid4 := attack_rays & Mid4Centre
	if attack_on_mid4 != 0 {

		n = attack_on_mid4.Count()
		score += float64(n) * mid4_sq_val / 2.5

		attack_on_mid := attack_rays & Mid4Centre
		n = attack_on_mid.Count()
		score += float64(n) * mid_sq_val / 2.5

	}

	return score
}