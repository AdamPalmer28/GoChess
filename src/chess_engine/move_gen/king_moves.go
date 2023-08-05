package move_gen

import (
	"chess/chess_engine/board"
)

func GenBasicKingMove(king_bb board.Bitboard, king_rays *[64]board.Bitboard,
				team_bb board.Bitboard, opp_bb board.Bitboard, opp_king_bb board.Bitboard) []uint {

	var movelist []uint
	var moveno uint

	ind := king_bb.Index()[0]

	
	move_ray := king_rays[ind]
	move_ray &= ^team_bb
	
	// make sure king doesn't move next to opp king
	opp_king_ind := opp_king_bb.Index()[0]
	move_ray &= ^king_rays[opp_king_ind]

	// captures
	caps := move_ray & opp_bb

	// non captures
	noncaps := move_ray & ^opp_bb

	// generate the moves nums
	for _, end_sq := range caps.Index() {
		moveno = 1 << 14 | uint(end_sq) << 6 | uint(ind)
		movelist = append(movelist, moveno)
	}

	for _, end_sq := range noncaps.Index() {
		moveno = uint(end_sq) << 6 | uint(ind)
		movelist = append(movelist, moveno)
	}


	return movelist

}


func KingRays(ind int) board.Bitboard {

	var moves board.Bitboard = 0

	var vals = []int{7, 8, 9, -7, -8, -9, 1, -1}
	var col_change = []int{-1, 0, 1, 1, 0, -1, 1, -1}
	var row_change = []int{1, 1, 1, -1, -1, -1, 0, 0}

	for i, val := range vals {
		col_c := col_change[i]
		row_c := row_change[i]

		// validate the move
		if ((ind+val) % 8 - ind % 8 != col_c) ||
		   ((ind+val) / 8 - ind / 8 != row_c) ||
		   ((ind+val) < 0 || (ind+val) > 63) {
			continue
		}

		moves |= 1 << uint(ind+val)
	}

	return moves
}