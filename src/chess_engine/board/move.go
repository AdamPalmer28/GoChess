package board

// Make a move on the board
func (cb *ChessBoard) Move(move_num uint, white_move bool) (uint, uint) {

	start_sq := move_num & 0x3F
	finish_sq := (move_num >> 6) & 0x3F
	//special := (move_num >> 12) & 0xF

	BB_list := [6]*Bitboard{}
	Opp_BB_list := [6]*Bitboard{}


	BB_list = cb.listBB(white_move) 
	Opp_BB_list = cb.listBB(!white_move)
		
	
	// update piece bitboards
	var piece_moved uint = 6
	for ind, BB := range BB_list {

		if *BB & (1 << start_sq) != 0 {

			piece_moved = uint(ind)
			BB.flip(start_sq)
			BB.flip(finish_sq)
			
			cb.UpdateSideBB(white_move)
			break
		}
	}
	// check for capture
	var cap_piece uint = 6
	for ind, BB := range Opp_BB_list {

		if *BB & (1 << finish_sq) != 0 {
			
			cap_piece = uint(ind)
			BB.flip(finish_sq)

			cb.UpdateSideBB(!white_move)
			break
		}
	}

	return piece_moved, cap_piece

}