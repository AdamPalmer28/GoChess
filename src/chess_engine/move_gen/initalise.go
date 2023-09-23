package move_gen

import (
	"chess/chess_engine/board"
)


func InitKnightRays() [64]board.Bitboard {

	moves := [64]board.Bitboard{}

	for i := 0; i < 64; i++ {
		moves[i] = KnightRays(i)
	}

	return moves
}

func InitKingRays() [64]board.Bitboard {

	moves := [64]board.Bitboard{}

	for i := 0; i < 64; i++ {
		moves[i] = KingRays(i)
	}

	return moves
}

func InitPawnCaptureRays() [2][64]board.Bitboard {

	moves := [2][64]board.Bitboard{}

	// white
	for i := 0; i < 64; i++ {
		moves[0][i] = get_pawn_attack(uint(i), 8)
	}

	// black
	for i := 0; i < 64; i++ {
		moves[1][i] = get_pawn_attack(uint(i), -8)
	}

	return moves
}