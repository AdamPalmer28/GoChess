package gamestate

import (
	"chess/chess_engine/move_gen"
)

// make move on gamestate
func (gs *GameState) Make_move(move uint) {

	// get the start and finish squares
	start_sq := move & 0x3F
	finish_sq := (move >> 6) & 0x3F
	special := (move >> 12) & 0xF

	// update piece bitboards (gs.Board)
	piece_moved, cap_piece := gs.Board.Move(move, gs.White_to_move)

	var CastleRight *uint
	var OppCastleRight *uint
	var row uint
	var fwd int

	if gs.White_to_move {
		CastleRight = &gs.WhiteCastle
		OppCastleRight = &gs.BlackCastle
		row = 0
		fwd = 8
	} else {
		CastleRight = &gs.BlackCastle
		OppCastleRight = &gs.WhiteCastle
		row = 7
		fwd = -8
	}
	gs.History.CastleRight = append(gs.History.CastleRight, *CastleRight) // castle rights at start of the move

	var sq uint
	// update castling rights
	if *CastleRight > 0 {

		if piece_moved == 5 { // king moved
			*CastleRight = 0

		} else if piece_moved == 3 { // rook moved

			// king side
			if *CastleRight&0b01 > 0 {
				sq = row*8 + 7
				if start_sq == sq {
					*CastleRight &^= 0b01
				}
			}
			// queen side
			if *CastleRight&0b10 > 0 {
				sq = row * 8
				if start_sq == sq {
					*CastleRight &^= 0b10
				}
			}
		}
	}
	if (*OppCastleRight > 0) && (cap_piece == 3) {

		opp_row := 7 - row
		// king side
		if *OppCastleRight&0b01 > 0 {
			sq = opp_row*8 + 7
			if finish_sq == sq {
				*OppCastleRight &^= 0b01
			}
		}
		// queen side
		if *OppCastleRight&0b10 > 0 {
			sq = opp_row * 8
			if finish_sq == sq {
				*OppCastleRight &^= 0b10
			}
		}

	}

	// update enpassent index
	if special == 0b0001 {
		// double pawn push
		gs.Enpass_ind = uint(int(finish_sq) - fwd)

	} else {
		gs.Enpass_ind = 64
	}

	// update gamestate

	// update previous moves
	gs.History.PrevMoves = append(gs.History.PrevMoves, move)
	gs.History.Cap_pieces = append(gs.History.Cap_pieces, cap_piece)

	// update move number
	gs.Moveno++

	// change move color
	gs.White_to_move = !gs.White_to_move

	gs.Next_move()
}


// Calcs for next move
func (gs *GameState) Next_move() {

	gs.Make_BP() // make board perspectives

	pinned_pieces := gs.GetCheck() // get check status

	if gs.InCheck {
		gs.GenCheckMoves() // generate moves // in check
	} else {
		gs.GenMoves() // generate moves
	}

	remove_illegal_moves(gs, pinned_pieces) // remove illegal moves

	// check for game over
	if len(gs.MoveList) == 0 {
		// no moves 
		gs.GameOver = true
		
		if gs.InCheck {
			if gs.White_to_move {
				//println("Black wins - checkmate")
			} else {
				//println("White wins - checkmate")
			}
		} else {
			//println("Stalemate")
		}

	} else {
		gs.GameOver = false
	}
}


// updates InCheck status depending on position
func (gs *GameState) GetCheck() map[uint][]uint {

	// get the king square
	var king_sq uint
	king_sq = gs.PlayerKingSaftey.King_sq

	// check if the king is attacked

	results, pinned_pieces := move_gen.BoardKingAnalysis(gs.PlayerKingSaftey,
		gs.MoveRays.KnightRays[king_sq],
		&gs.MoveRays.Magic.RookMagic[king_sq],
		&gs.MoveRays.Magic.BishopMagic[king_sq])

	gs.InCheck = !results

	return pinned_pieces
}