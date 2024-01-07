package evaluate

import (
	"chess/src/chess_engine/board"
	"chess/src/chess_engine/gamestate"
	"chess/src/chess_engine/move_gen/magic"
)

/*
Generate all move rays for each piece for use in evaluation
It includes all direct move rays and xrays

*/

type EvalMoveRays struct {

	// knight rays
	W_kn_rays []board.Bitboard
	B_kn_rays []board.Bitboard
	
	
	// bishop rays
	W_b_rays []board.Bitboard
	B_b_rays []board.Bitboard
	// ignoring other pieces
	W_b_xrays []board.Bitboard
	B_b_xrays []board.Bitboard
	
	
	// rook rays
	W_r_rays []board.Bitboard
	B_r_rays []board.Bitboard
	// ignoring other pieces
	W_r_xrays []board.Bitboard
	B_r_xrays []board.Bitboard
	
	
	// queen rays
	W_q_rays []board.Bitboard
	B_q_rays []board.Bitboard
	// ignoring other pieces
	W_q_xrays []board.Bitboard
	B_q_xrays []board.Bitboard


	// king rays - (sliding, knight, pawn)
	W_king_rays Threats 
	B_king_rays Threats
		// ignoring other pieces (only sliding)
	W_king_xrays XrayThreats
	B_king_xrays XrayThreats
}

type Threats struct { // threats to king
	Straight board.Bitboard
	Diag board.Bitboard
	Knight board.Bitboard
	Pawn board.Bitboard
}
type XrayThreats struct { // xray threats to king
	Straight board.Bitboard
	Diag board.Bitboard
}



