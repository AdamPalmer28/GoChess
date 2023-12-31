package board

/* Move type hierarchy
----------------------------------------------
1. split movenum and variables
2. check moves type
	a. normal move (special < 2)
	b. castle (special < 4)
	c. capture (special < 6)
		i. enpassent (special == 5)
	d. promotion (special > 7)
		i promotion capture (special > 11)
3. update piece bitboards
*/

// Make a move on the bitboards
func (cb *ChessBoard) Move(move_num uint, white_move bool) (uint, uint) {

	start_sq := move_num & 0x3F
	finish_sq := (move_num >> 6) & 0x3F
	special := (move_num >> 12) & 0xF

	BB_list := cb.ListBB(white_move) 
	Opp_BB_list := cb.ListBB(!white_move)

	var piece_moved uint
	var cap_piece uint = 6
		
	var piece_start_ind uint = 0
	var piece_cap_start_ind uint = 6
	if !white_move { 
		piece_start_ind = 6 
		piece_cap_start_ind = 0
	}


	// Move type (see top of file)
	// ========================================================================

	if special < 2 { // normal move

		piece_moved = normal_move(start_sq, finish_sq, BB_list)

		// ----------------------------------------------------------
	} else if special < 4 { // castle
		var rook_start_sq uint
		var rook_finish_sq uint

		if special == 2 { // king side castle
			
			rook_start_sq = finish_sq + 1 // rook start square
			rook_finish_sq = finish_sq - 1 // rook finish square

		} else { // queen side castle 
		
			rook_start_sq = finish_sq - 2 // rook start square
			rook_finish_sq = finish_sq + 1 // rook finish square
		}

		// move king
		piece_moved = normal_move(start_sq, finish_sq, BB_list)
		// move rook
		_ = normal_move(rook_start_sq, rook_finish_sq, BB_list)

		// update rook locations on castle
		cb.UpdatePieceLocations(piece_start_ind + 3)

		// ----------------------------------------------------------
	} else if special <= 5 { // capture
		var cap_sq uint

		if special == 5 {  // enpassent

			var fwd int = 8
			if !white_move { fwd = -8 }

			cap_sq = finish_sq - uint(fwd)
		
			} else {
			cap_sq = finish_sq
		}

		piece_moved = normal_move(start_sq, finish_sq, BB_list)
		cap_piece = cap_move(cap_sq, Opp_BB_list)

		cb.UpdateSideBB(!white_move) // update opp bitboards

		// ----------------------------------------------------------
	} else if special > 7 { // promotion

		if special > 11 { // promotion capture
			
			cap_piece = cap_move(finish_sq, Opp_BB_list)

			cb.UpdateSideBB(!white_move) // update opp bitboards
		}
		
		piece_moved = 0 // pawn
		prom_move(special, start_sq, finish_sq, BB_list)

		// update piece locations for knight, bishop, rook, queen
		cb.UpdatePieceLocations(piece_start_ind + 1)
		cb.UpdatePieceLocations(piece_start_ind + 2)
		cb.UpdatePieceLocations(piece_start_ind + 3)
		cb.UpdatePieceLocations(piece_start_ind + 4)
	}

	cb.UpdateSideBB(white_move) // update bitboards after move execution

	// update piece locations
	cb.UpdatePieceLocations(piece_start_ind + piece_moved)
	if cap_piece < 6 {
		cb.UpdatePieceLocations(piece_cap_start_ind + cap_piece)
	}

	return piece_moved, cap_piece

}

// ========================================================================
// Helper functions


// Standard moves
func normal_move(start_sq uint, finish_sq uint, BB_list [6]*Bitboard) uint {

	// update piece bitboards
	var piece_moved uint
	for ind, BB := range BB_list {

		if *BB & (1 << start_sq) != 0 {

			piece_moved = uint(ind)
			BB.flip(start_sq)
			BB.flip(finish_sq)
			
			break
		}
	}
	return piece_moved
}

// Captures
func cap_move(take_sq uint, BB_list [6]*Bitboard) uint {

	// update piece bitboards
	var cap_piece uint
	for ind, BB := range BB_list {

		if *BB & (1 << take_sq) != 0 {

			cap_piece = uint(ind)
			BB.flip(take_sq)
			
			break
		}
	}
	return cap_piece
}

// Promotions
func prom_move(special uint, start uint, finish_sq uint, BB_list [6]*Bitboard) {

	special = special & 0b0011 // relevant bits
	var promo_bb uint

	// special:   		 knight = 0, bishop = 1, rook = 2, queen = 3
	// BB_list postion:  knight = 1, bishop = 2, rook = 3, queen = 4
	promo_bb = special + 1 


	// update pawn
	BB_list[0].flip(start)

	// update promotion piece
	BB_list[promo_bb].flip(finish_sq)

}