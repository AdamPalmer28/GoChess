package gamestate

import (
	"chess/chess_engine/board"
	"chess/chess_engine/moves/move_gen"
)

func (gs *GameState) GenMoves() {

	// reset the moves
	MoveList := []uint{}

	var pawn_bb board.Bitboard
	var knight_bb board.Bitboard
	var bishop_bb board.Bitboard
	var rook_bb board.Bitboard
	var queen_bb board.Bitboard
	var king_bb board.Bitboard

	var team_bb board.Bitboard
	var opp_bb board.Bitboard

	if gs.White_to_move {
		pawn_bb = *gs.Board.WhitePawns
		knight_bb = *gs.Board.WhiteKnights
		bishop_bb = *gs.Board.WhiteBishops
		rook_bb = *gs.Board.WhiteRooks
		queen_bb = *gs.Board.WhiteQueens
		king_bb = *gs.Board.WhiteKing

		team_bb = gs.Board.White
		opp_bb = gs.Board.Black
	} else {
		pawn_bb = *gs.Board.BlackPawns
		knight_bb = *gs.Board.BlackKnights
		bishop_bb = *gs.Board.BlackBishops
		rook_bb = *gs.Board.BlackRooks
		queen_bb = *gs.Board.BlackQueens
		king_bb = *gs.Board.BlackKing

		team_bb = gs.Board.Black
		opp_bb = gs.Board.White
	}

	// generate pawn moves
	pawn_moves := move_gen.GenPawnMoves(pawn_bb, gs.White_to_move, 
				gs.Enpass_ind, team_bb, opp_bb)

	MoveList = append(MoveList, pawn_moves...)

	// generate knight moves
	knight_moves := move_gen.GenKnightMoves(knight_bb, 
				&gs.MoveRays.KnightRays, team_bb, opp_bb)

	MoveList = append(MoveList, knight_moves...)

	// generate bishop moves


	// generate rook moves


	// generate queen moves


	// generate king moves


	// temp - to stop error
	bishop_bb.Index()
	rook_bb.Index()
	queen_bb.Index()
	king_bb.Index()

	gs.MoveList = MoveList

}

