package gamestate

import (
	"chess/chess_engine/move_gen"
	"chess/chess_engine/move_gen/magic"
)

func (gs *GameState) Init() {

	// update
	gs.Board.UpdateSideBB(true)
	gs.Board.UpdateSideBB(false)

	
	// initialise move rays
	gs.makeMoveRays()
	
	// magic squares
	strt, diag := magic.Load_all_magicsq()

	gs.MoveRays.Magic.RookMagic = strt
	gs.MoveRays.Magic.BishopMagic = diag
	
	// next move
	gs.Next_move()


}


func (gs *GameState) makeMoveRays() {

	gs.MoveRays.KnightRays = move_gen.InitKnightRays()

	gs.MoveRays.KingRays = move_gen.InitKingRays()
}