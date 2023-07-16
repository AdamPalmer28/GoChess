package chess_engine

import (
	"chess/board"
	"chess/chess_engine/make_game"
)



func StartGame() *board.ChessBoard{

	println("Starting Chess Engine")

	board := make_game.MakeIntialChessBoard()
	
	return board
}
