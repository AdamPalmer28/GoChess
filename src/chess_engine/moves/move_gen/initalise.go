package move_gen

import (
	"chess/chess_engine/board"
)


func InitKnightRays() [64]board.Bitboard {

	moves := [64]board.Bitboard{}

	for i := 0; i < 64; i++ {
		moves[i] = KnightMoves(i)
	}

	return moves
}