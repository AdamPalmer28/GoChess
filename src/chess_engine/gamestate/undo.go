package gamestate

import (
	"chess/src/chess_engine/move_gen"
)

func (gs *GameState) Undo() {

	last_move_num := gs.MoveNo - 1
	last_move_ind := last_move_num - 1

	// restore board
	gs.MoveNo -= 1
	gs.DisplayMoveNo -= 1
	gs.HalfMoveNo -= 1
	gs.White_to_move = !gs.White_to_move
	gs.Make_BP()
	var player move_gen.BoardPerpective = gs.PlayerBoard

	// get last move
	last_move := gs.History.PrevMoves[last_move_ind]
	special := (last_move >> 12) & 0xF
	prev_end_sq := (last_move >> 6) & 0x3F
	prev_start_sq := last_move & 0x3F

	// piece locations
	var piece_start_ind uint = 0
	var piece_cap_start_ind uint = 6
	if !gs.White_to_move { 
		piece_start_ind = 6 
		piece_cap_start_ind = 0
	}
	var cap_piece uint = 6

	// reverse move
	rev_move := (prev_start_sq << 6) | prev_end_sq

	piece_moved, _ := gs.Board.Move(rev_move, gs.White_to_move) // moves piece back to start_sq


	// if promotion
	if special&0b1000 > 0 {
		// remove the promoted piece back to pawn
		player_bbs := gs.Board.ListBB(gs.White_to_move)

		promo_piece := special & 0b0011
		var bb_ind uint
		// get the promoted piece
		if promo_piece == 0b0011 { // queen
			bb_ind = 4
		} else if promo_piece == 0b0010 { // rook
			bb_ind = 3
		} else if promo_piece == 0b0001 { // bishop
			bb_ind = 2
		} else { // knight
			bb_ind = 1
		}

		// remove the promoted piece
		*player_bbs[bb_ind] &^= (1 << prev_start_sq)

		// add the pawn back
		*player_bbs[0] |= (1 << prev_start_sq)

	}

	// capture
	if special&0b0100 > 0 {

		opp_bbs := gs.Board.ListBB(!gs.White_to_move)
		cap_piece = gs.History.Cap_pieces[last_move_ind]

		// return old piece to board
		if special == 0b0101 { // enpassent capture
			// add the captured pawn back
			*opp_bbs[cap_piece] |= (1 << (int(prev_end_sq) - player.Fwd))

		} else { // normal or promotion capture
			*opp_bbs[cap_piece] |= (1 << prev_end_sq)
		}

		gs.Board.UpdateSideBB(!gs.White_to_move) // update BB

		// update piece locations - pawn
		gs.Board.UpdatePieceLocations(piece_start_ind + 0)
	}

	// castle
	if (special == 0b0010) || (special == 0b0011) {
		// move rook back - king is already moved back

		player_bbs := gs.Board.ListBB(gs.White_to_move)
		rook_bb := player_bbs[3]

		var rook_cur_sq uint
		var rook_prev_sq uint

		if special == 0b0010 { // king side castle
			rook_cur_sq = prev_end_sq - 1  // rook current square
			rook_prev_sq = prev_end_sq + 1 // rook square
		} else { // queen side castle
			rook_cur_sq = prev_end_sq + 1  // rook current square
			rook_prev_sq = prev_end_sq - 2 // rook start square
		}

		// move rook
		*rook_bb &^= (1 << rook_cur_sq)
		*rook_bb |= (1 << rook_prev_sq)

		gs.Board.UpdateSideBB(gs.White_to_move) // update

		// update piece locations - rook
		gs.Board.UpdatePieceLocations(piece_start_ind + 3)
	}

	// check update enpassent
	gs.Enpass_ind = gs.History.EnPassHist[last_move_ind]

	// update castle rights
	if gs.White_to_move {
		gs.WhiteCastle = gs.History.CastleRight[last_move_ind]
	} else {
		gs.BlackCastle = gs.History.CastleRight[last_move_ind]
	}

	gs.History.PrevMoves = gs.History.PrevMoves[:last_move_ind]
	gs.History.Cap_pieces = gs.History.Cap_pieces[:last_move_ind]
	gs.History.CastleRight = gs.History.CastleRight[:last_move_ind]
	gs.History.EnPassHist = gs.History.EnPassHist[:last_move_ind]


	// update piece locations
	gs.Board.UpdatePieceLocations(piece_start_ind + piece_moved)
	if cap_piece < 6 {
		gs.Board.UpdatePieceLocations(piece_cap_start_ind + cap_piece)
	}

	gs.Next_move() // generate new moves
}