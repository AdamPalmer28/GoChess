package gamestate

func (gs *GameState) Make_move(move uint) {

	// get the start and finish squares
	start_sq := move & 0x3F
	finish_sq := (move >> 6) & 0x3F
	special := (move >> 12) & 0xF

	// update piece bitboards (gs.Board)
	piece_moved, cap_piece := gs.Board.Move(move, gs.White_to_move)

	// piece_moved_str := PieceValLookup[int(piece_moved)]
	// cap_piece_str := PieceValLookup[int(cap_piece)]

	var CastleRight *uint
	var row uint
	var fwd int

	if gs.White_to_move {
		CastleRight = &gs.WhiteCastle
		row = 0
		fwd = 8
	} else {
		CastleRight = &gs.BlackCastle
		row = 7
		fwd = -8
	}

	var sq uint
	var intrest_sq uint
	// update castling rights
	if *CastleRight > 0 {

		if piece_moved == 5 { // king moved
			*CastleRight = 0

		} else if piece_moved == 3 || cap_piece == 3 { // rook moved

			intrest_sq = start_sq
			if cap_piece == 3 {
				intrest_sq = finish_sq
			}

			// king side
			if *CastleRight&0b01 > 0 {
				sq = row*8 + 7
				if intrest_sq == sq {
					*CastleRight &^= 0b01
				}
			}
			// queen side
			if *CastleRight&0b10 > 0 {
				sq = row * 8
				if intrest_sq == sq {
					*CastleRight &^= 0b10
				}
			}
		}
	}

	// update enpassent index
	if special == 0b0001 {
		// double pawn push
		gs.Enpass_ind = uint(int(finish_sq) - fwd)

	} else {
		gs.Enpass_ind = 0
	}

	// capture piece
	if (special & 4) > 0 {
		piece := 0 // placeholder
		gs.Cap_pieces = append(gs.Cap_pieces, [2]int{int(gs.Moveno), piece})
	}

	// update previous moves
	gs.PrevMoves = append(gs.PrevMoves, move)

	if cap_piece != 6 {
		gs.Cap_pieces = append(gs.Cap_pieces,
			[2]int{int(gs.Moveno), int(cap_piece)})
	}

	// update move number
	gs.Moveno++

	// change move color
	gs.White_to_move = !gs.White_to_move

	// generate new moves
	gs.GenMoves()
}