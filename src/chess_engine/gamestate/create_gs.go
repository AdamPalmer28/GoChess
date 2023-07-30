package gamestate

import (
	"chess/chess_engine/board"
)
func MakeGameState() *GameState {
	// Make a new game state

	chess_board	 := MakeIntialChessBoard()
	chess_board.UpdateSideBB(true)

	gs := &GameState{
		Board: *chess_board,
		WhiteCastle: [2]bool{true, true}, // [0] = queenside, [1] = kingside
		BlackCastle: [2]bool{true, true}, // [0] = queenside, [1] = kingside
		White_to_move: true,
		Enpassent_ind: 0,
		MoveList: []uint{},
		MoveHumanList: []string{},
		Moveno: 0,
		PrevMoves: []uint{},
		Cap_pieces: [][2]int{},
	}

	return gs
}

func MakeIntialChessBoard() *board.ChessBoard {
	// Starting chess position

	WP := board.Rank2
	WN := (board.FileB | board.FileG) & board.Rank1
	WB := (board.FileC | board.FileF) & board.Rank1
	WR := (board.FileA | board.FileH) & board.Rank1
	WQ := board.FileD & board.Rank1
	WK := board.FileE & board.Rank1
	BP := board.Rank7
	BN := (board.FileB | board.FileG) & board.Rank8
	BB := (board.FileC | board.FileF) & board.Rank8
	BR := (board.FileA | board.FileH) & board.Rank8
	BQ := board.FileD & board.Rank8
	BK := board.FileE & board.Rank8

	
	board := &board.ChessBoard{
		WhitePawns:   &WP,
		WhiteKnights: &WN,
		WhiteBishops: &WB,
		WhiteRooks:   &WR,
		WhiteQueens:  &WQ,
		WhiteKing:    &WK,
		BlackPawns:   &BP,
		BlackKnights: &BN,
		BlackBishops: &BB,
		BlackRooks:   &BR,
		BlackQueens:  &BQ,
		BlackKing:    &BK,
	}
	return board
}