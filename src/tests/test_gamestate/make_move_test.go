package test_gamestate

import (
	"chess/chess_engine/board"
	"chess/chess_engine/gamestate"
	"testing"
)

/*
Make initial gamestate
Make a sequence of moves
	Veryify:
		- piece bitboards
		- castle rights
		- enpassent index
		- move number
		- history
			- move list
			- capture list
			- castle rights
*/

func Test_Make_Move(t *testing.T) {

	fen := "rnbqkbnr/1p4p1/8/8/8/8/1P4P1/RNBQKBNR w KQkq - 0 1"
	gs := gamestate.FEN_to_gs(fen)

	var move_num uint 

	// 1 - move pawn double fwd
	move_num = 0b0001_011001_001001
	gs.Make_move(move_num)

	// test first move
	exp_bb := board.Make_bitboard([]uint{25, 14})
	if *gs.Board.WhitePawns != *exp_bb { // piece bitboards
		t.Errorf("Pawn bitboard not correct - 1st move")
	}
	if gs.Enpass_ind != 17 { // enpassent index
		t.Errorf("Enpass index not correct - got %d (expected 17)", gs.Enpass_ind)
	}
	if gs.Moveno != 2 { // move number
		t.Errorf("Move number not correct - got %d (expected 2)", gs.Moveno)
	}
	if gs.White_to_move != false { // player to move
		t.Errorf("Should be black to move - 1st move")
	}
	if (gs.WhiteCastle != 0b11) && (gs.BlackCastle != 0b11) { // castle rights
		t.Errorf("Castle rights not correct - 1st move")
	}
	if gs.History.PrevMoves[0] != move_num { // move history
		t.Errorf("Move list not correct - 1st move")
	}
	if gs.History.CastleRight[0] != 0b11 { // castle right history
		t.Errorf("CastleRight not correct - 1st move")
	}
	if gs.History.Cap_pieces[0] != 6 { // capture history
		t.Errorf("Capture history not correct - 1st move")
	}

	// 2 - rook moves a8 to a6
	move_num = 0b0000_101000_111000
	gs.Make_move(move_num)

	// test second move
	exp_bb = board.Make_bitboard([]uint{40, 63})
	if *gs.Board.BlackRooks != *exp_bb { // piece bitboards
		t.Errorf("Rook bitboard not correct - 2nd move")
	}
	if gs.Enpass_ind != 64 { // enpassent index
		t.Errorf("Enpass index not correct - got %d (expected 64)", gs.Enpass_ind)
	}
	if gs.Moveno != 3 { // move number
		t.Errorf("Move number not correct - got %d (expected 3)", gs.Moveno)
	}
	if gs.White_to_move != true { // player to move
		t.Errorf("Should be white to move - 2nd move")
	}
	if (gs.WhiteCastle != 0b11) && (gs.BlackCastle != 0b01) { // castle rights		
		t.Errorf("Castle rights not correct (expected: 0b10) - 2nd move")
	}
	if gs.History.PrevMoves[1] != move_num { // move history
		t.Errorf("Move list not correct - 2nd move")
	}
	if gs.History.CastleRight[1] != 0b11 { // castle right history
		t.Errorf("CastleRight not correct - 2nd move")
	}
	if gs.History.Cap_pieces[1] != 6 { // capture history
		t.Errorf("Capture history not correct - 2nd move")
	}

	// 3 rook capture h1 to h8
	move_num = 0b0100_111111_000111
	gs.Make_move(move_num)

	// test third move
	exp_bb = board.Make_bitboard([]uint{40})
	if *gs.Board.BlackRooks != *exp_bb { // piece bitboards
		t.Errorf("Rook bitboard not correct - 3rd move")
	}
	if gs.Enpass_ind != 64 { // enpassent index
		t.Errorf("Enpass index not correct - got %d (expected 64)", gs.Enpass_ind)
	}
	if gs.Moveno != 4 { // move number
		t.Errorf("Move number not correct - got %d (expected 4)", gs.Moveno)
	}
	if gs.White_to_move != false { // player to move
		t.Errorf("Should be black to move - 3rd move")
	}
	if (gs.WhiteCastle != 0b10) && (gs.BlackCastle != 0b00) { // castle rights
		t.Errorf("Castle rights not correct - 3rd move")
	}
	if gs.History.PrevMoves[2] != move_num { // move history
		t.Errorf("Move list not correct - 3rd move")
	}
	if gs.History.CastleRight[2] != 0b11 { // castle right history
		t.Errorf("CastleRight not correct - 3rd move")
	}
	if gs.History.Cap_pieces[2] != 3 { // capture history
		t.Errorf("Capture history not correct - 3rd move")
	}

	// ------------------------------------------------------------------------

	// new gamestate
	fen = "rnbqkbnr/1p4p1/8/8/8/8/1P4P1/RNBQKBNR w KQkq - 0 1"
	gs = gamestate.FEN_to_gs(fen)

	// king moves e1 to e2
	move_num = 0b0000_001100_000100
	gs.Make_move(move_num)

	// test first move
	if gs.Enpass_ind != 64 { // enpassent index
		t.Errorf("Enpass index not correct - got %d (expected 17)", gs.Enpass_ind)
	}
	if gs.Moveno != 2 { // move number
		t.Errorf("Move number not correct - got %d (expected 2)", gs.Moveno)
	}
	if gs.White_to_move != false { // player to move
		t.Errorf("Should be black to move - king move")
	}
	if (gs.WhiteCastle != 0b00) && (gs.BlackCastle != 0b11) { // castle rights
		t.Errorf("Castle rights not correct - king move")
	}
	if gs.History.PrevMoves[0] != move_num { // move history
		t.Errorf("Move list not correct - king move")
	}
	if gs.History.CastleRight[0] != 0b11 { // castle right history
		t.Errorf("CastleRight not correct - king move")
	}
	if gs.History.Cap_pieces[0] != 6 { // capture history
		t.Errorf("Capture history not correct - king move")
	}

}