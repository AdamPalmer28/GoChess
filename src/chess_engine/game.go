package chess_engine

import (
	"chess/chess_engine/make_game"
)



func StartGame() *make_game.GameState {

	println("Starting Chess Engine")

	gs := make_game.MakeGameState()
	
	return gs
}
