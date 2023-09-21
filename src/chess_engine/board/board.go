package board

import (
	"fmt"
	"strconv"
)

type ChessBoard struct {
	WhitePawns   *Bitboard
	WhiteKnights *Bitboard
	WhiteBishops *Bitboard
	WhiteRooks   *Bitboard
	WhiteQueens  *Bitboard
	WhiteKing    *Bitboard

	BlackPawns   *Bitboard
	BlackKnights *Bitboard
	BlackBishops *Bitboard
	BlackRooks   *Bitboard
	BlackQueens  *Bitboard
	BlackKing    *Bitboard

	White        Bitboard
	Black        Bitboard
}

// ListBB returns a list of bitboards for each size
func (cb *ChessBoard) ListBB(white bool) [6]*Bitboard {
	// return a list of bitboards for each piece type
	if white {
		return [6]*Bitboard{cb.WhitePawns, cb.WhiteKnights, cb.WhiteBishops, 
			cb.WhiteRooks, cb.WhiteQueens, cb.WhiteKing}
	} else {
		return [6]*Bitboard{cb.BlackPawns, cb.BlackKnights, cb.BlackBishops, 
			cb.BlackRooks, cb.BlackQueens, cb.BlackKing}
	}
}

// Creates a BB of all the pieces on the board for a given side
func (cb *ChessBoard) UpdateSideBB(white bool) {
	
	if white {
		cb.White = *cb.WhitePawns | *cb.WhiteKnights | *cb.WhiteBishops | 
		*cb.WhiteRooks | *cb.WhiteQueens | *cb.WhiteKing
	} else {
		cb.Black = *cb.BlackPawns | *cb.BlackKnights | *cb.BlackBishops | 
		*cb.BlackRooks | *cb.BlackQueens | *cb.BlackKing
	}
}	
	
	
func (cb *ChessBoard) Print() {

	fmt.Println("  | a   b   c   d   e   f   g   h")
	fmt.Println("--+---+---+---+---+---+---+---+---+")
	for rank := uint(0); rank < 8; rank++ {
	
		fmt.Printf("%d |", 8 - rank)
		for file := uint(0); file < 8; file++ {

			square := ((7-rank) * 8) + file
			mask := Bitboard(1) << square

			// Check if any piece is present on the square
			piece := "   |"
			if (*cb.WhitePawns & mask) != 0 {
				piece = " P |"
			} else if (*cb.WhiteKnights & mask) != 0 {
				piece = " N |"
			} else if (*cb.WhiteBishops & mask) != 0 {
				piece = " B |"
			} else if (*cb.WhiteRooks & mask) != 0 {
				piece = " R |"
			} else if (*cb.WhiteQueens & mask) != 0 {
				piece = " Q |"
			} else if (*cb.WhiteKing & mask) != 0 {
				piece = " K |"

			} else if (*cb.BlackPawns & mask) != 0 {
				piece = " p |"
			} else if (*cb.BlackKnights & mask) != 0 {
				piece = " n |" 
			} else if (*cb.BlackBishops & mask) != 0 {
				piece = " b |"
			} else if (*cb.BlackRooks & mask) != 0 {
				piece = " r |"
			} else if (*cb.BlackQueens & mask) != 0 {
				piece = " q |"
			} else if (*cb.BlackKing & mask) != 0 {
				piece = " k |"
			}

			fmt.Print(piece)
		}
		fmt.Println()
		fmt.Println("  +---+---+---+---+---+---+---+---+")
	}
}

func (cb *ChessBoard) Copy() ChessBoard {

	var new_cb ChessBoard

	wp, wn, wb, wr, wq, wk := cb.WhitePawns.Copy(), cb.WhiteKnights.Copy(), cb.WhiteBishops.Copy(), 
							cb.WhiteRooks.Copy(), cb.WhiteQueens.Copy(), cb.WhiteKing.Copy()
	bp, bn, bb, br, bq, bk := cb.BlackPawns.Copy(), cb.BlackKnights.Copy(), cb.BlackBishops.Copy(),
							cb.BlackRooks.Copy(), cb.BlackQueens.Copy(), cb.BlackKing.Copy()

	new_cb.WhitePawns = &wp
	new_cb.WhiteKnights = &wn
	new_cb.WhiteBishops = &wb
	new_cb.WhiteRooks = &wr
	new_cb.WhiteQueens = &wq
	new_cb.WhiteKing = &wk

	new_cb.BlackPawns = &bp
	new_cb.BlackKnights = &bn
	new_cb.BlackBishops = &bb
	new_cb.BlackRooks = &br
	new_cb.BlackQueens = &bq
	new_cb.BlackKing = &bk

	new_cb.White = cb.White
	new_cb.Black = cb.Black

	return new_cb
}



func Move_to_index(cord string) uint {
	// convert a chess cord to an index

	var ind uint

	cord = cord[0:2]
	
	letter := cord[0]
	number := cord[1]

	ind = uint(number-'1')*8 + uint(letter-'a')

	return ind
}

func Index_to_move(ind uint) string {

	var cord string

	letter := rune(ind%8) + 'a'
	rank := ind/8 + 1
	cord = string(letter) + strconv.Itoa( int(rank) )

	return cord
}

// check if two chessboards are identical
func (cb ChessBoard) Identical(new_cb ChessBoard) bool {

	if (*cb.WhitePawns == *new_cb.WhitePawns) &&
		(*cb.WhiteKnights == *new_cb.WhiteKnights) &&
		(*cb.WhiteBishops == *new_cb.WhiteBishops) &&
		(*cb.WhiteRooks == *new_cb.WhiteRooks) &&
		(*cb.WhiteQueens == *new_cb.WhiteQueens) &&
		(*cb.WhiteKing == *new_cb.WhiteKing) &&
		(*cb.BlackPawns == *new_cb.BlackPawns) &&
		(*cb.BlackKnights == *new_cb.BlackKnights) &&
		(*cb.BlackBishops == *new_cb.BlackBishops) &&
		(*cb.BlackRooks == *new_cb.BlackRooks) &&
		(*cb.BlackQueens == *new_cb.BlackQueens) &&
		(*cb.BlackKing == *new_cb.BlackKing) {
		return true
	} 
	
	return false
}