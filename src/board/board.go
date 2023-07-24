package board

import "fmt"

type ChessBoard struct {
	WhitePawns   Bitboard
	WhiteKnights Bitboard
	WhiteBishops Bitboard
	WhiteRooks   Bitboard
	WhiteQueens  Bitboard
	WhiteKing    Bitboard

	BlackPawns   Bitboard
	BlackKnights Bitboard
	BlackBishops Bitboard
	BlackRooks   Bitboard
	BlackQueens  Bitboard
	BlackKing    Bitboard

}


func (cb *ChessBoard) Print() {

	fmt.Println("   a  b  c  d  e  f  g  h")
	for rank := uint(0); rank < 8; rank++ {
		fmt.Printf("%d ", 8 - rank)
		for file := uint(0); file < 8; file++ {

			square := ((7-rank) * 8) + file
			mask := Bitboard(1) << square

			// Check if any piece is present on the square
			piece := " - "
			if (cb.WhitePawns & mask) != 0 {
				piece = "wP "
			} else if (cb.WhiteKnights & mask) != 0 {
				piece = "wN "
			} else if (cb.WhiteBishops & mask) != 0 {
				piece = "wB "
			} else if (cb.WhiteRooks & mask) != 0 {
				piece = "wR "
			} else if (cb.WhiteQueens & mask) != 0 {
				piece = "wQ "
			} else if (cb.WhiteKing & mask) != 0 {
				piece = "wK "

			} else if (cb.BlackPawns & mask) != 0 {
				piece = " p "
			} else if (cb.BlackKnights & mask) != 0 {
				piece = " n " 
			} else if (cb.BlackBishops & mask) != 0 {
				piece = " b "
			} else if (cb.BlackRooks & mask) != 0 {
				piece = " r "
			} else if (cb.BlackQueens & mask) != 0 {
				piece = " q "
			} else if (cb.BlackKing & mask) != 0 {
				piece = " k "
			}

			fmt.Print(piece)
		}
		fmt.Println()
	}
}

func (cb *ChessBoard) listBB(white bool) [6]Bitboard {
	// return a list of bitboards for each piece type
	if white {
		return [6]Bitboard{cb.WhitePawns, cb.WhiteKnights, cb.WhiteBishops, 
							cb.WhiteRooks, cb.WhiteQueens, cb.WhiteKing}
	} else {
		return [6]Bitboard{cb.BlackPawns, cb.BlackKnights, cb.BlackBishops, 
							cb.BlackRooks, cb.BlackQueens, cb.BlackKing}
	}
}


func (cb *ChessBoard) move(move_num uint, white_move bool) (uint, uint) {

	start_sq := move_num & 0x3F
	finish_sq := (move_num >> 6) & 0x3F
	//special := (move_num >> 12) & 0xF

	BB_list := [6]Bitboard{}
	Opp_BB_list := [6]Bitboard{}

	if white_move {
		// white move
		BB_list = cb.listBB(true)
		Opp_BB_list = cb.listBB(false)
		} else {
		// black move
		BB_list = cb.listBB(false) 
		Opp_BB_list = cb.listBB(true)
	}

	// update piece bitboards
	var piece_moved uint = 6
	for ind, BB := range BB_list {

		if BB & (1 << start_sq) != 0 {

			piece_moved = uint(ind)
			BB_list[ind].flip(start_sq)
			BB_list[ind].flip(finish_sq)
			break
		}
	}
	// check for capture
	var cap_piece uint = 6
	for ind, BB := range Opp_BB_list {

		if BB & (1 << finish_sq) != 0 {
			
			cap_piece = uint(ind)
			Opp_BB_list[ind].flip(finish_sq)
			break
		}
	}

	return piece_moved, cap_piece

}