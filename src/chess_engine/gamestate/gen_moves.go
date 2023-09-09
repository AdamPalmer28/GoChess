package gamestate

import (
	"chess/chess_engine/move_gen"
	"chess/chess_engine/move_gen/magic"
)

func (gs *GameState) Make_BP() {

	var bp move_gen.BoardPerpective
	var king_safety move_gen.KingSafetyRelBB

	if gs.White_to_move {
		
		bp = move_gen.BoardPerpective{
			
			Pawn_bb: *gs.Board.WhitePawns,
			Knight_bb: *gs.Board.WhiteKnights,
			Bishop_bb: *gs.Board.WhiteBishops,
			Rook_bb: *gs.Board.WhiteRooks,
			Queen_bb: *gs.Board.WhiteQueens,
			King_bb: *gs.Board.WhiteKing,
			Opp_pawn_bb: *gs.Board.BlackPawns,
			Opp_knight_bb: *gs.Board.BlackKnights,
			Opp_bishop_bb: *gs.Board.BlackBishops,
			Opp_rook_bb: *gs.Board.BlackRooks,
			Opp_queen_bb: *gs.Board.BlackQueens,
			Opp_king_bb: *gs.Board.BlackKing,
			Team_bb: gs.Board.White,
			Opp_bb: gs.Board.Black,

			Fwd: 8,
			P_start_row: 1,
			Castle_rights: gs.WhiteCastle,
			Enpass_ind: gs.Enpass_ind,
		}

	} else {

		bp = move_gen.BoardPerpective{

			Pawn_bb: *gs.Board.BlackPawns,
			Knight_bb: *gs.Board.BlackKnights,
			Bishop_bb: *gs.Board.BlackBishops,
			Rook_bb: *gs.Board.BlackRooks,
			Queen_bb: *gs.Board.BlackQueens,
			King_bb: *gs.Board.BlackKing,
			Opp_pawn_bb: *gs.Board.WhitePawns,
			Opp_knight_bb: *gs.Board.WhiteKnights,
			Opp_bishop_bb: *gs.Board.WhiteBishops,
			Opp_rook_bb: *gs.Board.WhiteRooks,
			Opp_queen_bb: *gs.Board.WhiteQueens,
			Opp_king_bb: *gs.Board.WhiteKing,
			Team_bb: gs.Board.Black,
			Opp_bb: gs.Board.White,

			Fwd: -8,
			P_start_row: 6,
			Castle_rights: gs.BlackCastle,
			Enpass_ind: gs.Enpass_ind,
		}
	}

	// king safety struct
	opp_king_bubble := gs.MoveRays.KingRays[bp.Opp_king_bb.Index()[0]]
	
	king_safety = move_gen.KingSafetyRelBB{
		King_sq: bp.King_bb.Index()[0],
		King_bb: bp.King_bb,
		Team_bb: bp.Team_bb,
		Team_bb_no_king: bp.Team_bb & ^bp.King_bb,
		Fwd: bp.Fwd,
		P_start_row: bp.P_start_row,
		Opp_bb: bp.Opp_bb,

		Opp_pawn_bb: bp.Opp_pawn_bb,
		Opp_knight_bb: bp.Opp_knight_bb,
		Opp_bishop_bb: bp.Opp_bishop_bb,
		Opp_rook_bb: bp.Opp_rook_bb,
		Opp_queen_bb: bp.Opp_queen_bb,
		Opp_king_bubble: opp_king_bubble,
	}

	gs.PlayerBoard = bp
	gs.PlayerKingSaftey = king_safety
}


// ----------------------------------------------------------------------------

// generate moves when king not in check
func (gs *GameState) GenMoves() {

	// reset the moves
	MoveList := move_gen.MoveList{}

	var player move_gen.BoardPerpective = gs.PlayerBoard

	// generate pawn moves
	pawn_moves := move_gen.GenPawnMoves(player.Pawn_bb, gs.White_to_move, 
				gs.Enpass_ind, player.Team_bb, player.Opp_bb)

	MoveList = append(MoveList, pawn_moves...)

	// generate knight moves
	knight_moves := move_gen.GenKnightMoves(player.Knight_bb, 
				&gs.MoveRays.KnightRays, player.Team_bb, player.Opp_bb)

	MoveList = append(MoveList, knight_moves...)
	

	// generate rook moves
	rook_moves := magic.GenMagicMoves(player.Rook_bb, &gs.MoveRays.Magic.RookMagic,
				player.Team_bb, player.Opp_bb)

	MoveList = append(MoveList, rook_moves...)


	// generate bishop moves
	bishop_moves := magic.GenMagicMoves(player.Bishop_bb, &gs.MoveRays.Magic.BishopMagic,
				player.Team_bb, player.Opp_bb)
	
	MoveList = append(MoveList, bishop_moves...)


	// generate queen moves
	queen_moves := magic.GenMagicMoves(player.Queen_bb, &gs.MoveRays.Magic.BishopMagic,
				player.Team_bb, player.Opp_bb)
	MoveList = append(MoveList, queen_moves...)

	queen_moves = magic.GenMagicMoves(player.Queen_bb, &gs.MoveRays.Magic.RookMagic,
				player.Team_bb, player.Opp_bb)
	MoveList = append(MoveList, queen_moves...)



	// generate king moves
	king_moves := move_gen.GenKingMoves(gs.PlayerKingSaftey, player.Castle_rights,
				&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.Magic.BishopMagic,
				&gs.MoveRays.KnightRays, &gs.MoveRays.KingRays)
	MoveList = append(MoveList, king_moves...)
	
	
	gs.MoveList = MoveList
}

// generate moves when king in check
func (gs *GameState) GenCheckMoves() {

	var MoveList []uint

	king_sq := gs.PlayerKingSaftey.King_sq
	// get threat details
	threats, threat_path := move_gen.CheckDetails(gs.PlayerKingSaftey, gs.MoveRays.KnightRays[king_sq], 
				&gs.MoveRays.Magic.RookMagic[king_sq],&gs.MoveRays.Magic.BishopMagic[king_sq])
	
	// threats - must be captures
	// threat_path - must be blocks (i.e moved into)

	if len(threats) == 1 { // calculate blockers

		threat_sq := threats[0]

		defender_moves := move_gen.DefenderMoves(threat_sq, threat_path, 
			gs.PlayerBoard, &gs.MoveRays.KnightRays,
			&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.Magic.BishopMagic)

		MoveList = append(MoveList, defender_moves...)

	} 

	king_moves := move_gen.GenKingMoves(gs.PlayerKingSaftey, 0,
		&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.Magic.BishopMagic,
		&gs.MoveRays.KnightRays, &gs.MoveRays.KingRays)
	MoveList = append(MoveList, king_moves...)

	gs.MoveList = MoveList
}