package magic

import "chess/chess_engine/board"


func Get_magic_rays(magic_sq Magicsq,
	occ board.Bitboard) board.Bitboard {

	// get magic square
	magic_num := magic_sq.magic
	shift := magic_sq.shift

	rel_occ := occ & magic_sq.occ_mask

	// get attack rays
	magic_index := (rel_occ * magic_num) >> shift

	return magic_sq.attack_rays[magic_index]
}

func GenMagicMoves(bb board.Bitboard,
	magic_sqs *[64]Magicsq,
	team_bb board.Bitboard, opp_bb board.Bitboard) []uint {

	var movelist []uint
	var moveno uint
	var move_ray board.Bitboard

	// loop through all the rooks
	p_inds := bb.Index()
	for _, ind := range p_inds {

		move_ray = Get_magic_rays(magic_sqs[ind], team_bb|opp_bb)
		move_ray &= ^team_bb

		// generate the moves nums

		// captures
		caps := move_ray & opp_bb
		caps_inds := caps.Index()

		for _, end_sq := range caps_inds {

			moveno = 1<<14 | end_sq<<6 | ind
			movelist = append(movelist, moveno)
		}

		// non captures
		noncaps := move_ray & ^opp_bb
		noncaps_inds := noncaps.Index()

		for _, end_sq := range noncaps_inds {

			moveno = end_sq <<6 | ind
			movelist = append(movelist, moveno)
		}
	}
	return movelist
}
