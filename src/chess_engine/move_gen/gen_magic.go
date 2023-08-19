package move_gen

import "chess/chess_engine/board"

type magicsq struct {
	index uint
	magic uint64

	diag bool

	// attack rays
	rays []board.Bitboard

	// shift
	shift uint

	// occupancy mask
	occ_mask board.Bitboard
}

type gen_magicsq struct {
	
	magicsq [64]magicsq

	all_occ [64][]board.Bitboard
}

// ============================================================================


