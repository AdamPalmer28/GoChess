package gamestate

func (gs *GameState) Make_move(move uint) {

	// get the start and finish squares

	special := (move >> 12) & 0xF

	// update piece bitboards (gs.Board)
	piece_moved, cap_piece := gs.Board.Move(move, gs.White_to_move)

	piece_moved_str := PieceValLookup[int(piece_moved)]
	cap_piece_str := PieceValLookup[int(cap_piece)]

	println(piece_moved_str, cap_piece_str)

	// update castling rights

	// update enpassent index

	// capture piece
	if (special & 4) > 0 {
		piece := 0 // placeholder
		gs.Cap_pieces = append(gs.Cap_pieces, [2]int{int(gs.Moveno), piece})
	}

	// update previous moves
	gs.PrevMoves = append(gs.PrevMoves, move)

	// update move number
	gs.Moveno++

	// change move color
	gs.White_to_move = !gs.White_to_move

	// generate new moves
	gs.GenMoves()
}