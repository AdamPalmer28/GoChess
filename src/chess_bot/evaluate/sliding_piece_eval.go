package evaluate

import (
	"chess/chess_engine/board"
)

/*
Evaluation:
	All sliding pieces:
		- mobility
		- ray/xray attacks
		- ray/xray attacks on valuable pieces
	Rook
	Bishop
	Queen


*/

const (
	// rook vals
	rk_mv_count_val float64 = 0.015
	rk_ray_xray_ratio float64 = 0.2

	rk_opp_q_val float64 = 0.12
	rk_opp_b_val float64 = 0.08
	rk_opp_r_val float64 = 0.00
	rk_opp_kn_val float64 = 0.05


	// bishop vals
	b_mv_count_val float64 = 0.014
	b_ray_xray_ratio float64 = 0.4

	b_opp_q_val float64 = 0.12
	b_opp_b_val float64 = 0.00
	b_opp_r_val float64 = 0.08
	b_opp_kn_val float64 = 0.06


	// queen vals
	q_mv_count_val float64 = 0.01
	q_ray_xray_ratio float64 = 0.3

	q_opp_q_val float64 = 0.0
	q_opp_b_val float64 = 0.02
	q_opp_r_val float64 = 0.02
	q_opp_kn_val float64 = 0.05
)

// Rook eval
// -----------------
func RookEval(eval_move *EvalMoveRays, cb board.ChessBoard) float64 {
	
	var score float64 = 0.0

	var ind int
	var ray, xray board.Bitboard
	var opp_q_atks, opp_b_atks, opp_rk_atks, opp_kn_atks board.Bitboard

	// evaluate move rays + xrays

	for ind, ray = range eval_move.W_r_rays { // white
		xray = eval_move.W_r_xrays[ind]

		opp_q_atks = xray & *cb.BlackQueens
		opp_b_atks = xray & *cb.BlackBishops
		opp_rk_atks = xray & *cb.BlackRooks
		opp_kn_atks = xray & *cb.BlackKnights
		
		score += eval_sliding(ray, opp_q_atks, opp_b_atks, opp_rk_atks, opp_kn_atks,
			rk_mv_count_val, rk_opp_q_val, rk_opp_b_val, rk_opp_r_val, rk_opp_kn_val, rk_ray_xray_ratio)
	}

	for ind, ray = range eval_move.B_r_rays { // black
		xray = eval_move.B_r_xrays[ind]

		opp_q_atks = xray & *cb.WhiteQueens
		opp_b_atks = xray & *cb.WhiteBishops
		opp_rk_atks = xray & *cb.WhiteRooks
		opp_kn_atks = xray & *cb.WhiteKnights
		

		score -= eval_sliding(ray, opp_q_atks, opp_b_atks, opp_rk_atks, opp_kn_atks,
			rk_mv_count_val, rk_opp_q_val, rk_opp_b_val, rk_opp_r_val, rk_opp_kn_val, rk_ray_xray_ratio)
	}
	
	return score
}


// Bishop eval
// -----------------
func BishopEval(eval_move *EvalMoveRays, cb board.ChessBoard) float64 {

	var score float64 = 0.0

	var ind int
	var ray, xray board.Bitboard
	var opp_q_atks, opp_b_atks, opp_rk_atks, opp_kn_atks board.Bitboard

	// evaluate move rays + xrays

	for ind, ray = range eval_move.W_b_rays { // white
		xray = eval_move.W_b_xrays[ind]

		opp_q_atks = xray & *cb.BlackQueens
		opp_b_atks = xray & *cb.BlackBishops
		opp_rk_atks = xray & *cb.BlackRooks
		opp_kn_atks = xray & *cb.BlackKnights


		score += eval_sliding(ray, opp_q_atks, opp_b_atks, opp_rk_atks, opp_kn_atks,
			b_mv_count_val, b_opp_q_val, b_opp_b_val, b_opp_r_val, b_opp_kn_val, b_ray_xray_ratio)
	}

	for ind, ray = range eval_move.B_b_rays { // black
		xray = eval_move.B_b_xrays[ind]

		opp_q_atks = xray & *cb.WhiteQueens
		opp_b_atks = xray & *cb.WhiteBishops
		opp_rk_atks = xray & *cb.WhiteRooks
		opp_kn_atks = xray & *cb.WhiteKnights


		score -= eval_sliding(ray, opp_q_atks, opp_b_atks, opp_rk_atks, opp_kn_atks,
			b_mv_count_val, b_opp_q_val, b_opp_b_val, b_opp_r_val, b_opp_kn_val, b_ray_xray_ratio)
	}

	return score
}


