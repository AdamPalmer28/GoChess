package make_game

import (
	"chess/board"
)

type GameState struct {
	Board board.ChessBoard

	WhiteKingCastle [2]bool
	BlackKingCastle [2]bool

	WhiteToMove bool
	EnPassantSquare uint

	MoveList []string
	MoveNumber uint

	CapturedPieces [][2]int // [move number][piece type]

}


var PieceValLookup = map[int]string{
	0 : "P",
	1 : "K",
	2 : "B",
	3 : "R",
	4 : "Q",
	5 : "K",
}

