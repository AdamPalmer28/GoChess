package make_game

import "chess/board"

type GameState struct {
	Board board.ChessBoard

	WhiteCastle [2]bool
	BlackCastle [2]bool

	White_to_move bool
	Enpassent_ind uint

	Movehist []string
	Moveno   uint

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
