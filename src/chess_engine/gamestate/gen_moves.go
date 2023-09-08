package gamestate

import (
	"chess/chess_engine/board"
	"chess/chess_engine/move_gen"
	"chess/chess_engine/move_gen/magic"
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

	fwd int
	castle_rights uint
}

func (gs *GameState) Make_BP() {

	var bp BoardPerpective
	var king_safety move_gen.KingSafetyRelBB

	if gs.White_to_move {
		
		bp = BoardPerpective{
			

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

			fwd: 8,
			castle_rights: gs.WhiteCastle,
		}

	} else {

		bp = BoardPerpective{

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

			fwd: -8,
			castle_rights: gs.BlackCastle,
		}
	}

	// king safety struct
	opp_king_bubble := gs.MoveRays.KingRays[bp.opp_king_bb.Index()[0]]
	
	king_safety = move_gen.KingSafetyRelBB{
		King_sq: bp.king_bb.Index()[0],
		King_bb: bp.king_bb,
		Team_bb: bp.team_bb,
		Team_bb_no_king: (bp.pawn_bb | bp.knight_bb | 
				bp.bishop_bb | bp.rook_bb | bp.queen_bb),
		Opp_bb: bp.opp_bb,
		Fwd: bp.fwd,

		Opp_pawn_bb: bp.opp_pawn_bb,
		Opp_knight_bb: bp.opp_knight_bb,
		Opp_bishop_bb: bp.opp_bishop_bb,
		Opp_rook_bb: bp.opp_rook_bb,
		Opp_queen_bb: bp.opp_queen_bb,
		Opp_king_bubble: opp_king_bubble,
	}

	gs.PlayerBoard = bp
	gs.PlayerKingSaftey = king_safety
}

// Should add initial check to see if king is in check
// then only generate relevant moves


func (gs *GameState) GenMoves() {

	// reset the moves
	MoveList := move_gen.MoveList{}

	var player BoardPerpective = gs.PlayerBoard

	// generate pawn moves
	pawn_moves := move_gen.GenPawnMoves(player.pawn_bb, gs.White_to_move, 
				gs.Enpass_ind, player.team_bb, player.opp_bb)

	MoveList = append(MoveList, pawn_moves...)

	// generate knight moves
	knight_moves := move_gen.GenKnightMoves(player.knight_bb, 
				&gs.MoveRays.KnightRays, player.team_bb, player.opp_bb)

	MoveList = append(MoveList, knight_moves...)
	

	// generate rook moves
	rook_moves := magic.GenMagicMoves(player.rook_bb, &gs.MoveRays.Magic.RookMagic,
				player.team_bb, player.opp_bb)

	MoveList = append(MoveList, rook_moves...)


	// generate bishop moves
	bishop_moves := magic.GenMagicMoves(player.bishop_bb, &gs.MoveRays.Magic.BishopMagic,
				player.team_bb, player.opp_bb)
	
	MoveList = append(MoveList, bishop_moves...)


	// generate queen moves
	queen_moves := magic.GenMagicMoves(player.queen_bb, &gs.MoveRays.Magic.BishopMagic,
				player.team_bb, player.opp_bb)
	MoveList = append(MoveList, queen_moves...)

	queen_moves = magic.GenMagicMoves(player.queen_bb, &gs.MoveRays.Magic.RookMagic,
				player.team_bb, player.opp_bb)
	MoveList = append(MoveList, queen_moves...)



	// generate king moves
	
	

	king_moves := move_gen.GenKingMoves(gs.PlayerKingSaftey, player.castle_rights,
				&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.Magic.BishopMagic,
				&gs.MoveRays.KnightRays, &gs.MoveRays.KingRays)
	MoveList = append(MoveList, king_moves...)
	
	
	gs.MoveList = MoveList
}
