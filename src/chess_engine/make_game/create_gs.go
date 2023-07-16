package make_game

import (
	"chess/board"
)

type GameState struct {
	Board board.ChessBoard

	WhiteCastle [2]bool
	BlackCastle [2]bool

	White_to_move bool
	Enpassent_ind uint

	Movehist []string
	Moveno uint

	Cap_pieces [][2]int // [move number][piece type]

}


func MakeIntialChessBoard() *board.ChessBoard {
	// Starting chess position
	
	board := &board.ChessBoard{
		WhitePawns:   board.Rank2,
		WhiteKnights: (board.FileB | board.FileG) & board.Rank1,
		WhiteBishops: (board.FileC | board.FileF) & board.Rank1,
		WhiteRooks:   (board.FileA | board.FileH) & board.Rank1,
		WhiteQueens:  board.FileD & board.Rank1,
		WhiteKing:    board.FileE & board.Rank1,
		BlackPawns:   board.Rank7,
		BlackKnights: (board.FileB | board.FileG) & board.Rank8,
		BlackBishops: (board.FileC | board.FileF) & board.Rank8,
		BlackRooks:   (board.FileA | board.FileH) & board.Rank8,
		BlackQueens:  board.FileD & board.Rank8,
		BlackKing:    board.FileE & board.Rank8,
	}
	return board
}