func GetEvalMoveRays(gs *gamestate.GameState) EvalMoveRays {
	
	var eval_move_rays EvalMoveRays
	
	occ_board := gs.Board.White | gs.Board.Black
	
	// ------------------------------------------------------------------------
	// sliding pieces

	// bishop rays
	eval_move_rays.W_b_rays, eval_move_rays.W_b_xrays = getMoveRay(*gs.Board.WhiteBishops, occ_board,
		&gs.MoveRays.Magic.BishopMagic, &gs.MoveRays.BishopRays)
	eval_move_rays.B_b_rays, eval_move_rays.B_b_xrays = getMoveRay(*gs.Board.BlackBishops, occ_board,
		&gs.MoveRays.Magic.BishopMagic, &gs.MoveRays.BishopRays)


	// rook rays
	eval_move_rays.W_r_rays, eval_move_rays.W_r_xrays = getMoveRay(*gs.Board.WhiteRooks, occ_board,
		&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.RookRays)
	eval_move_rays.B_r_rays, eval_move_rays.B_r_xrays = getMoveRay(*gs.Board.BlackRooks, occ_board,
		&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.RookRays)
	

	// queen rays
		// white
	queen_diag_rays, queen_diag_xrays := getMoveRay(*gs.Board.WhiteQueens, occ_board,
		&gs.MoveRays.Magic.BishopMagic, &gs.MoveRays.BishopRays)
	queen_straight_rays, queen_straight_xrays  := getMoveRay(*gs.Board.WhiteQueens, occ_board,
		&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.RookRays)

	eval_move_rays.W_q_rays = combineRaysSets(queen_diag_rays, queen_straight_rays)
	eval_move_rays.W_q_xrays = combineRaysSets(queen_diag_xrays, queen_straight_xrays)

		// black
	queen_diag_rays, queen_diag_xrays = getMoveRay(*gs.Board.BlackQueens, occ_board,
		&gs.MoveRays.Magic.BishopMagic, &gs.MoveRays.BishopRays)
	queen_straight_rays, queen_straight_xrays  = getMoveRay(*gs.Board.BlackQueens, occ_board,
		&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.RookRays)

	eval_move_rays.B_q_rays = combineRaysSets(queen_diag_rays, queen_straight_rays)
	eval_move_rays.B_q_xrays = combineRaysSets(queen_diag_xrays, queen_straight_xrays)


	// ------------------------------------------------------------------------
	// knight rays
	for _, ind := range gs.Board.WhiteKnights.Index() {
		eval_move_rays.W_kn_rays = append(eval_move_rays.W_kn_rays, gs.MoveRays.KnightRays[ind])
	}
	for _, ind := range gs.Board.BlackKnights.Index() {
		eval_move_rays.B_kn_rays = append(eval_move_rays.B_kn_rays, gs.MoveRays.KnightRays[ind])
	}

	// ------------------------------------------------------------------------
	// king threats
	var king_ind uint
	
	// -----------------------
		// white
	king_ind = gs.Board.WhiteKing.Index()[0]
			// sliding rays
	king_diag_rays, king_diag_xrays := getMoveRay(*gs.Board.WhiteKing, occ_board,
			&gs.MoveRays.Magic.BishopMagic, &gs.MoveRays.BishopRays)
	king_straight_rays, king_straight_xrays := getMoveRay(*gs.Board.WhiteKing, occ_board,
			&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.RookRays)

	eval_move_rays.W_king_rays.Diag = king_diag_rays[0]
	eval_move_rays.W_king_rays.Straight = king_straight_rays[0]
	eval_move_rays.W_king_xrays.Diag = king_diag_xrays[0]
	eval_move_rays.W_king_xrays.Straight = king_straight_xrays[0]

			// knight
	eval_move_rays.W_king_rays.Knight = gs.MoveRays.KnightRays[king_ind]
			// pawn
	eval_move_rays.W_king_rays.Pawn = gs.MoveRays.PawnCapRays[0][king_ind]

	// -----------------------
		// black
	king_ind = gs.Board.BlackKing.Index()[0]
			// sliding rays
	king_diag_rays, king_diag_xrays = getMoveRay(*gs.Board.BlackKing, occ_board,
			&gs.MoveRays.Magic.BishopMagic, &gs.MoveRays.BishopRays)
	king_straight_rays, king_straight_xrays = getMoveRay(*gs.Board.BlackKing, occ_board,
			&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.RookRays)

	eval_move_rays.B_king_rays.Diag = king_diag_rays[0]
	eval_move_rays.B_king_rays.Straight = king_straight_rays[0]
	eval_move_rays.B_king_xrays.Diag = king_diag_xrays[0]
	eval_move_rays.B_king_xrays.Straight = king_straight_xrays[0]

			// knight
	eval_move_rays.B_king_rays.Knight = gs.MoveRays.KnightRays[king_ind]
			// pawn
	eval_move_rays.B_king_rays.Pawn = gs.MoveRays.PawnCapRays[1][king_ind]


	return eval_move_rays
}

// get move rays and xrays for each piece (for sliding pieces)
func getMoveRay(pieces board.Bitboard, occ_board board.Bitboard, magic_sqs *[64]magic.Magicsq, 
				all_rays *[64]board.Bitboard) ([]board.Bitboard, []board.Bitboard) {

	var rays_slice, xrays_slice []board.Bitboard // output

	var rays, xrays board.Bitboard
	for _, ind := range pieces.Index() {
		
		// get attack rays
		magic_sq := magic_sqs[ind]
		rays = magic.Get_magic_rays(magic_sq, occ_board)
		rays_slice = append(rays_slice, rays)

		// get xray rays
		xrays = all_rays[ind]
		xrays_slice = append(xrays_slice, xrays)
	}

	return rays_slice, xrays_slice
}

// ----------------------------------------------------------------------------
// helper

// combine rays sets 
func combineRaysSets(ray_set []board.Bitboard, ray_set2 []board.Bitboard) []board.Bitboard {

	var combined []board.Bitboard
	for ind, ray := range ray_set {
		combined = append(combined, ray|ray_set2[ind])
	}

	return combined
}