package move_gen

import (
	"chess/chess_engine/board"
)


func InitKnightMoves() []board.Bitboard {

	moves := make([]board.Bitboard, 64)

	for i := 0; i < 64; i++ {
		moves[i] = KnightMoves(i)
	}

	return moves
}