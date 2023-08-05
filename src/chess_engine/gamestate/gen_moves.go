package gamestate

import (
	"chess/chess_engine/board"
	"chess/chess_engine/move_gen"
)

type BoardPerpective struct {
	// friendly pieces
	pawn_bb board.Bitboard
	knight_bb board.Bitboard
	bishop_bb board.Bitboard
	rook_bb board.Bitboard
	queen_bb board.Bitboard
	king_bb board.Bitboard
	// opp pieces
	opp_pawn_bb board.Bitboard
	opp_knight_bb board.Bitboard
	opp_bishop_bb board.Bitboard
	opp_rook_bb board.Bitboard
	opp_queen_bb board.Bitboard
	opp_king_bb board.Bitboard

	team_bb board.Bitboard
	opp_bb board.Bitboard
}

func (gs *GameState) GenMoves() {

	// reset the moves
	MoveList := move_gen.MoveList{}

	var player BoardPerpective

	if gs.White_to_move {

		player = BoardPerpective{

			pawn_bb: *gs.Board.WhitePawns,
			knight_bb: *gs.Board.WhiteKnights,
			bishop_bb: *gs.Board.WhiteBishops,
			rook_bb: *gs.Board.WhiteRooks,
			queen_bb: *gs.Board.WhiteQueens,
			king_bb: *gs.Board.WhiteKing,

			opp_pawn_bb: *gs.Board.BlackPawns,
			opp_knight_bb: *gs.Board.BlackKnights,
			opp_bishop_bb: *gs.Board.BlackBishops,
			opp_rook_bb: *gs.Board.BlackRooks,
			opp_queen_bb: *gs.Board.BlackQueens,
			opp_king_bb: *gs.Board.BlackKing,

			team_bb: gs.Board.White,
			opp_bb: gs.Board.Black,
		}

	} else {
		player = BoardPerpective{

			pawn_bb: *gs.Board.BlackPawns,
			knight_bb: *gs.Board.BlackKnights,
			bishop_bb: *gs.Board.BlackBishops,
			rook_bb: *gs.Board.BlackRooks,
			queen_bb: *gs.Board.BlackQueens,
			king_bb: *gs.Board.BlackKing,

			opp_pawn_bb: *gs.Board.WhitePawns,
			opp_knight_bb: *gs.Board.WhiteKnights,
			opp_bishop_bb: *gs.Board.WhiteBishops,
			opp_rook_bb: *gs.Board.WhiteRooks,
			opp_queen_bb: *gs.Board.WhiteQueens,
			opp_king_bb: *gs.Board.WhiteKing,

			team_bb: gs.Board.Black,
			opp_bb: gs.Board.White,
		}
	}

	// generate pawn moves
	pawn_moves := move_gen.GenPawnMoves(player.pawn_bb, gs.White_to_move, 
				gs.Enpass_ind, player.team_bb, player.opp_bb)

	MoveList = append(MoveList, pawn_moves...)


	// generate knight moves
	knight_moves := move_gen.GenKnightMoves(player.knight_bb, 
				&gs.MoveRays.KnightRays, player.team_bb, player.opp_bb)

	MoveList = append(MoveList, knight_moves...)

	// generate bishop moves


	// generate rook moves


	// generate queen moves


	// generate king moves
	king_basicmoves := move_gen.GenBasicKingMove(player.king_bb, &gs.MoveRays.KingRays, 
			player.team_bb, player.opp_bb, player.opp_king_bb)
	
	MoveList = append(MoveList, king_basicmoves...)
	

	gs.MoveList = MoveList

}
