package move_gen

import (
	"chess/chess_engine/board"
	"chess/chess_engine/move_gen/magic"
)

type KingSafetyRelBB struct {
	// friendly pieces
	King_sq uint
	King_bb board.Bitboard
	Team_bb board.Bitboard
	Team_bb_no_king board.Bitboard
	Fwd int

	// opp pieces
	Opp_bb board.Bitboard
	Opp_pawn_bb board.Bitboard
	Opp_knight_bb board.Bitboard
	Opp_bishop_bb board.Bitboard
	Opp_rook_bb board.Bitboard
	Opp_queen_bb board.Bitboard

	Opp_king_bubble board.Bitboard // squares around opp king
}



func GenKingMoves(kingRel KingSafetyRelBB, castle_rights uint,
		magic_str_sqs *[64]magic.Magicsq, magic_diag_sqs *[64]magic.Magicsq,
		knight_rays *[64]board.Bitboard, king_rays *[64]board.Bitboard) []uint {

	var movelist []uint


	// generate basic moves
	basic_moves := king_rays[kingRel.King_sq]
	basic_moves &= (^kingRel.Team_bb & ^kingRel.Opp_king_bubble)

	cap_moves := basic_moves & kingRel.Opp_bb
	noncap_moves := basic_moves & ^kingRel.Opp_bb

	// generate the basic moves
	cap_moves_nums := legal_king_moves(cap_moves, kingRel, 0b0100, 
							magic_str_sqs, magic_diag_sqs, knight_rays)
	movelist = append(movelist, cap_moves_nums...)

	noncap_moves_nums := legal_king_moves(noncap_moves, kingRel, 0b0000,
							magic_str_sqs, magic_diag_sqs, knight_rays)
	movelist = append(movelist, noncap_moves_nums...)


	// castle moves
	if castle_rights > 0 {
		castle_moves := GenCastling(castle_rights, kingRel,
									magic_str_sqs, magic_diag_sqs, knight_rays)
		movelist = append(movelist, castle_moves...)
	}

	return movelist
}


// generate castling moves
func GenCastling(castle_rights uint, kingRel KingSafetyRelBB,
	magic_str_sqs *[64]magic.Magicsq, magic_diag_sqs *[64]magic.Magicsq, 
	knight_rays *[64]board.Bitboard) []uint {

	var movelist []uint
	var castle_occ board.Bitboard

	var castle_checks []uint 
	// king side castle
	if castle_rights&0b01 > 0 {
		// check safe for king
		castle_occ = kingRel.King_bb << 1 | kingRel.King_bb << 2
		
		if (castle_occ) & (kingRel.Team_bb | kingRel.Opp_bb) == 0 {
			// check safe for king
			castle_checks = legal_king_moves(castle_occ, kingRel, 0b0010,
							magic_str_sqs, magic_diag_sqs, knight_rays)
			if len(castle_checks) == 2 {
				movelist = append(movelist, castle_checks[1])
			}
		}
	}

	// queen side castle
	if castle_rights&0b10 > 0 {

		// check safe for king
		queen_side_castle := kingRel.King_bb >> 1 | kingRel.King_bb >> 2
		// check square occs
		castle_occ = queen_side_castle | queen_side_castle >> 1

		if (castle_occ) & (kingRel.Team_bb | kingRel.Opp_bb) == 0 {
			// check safe for king
			castle_checks = legal_king_moves(queen_side_castle, kingRel, 0b0011,
							magic_str_sqs, magic_diag_sqs, knight_rays)	
			if len(castle_checks) == 2 {
				movelist = append(movelist, castle_checks[0])
			}
		}
	}

	return movelist
}

// ==================================================================
// helper function

// create king moves are legal
func legal_king_moves(movebb board.Bitboard, kingRel KingSafetyRelBB, special uint,
		magic_str_sqs *[64]magic.Magicsq, magic_diag_sqs *[64]magic.Magicsq, 
		knight_rays *[64]board.Bitboard) []uint {

	var movelist []uint
	special = special << 12

	var moveno uint
	var magic_str_sq *magic.Magicsq
	var magic_diag_sq *magic.Magicsq
	var knight_ray board.Bitboard

	var legal bool
	// loop through all the king moves
	for _, end_sq := range movebb.Index() {

		magic_str_sq = &magic_str_sqs[end_sq]
		magic_diag_sq = &magic_diag_sqs[end_sq]
		knight_ray = knight_rays[end_sq]

		legal = check_king_safety(end_sq, kingRel, knight_ray, magic_str_sq, magic_diag_sq)
		if legal {
			// make move number
			moveno = special | (end_sq << 6)  | (kingRel.King_sq)
			movelist = append(movelist, moveno)
		}
	}

	return movelist
}


// ==================================================================
// Intialiase function

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