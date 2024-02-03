package src

import (
	"chess/src/chess_engine"
	"chess/src/chess_engine/gamestate"
)

type GameHost struct {

	GameState *gamestate.GameState

	// 
}


// Server Host of the chess gamestate
func StartGameHost() *GameHost {

	gamestate.InitZobrist() // init zobrist keys

	// start the game
	gs := chess_engine.StartGame()
	
	gh := &GameHost{
		GameState: gs,
	}

	return gh
}

