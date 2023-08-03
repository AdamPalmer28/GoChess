package chess_engine

import (
	"chess/chess_engine/gamestate"
)



func StartGame() *gamestate.GameState {

	println("Starting Chess Engine")

	// fen_str := "rnbqkbnr/pp2p1pp/2p2p2/3p4/2PP4/4PN2/PP3PPP/RNBQKB1R b KQkq - 1 4"
	// gs := gamestate.FEN_to_gs(fen_str)

	gs := gamestate.MakeGameState()
	
	
	return gs
}
