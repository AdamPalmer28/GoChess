package move_gen

import (
	"chess/chess_engine/board"
	"chess/chess_engine/move_gen/magic"
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

// ----------------------------------------------------------------------------
// Xray moves of sliding pieces - i.e. the moves that are blocked by other pieces

func InitRookXRays() [64]board.Bitboard {

	moves := [64]board.Bitboard{}

	var i uint
	for i = 0; i < 64; i++ {
		moves[i] = magic.Fullrays(i, false)
	}

	return moves
}

func InitBishopXRays() [64]board.Bitboard {

	moves := [64]board.Bitboard{}

	var i uint
	for i = 0; i < 64; i++ {
		moves[i] = magic.Fullrays(i, true)
	}

	return moves
}