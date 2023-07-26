package chess_engine

import (
	"chess/chess_engine/gamestate"
)



func StartGame() *gamestate.GameState {

	println("Starting Chess Engine")

	gs := gamestate.MakeGameState()
	
	return gs
}
