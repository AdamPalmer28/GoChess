package move_gen

import (
	"chess/chess_engine/board"
	"chess/chess_engine/move_gen/magic"
)

// ! need to finish king moves
/*
Moves still to code:
	- castle
	- check move is legal
		- king can't move
	- double checks
	- check pins

*/

func GenKingMoves(king_bb board.Bitboard, team_bb board.Bitboard, castle_rights uint,
		magic_str_sqs *[64]magic.Magicsq, magic_diag_sqs *[64]magic.Magicsq,
		knight_rays *[64]board.Bitboard, king_rays *[64]board.Bitboard,
		opp_pawn board.Bitboard, opp_knight board.Bitboard, opp_bishop board.Bitboard,
		opp_rook board.Bitboard, opp_queen board.Bitboard, opp_king board.Bitboard) []uint {

	var movelist []uint

	// // generate basic moves
	// basic_moves = genBasicKingMove(king_bb, magic_str_sqs, team_bb, opp_pawn, opp_king)
	// movelist = append(movelist, basic_moves...)

	// // generate castle moves
	//if castle_rights > 0 {
	// castle_moves = genCastleMoves(king_bb, castle_rights, team_bb, opp_pawn, opp_knight, opp_bishop, opp_rook, opp_queen, opp_king)
	// movelist = append(movelist, castle_moves...)
	//}

	return movelist
}


// ! needs to check legal moves
func genBasicKingMove(king_bb board.Bitboard, king_rays *[64]board.Bitboard,
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

func GenCastling(king_bb board.Bitboard, castle_rights uint, team_bb board.Bitboard, 
	opp_pawn board.Bitboard, opp_knight board.Bitboard, opp_bishop board.Bitboard,
	opp_rook board.Bitboard, opp_queen board.Bitboard, opp_king board.Bitboard) []uint {

var movelist []uint
//relevant_sqs := board.Bitboard(0)

// king side castle
if castle_rights&0b01 > 0 {


}

return movelist
}

// ==================================================================
// helper function

// king ray generator for given square
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