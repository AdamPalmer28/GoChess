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