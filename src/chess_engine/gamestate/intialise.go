package gamestate

import (
	"chess/src/chess_engine/move_gen"
	"chess/src/chess_engine/move_gen/magic"
)

func (gs *GameState) Init() {

	// update board
	gs.Board.UpdateSideBB(true)
	gs.Board.UpdateSideBB(false)

		// piece locations
	for i := uint(0); i < 12; i++ {
		gs.Board.UpdatePieceLocations(i)
	}

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

	gs.MoveRays.PawnCapRays = move_gen.InitPawnCaptureRays()

	gs.MoveRays.RookRays = move_gen.InitRookXRays()

	gs.MoveRays.BishopRays = move_gen.InitBishopXRays()


}