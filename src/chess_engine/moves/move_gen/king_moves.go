package move_gen

import (
	"chess/chess_engine/board"
)

func GenBasicKings(king_bb board.Bitboard, king_rays *[64]board.Bitboard,
				team_bb board.Bitboard, opp_bb board.Bitboard) []uint {

	var movelist []uint
	var moveno uint

	ind := king_bb.Index()[0]

	move_ray := king_rays[ind]

	return movelist

}