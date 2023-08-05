package move_gen

import (
	"chess/chess_engine/board"
)

// simplify this function to take inputs
func GenPawnMoves(pawn_bb board.Bitboard, w_move bool, enpass uint,
			team_bb board.Bitboard, opp_bb board.Bitboard) []uint {

	var movelist []uint
	var moveno uint

	// pawn data
	var pawnstep int
	var prom_row uint
	var start_row uint

	if w_move {
		pawnstep = 8
		prom_row = 7
		start_row = 1
	} else {
		pawnstep = -8
		prom_row = 0
		start_row = 6
	}
	
	comb_occ := team_bb | opp_bb

	// loop through all the pawns
	pawn_inds := pawn_bb.Index()
	for _, ind := range pawn_inds {

		moveno = 0 // reset moveno
		col := ind % 8
		row := ind / 8
		
		// single pawn push
		en_sq := uint(int(ind) + pawnstep)

		if comb_occ & (1 << en_sq) == 0 {

			moveno |= en_sq << 6
			moveno |= ind

			// double pawn push
			if (row == start_row) {
				en_sq := uint(int(ind) + (pawnstep * 2))

				// check if the square is occupied
				if comb_occ & (1 << en_sq) == 0 {
					double_moveno := (en_sq << 6) | ind
					double_moveno |= 0b0001 << 12

					movelist = append(movelist, double_moveno)
				}
			// promotion
			} else if (row == prom_row) {

				promotion_list := promotion(moveno)
				movelist = append(movelist, promotion_list[:]...)

			// single push
			}

			movelist = append(movelist, moveno)
			
		}

		// pawn captures
		cap_sq := []uint{}
		if col != 0 { // left capture
			en_sq := uint(int(ind) + pawnstep - 1)
			cap_sq = append(cap_sq, en_sq)
		}
		if col != 7 { // right capture
			en_sq := uint(int(ind) + pawnstep + 1)
			cap_sq = append(cap_sq, en_sq)
		}

		for _, sq := range cap_sq {
			if comb_occ & (1 << sq) != 0 {
				(comb_occ & (1 << sq)).Print()
				moveno = 1 << 14
				moveno |= sq << 6 | ind

				// promotion
				if (row == prom_row) {
					promotion_list := promotion(moveno)
					movelist = append(movelist, promotion_list[:]...)
				} else { 
					movelist = append(movelist, moveno)
				}
			}
		}

		// enpassent capture
		if enpass < 64 {

			for _, sq := range cap_sq {
				if enpass == sq {
					moveno = 0b0101 << 12
					moveno |= enpass << 6 | ind

					movelist = append(movelist, moveno)
				}
			}
		}
	}

	return movelist

}


func promotion(moveno uint) [4]uint {

	promotion_list := [4]uint{}

	// queen
	queen_moveno := 0b1011 << 12 | moveno
	promotion_list[0] = queen_moveno

	// knight
	knight_moveno := 0b1000 << 12 | moveno
	promotion_list[1] = knight_moveno

	// rook
	rook_moveno := 0b1010 << 12 | moveno
	promotion_list[2] = rook_moveno

	// bishop
	bishop_moveno := 0b1001 << 12 | moveno
	promotion_list[3] = bishop_moveno

	return promotion_list
}