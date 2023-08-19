package move_gen

import (
	"chess/chess_engine/board"
	"chess/chess_engine/move_gen/magic"
)


func GenSlidingMoves(bb board.Bitboard, 
	rays *[64]board.Bitboard, piece_type uint,
	team_bb board.Bitboard, opp_bb board.Bitboard) []uint {

	var movelist []uint
	var moveno uint
	var move_ray board.Bitboard

	// loop through all the rooks
	p_inds := bb.Index()
	for _, ind := range p_inds {

		// basic generation
		if piece_type == 0 {  
			// rook
			move_ray = magic.SlidingRays(ind, team_bb|opp_bb)
			move_ray &= ^team_bb
		} else if piece_type == 1 {
			// bishop
			move_ray = magic.DiagonalRays(ind, team_bb|opp_bb)
			move_ray &= ^team_bb
		} else { // piece_type == 2
			// queen
			move_ray = magic.SlidingRays(ind, team_bb|opp_bb)
			move_ray |= magic.DiagonalRays(ind, team_bb|opp_bb)
			move_ray &= ^team_bb
		}
		// generate the moves nums

		// captures
		caps := move_ray & opp_bb
		caps_inds := caps.Index()

		for _, end_sq := range caps_inds {

			moveno = 1<<14 | uint(end_sq)<<6 | uint(ind)
			movelist = append(movelist, moveno)
		}

		// non captures
		noncaps := move_ray & ^opp_bb
		noncaps_inds := noncaps.Index()

		for _, end_sq := range noncaps_inds {

			moveno = uint(end_sq)<<6 | uint(ind)
			movelist = append(movelist, moveno)
		}
	}
	return movelist
}

