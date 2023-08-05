package gamestate

import (
	"chess/chess_engine/moves/move_gen"
)

func (gs *GameState) Init() {

	// update
	gs.Board.UpdateSideBB(true)
	gs.Board.UpdateSideBB(false)

	
	// initialise move rays
	gs.makeMoveRays()
	
	
	// move gen
	gs.GenMoves()

}


func (gs *GameState) makeMoveRays() {

	gs.MoveRays.KnightRays = move_gen.InitKnightRays()

	gs.MoveRays.KingRays = move_gen.InitKingRays()
}