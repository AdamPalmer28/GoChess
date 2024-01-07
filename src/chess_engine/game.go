package chess_engine

import (
	"chess/src/chess_engine/gamestate"
)



func StartGame() *gamestate.GameState {

	gs := gamestate.MakeGameState()
	gs.Init()
	
	return gs

}

func CreateGameFen(fen string) *gamestate.GameState {

	gs := gamestate.FEN_to_gs(fen)
	gs.Init()

	return gs


}




