package gamestate

import "chess/chess_engine/board"

type GameState struct {
	Board board.ChessBoard

	WhiteCastle [2]bool
	BlackCastle [2]bool

	White_to_move bool
	Enpassent_ind uint

	MoveList []uint
	MoveHumanList []string // CLI referencing

	Moveno   uint
	HalfMoveNo uint
	PrevMoves []uint // previous moves (0000 000000 000000 form)
	Cap_pieces [][2]int // [move number][piece type]

}

var PieceValLookup = map[int]string{
	0: "P",
	1: "K",
	2: "B",
	3: "R",
	4: "Q",
	5: "K",
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


