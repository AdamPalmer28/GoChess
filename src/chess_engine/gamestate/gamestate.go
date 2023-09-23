package gamestate

import (
	"chess/chess_engine/board"
	"chess/chess_engine/move_gen"
	"chess/chess_engine/move_gen/magic"
)

type GameState struct {
	Board board.ChessBoard

	// king 0b01, queen 0b10, both 0b11 
	WhiteCastle uint 
	BlackCastle uint 
	
	// game state
	White_to_move bool
	InCheck bool
	GameOver bool
	
	// move list
	MoveList move_gen.MoveList
	Enpass_ind uint
	MoveHumanList []string // CLI referencing
	PlayerBoard  move_gen.BoardPerpective
	PlayerKingSaftey move_gen.KingSafetyRelBB

	Moveno   uint
	HalfMoveNo uint
	
	// history data - for undo
	History struct {
		PrevMoves []uint // previous moves (0000 000000 000000 form)
		EnPassHist []uint // enpassant history
		Cap_pieces []uint // History of piece type
		CastleRight []uint // castle rights at end of the move
	}
	
	// move arrays - for move gen
	MoveRays struct {
		// rays for each piece
		KnightRays [64]board.Bitboard
		KingRays [64]board.Bitboard
		PawnCapRays [2][64]board.Bitboard

		Magic struct {
			// magic squares for each piece
			RookMagic [64]magic.Magicsq
			BishopMagic [64]magic.Magicsq
		}
	}
}

// BoardPerpective struct
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



