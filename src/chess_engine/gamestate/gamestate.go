package gamestate

import (
	"chess/chess_engine/board"
	"chess/chess_engine/move_gen"
	"chess/chess_engine/move_gen/magic"
)

type GameState struct {
	Board board.ChessBoard

	WhiteCastle [2]bool
	BlackCastle [2]bool

	White_to_move bool
	Enpass_ind uint

	MoveList move_gen.MoveList
	MoveHumanList []string // CLI referencing

	Moveno   uint
	HalfMoveNo uint
	PrevMoves []uint // previous moves (0000 000000 000000 form)
	Cap_pieces [][2]int // [move number][piece type]

	MoveRays struct {
		// rays for each piece
		KnightRays [64]board.Bitboard
		BishopRays [64]board.Bitboard
		RookRays [64]board.Bitboard
		QueenRays [64]board.Bitboard
		KingRays [64]board.Bitboard

		Magic struct {
			// magic squares for each piece
			RookMagic [64]magic.Magicsq
			BishopMagic [64]magic.Magicsq
		}
	}

	PlayerBoard  BoardPerpective

}

var PieceValLookup = map[int]string{
	0: "P",
	1: "N",
	2: "B",
	3: "R",
	4: "Q",
	5: "K",
	6: "p",
	7: "n",
	8: "b",
	9: "r",
	10: "q",
	11: "k",
}

var PieceBBind = map[string]uint{
	"P": 0,
	"N": 1,
	"B": 2,
	"R": 3,
	"Q": 4,
	"K": 5,
	"p": 6,
	"n": 7,
	"b": 8,
	"r": 9,
	"q": 10,
	"k": 11,
}