// Queen eval
// -----------------
func QueenEval(eval_move *EvalMoveRays, cb board.ChessBoard) float64 {
	
	var score float64 = 0.0

	var ind int
	var ray, xray board.Bitboard
	var opp_q_atks, opp_b_atks, opp_rk_atks, opp_kn_atks board.Bitboard

	// evaluate move rays + xrays

	for ind, ray = range eval_move.W_q_rays { // white
		xray = eval_move.W_q_xrays[ind]

		opp_q_atks = xray & *cb.BlackQueens
		opp_b_atks = xray & *cb.BlackBishops
		opp_rk_atks = xray & *cb.BlackRooks
		opp_kn_atks = xray & *cb.BlackKnights


		score += eval_sliding(ray, opp_q_atks, opp_b_atks, opp_rk_atks, opp_kn_atks,
			q_mv_count_val, q_opp_q_val, q_opp_b_val, q_opp_r_val, q_opp_kn_val, q_ray_xray_ratio)
	}

	for ind, ray = range eval_move.B_q_rays { // black
		xray = eval_move.B_q_xrays[ind]

		opp_q_atks = xray & *cb.WhiteQueens
		opp_b_atks = xray & *cb.WhiteBishops
		opp_rk_atks = xray & *cb.WhiteRooks
		opp_kn_atks = xray & *cb.WhiteKnights

		score -= eval_sliding(ray, opp_q_atks, opp_b_atks, opp_rk_atks, opp_kn_atks,
			q_mv_count_val, q_opp_q_val, q_opp_b_val, q_opp_r_val, q_opp_kn_val, q_ray_xray_ratio)
	}

	return score
}


// ============================================================================

// eval_sliding - general sliding piece evaluation
func eval_sliding(rays, opp_q_xray, opp_b_xray, opp_rk_xray, opp_kn_xray board.Bitboard,
		move_cnt_val, opp_q_val, opp_b_val, opp_rk_val, opp_kn_val, ray_xray_ratio float64) float64 {

	var score float64 = 0.0

	score += float64(rays.Count()) * move_cnt_val // mobility

	// ray/xray attacks on valuable pieces

	// queen
	if (opp_q_xray) != 0 {
		// queen is on full move rays
		score += opp_q_val * float64(opp_q_xray.Count()) * ray_xray_ratio
		// queen is on direct move rays
		score += opp_q_val * float64((rays & opp_q_xray).Count()) * (1- ray_xray_ratio)
	}
	// rook
	if (opp_rk_xray) != 0 {
		// rook is on full move rays
		score += opp_rk_val * float64(opp_rk_xray.Count()) * ray_xray_ratio
		// rook is on direct move rays
		score += opp_rk_val * float64((rays & opp_rk_xray).Count()) * (1- ray_xray_ratio)
	}
	// bishop
	if (opp_b_xray) != 0 {
		// bishop is on full move rays
		score += opp_b_val * float64(opp_b_xray.Count()) * ray_xray_ratio
		// bishop is on direct move rays
		score += opp_b_val * float64((rays & opp_b_xray).Count()) * (1- ray_xray_ratio)
	}
	// knight
	if (opp_kn_xray) != 0 {
		// knight is on full move rays
		score += opp_kn_val * float64(opp_kn_xray.Count()) * ray_xray_ratio
		// knight is on direct move rays
		score += opp_kn_val * float64((rays & opp_kn_xray).Count()) * (1- ray_xray_ratio)
	}

	return score
}
