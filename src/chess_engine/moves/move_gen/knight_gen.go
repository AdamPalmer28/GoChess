package move_gen

import (
	"chess/chess_engine/board"
)

func GenKnightMoves(knight_bb board.Bitboard, knight_rays *[64]board.Bitboard,
				team_bb board.Bitboard, opp_bb board.Bitboard) []uint {

	var movelist []uint
	var moveno uint


	// loop through all the knights
	knight_inds := knight_bb.Index()
	for _, ind := range knight_inds {

		move_ray := knight_rays[ind]
		move_ray &= ^team_bb
		// generate the moves nums

		// captures
		caps := move_ray & opp_bb
		caps_inds := caps.Index()

		for _, end_sq := range caps_inds {
			moveno = 1 << 14 | uint(end_sq) << 6 | uint(ind)
			movelist = append(movelist, moveno)
		}

		// non captures
		noncaps := move_ray & ^opp_bb
		noncaps_inds := noncaps.Index()

		for _, end_sq := range noncaps_inds {
			moveno = uint(end_sq) << 6 | uint(ind)
			movelist = append(movelist, moveno)
		}

	}
	return movelist
}


func KnightMoves(ind int) board.Bitboard {

	var moves board.Bitboard = 0

	var vals = []int{6, 10, 15, 17, -6, -10, -15, -17}
	var col_change = []int{-2, 2, -1, 1, 2, -2, 1, -1}
	var row_change = []int{1, 1, 2, 2, -1, -1, -2, -2}

	for i, val := range vals {
		col_c := col_change[i]
		row_c := row_change[i]

		// validate the move
		if ((ind+val) % 8 - ind % 8 != col_c) ||
		   ((ind+val) / 8 - ind / 8 != row_c) ||
		   ((ind+val) < 0 || (ind+val) > 63) {
			continue
		} 		

		moves |= (1 << (ind + val))
	}

	return moves
}