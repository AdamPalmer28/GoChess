package cli_debug

import (
	"chess/src/chess_engine/board"
	"chess/src/chess_engine/move_gen"
)

func Move_cli(movelist move_gen.MoveList, sq string) {
	// print bitboard of current gamestate

	square := sq
	ind := Move_to_index(square)

	Sq_moves(movelist, ind)

}


func Sq_moves(list move_gen.MoveList, ind uint)  {

	// print bb for a square
	var sq_moves board.Bitboard = 0
	var end_sq uint

	for _, move := range list {
		
		if move & 0b111111 == ind {

			end_sq = (move >> 6) & 0b111111
			sq_moves |= 1 << end_sq
		}
	}

	sq_moves.Print()
}


func Move_to_index(cord string) uint {
	// convert a chess cord to an index

	var ind uint

	cord = cord[0:2]
	
	letter := cord[0]
	number := cord[1]

	ind = uint(number-'1')*8 + uint(letter-'a')

	return ind
}