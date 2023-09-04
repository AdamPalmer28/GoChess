package board

import (
	"testing"
)

// test quiet move
func Test_basic_move(t *testing.T) {

	var basic_move uint
	var exp_bb *Bitboard
	var piece_moved uint
	var cap_piece uint

	// create a board with all the pieces in their initial positions
	cb := make_inital_cb()

	// test double pawn push -------------------------------------------------
	basic_move = 0b0001_011100_001100 
	piece_moved, cap_piece = cb.Move(basic_move, true)

		// test correct
	exp_bb = make_bitboard([]uint{8,9,10,11,28,13,14,15})
	if piece_moved != 0 || cap_piece != 6 {
		t.Errorf("Double pawn push failed. \nPiece moved should be 6 (%d) Piece captured should be 0 (%d)", piece_moved, cap_piece) 
	}
	if *cb.WhitePawns != *exp_bb {
		t.Errorf("Double pawn push failed (expected bb followed by result)")
		exp_bb.Print()
		cb.WhitePawns.Print()
	}


	// test bishop move -------------------------------------------------------
	basic_move = 0b0000_011010_000101
	piece_moved, cap_piece = cb.Move(basic_move, true)

		// test correct
	exp_bb = make_bitboard([]uint{2,26})
	if piece_moved != 2 || cap_piece != 6 {
		t.Errorf("Bishop move failed. \nPiece moved should be 2 (%d) Piece captured should be 6 (%d)", piece_moved, cap_piece)
	}
	if *cb.WhiteBishops != *exp_bb {
		t.Errorf("Bishop move failed (expected bb followed by result)")
		exp_bb.Print()
		cb.WhiteBishops.Print()
	}

	// test knight move -------------------------------------------------------
	basic_move = 0b0000_010101_000110
	piece_moved, cap_piece = cb.Move(basic_move, true)

		// test correct
	exp_bb = make_bitboard([]uint{1,21})
	if piece_moved != 1 || cap_piece != 6 {
		t.Errorf("Knight move failed. \nPiece moved should be 1 (%d) Piece captured should be 6 (%d)", piece_moved, cap_piece)
	}
	if *cb.WhiteKnights != *exp_bb {
		t.Errorf("Knight move failed (expected bb followed by result)")
		exp_bb.Print()
		cb.WhiteKnights.Print()
	}
}

// test capture move (including enpassent)
func Test_capture_move(t *testing.T) {
	
	var basic_move uint
	var piece_moved uint
	var cap_piece uint

	// create a board with all the pieces in their initial positions
	cb := make_inital_cb()
	bpawns := make_bitboard([]uint{48,49,50,51,36,53,54,55})
	wpawns := make_bitboard([]uint{8,9,10,35,12,13,14,15})
	cb.BlackPawns = bpawns
	cb.WhitePawns = wpawns


	// normal capture --------------------------------------------------------
	basic_move = 0b0100_110001_001000
	piece_moved, cap_piece = cb.Move(basic_move, true)

		// test correct

	exp_bb_wp := make_bitboard([]uint{49,9,10,35,12,13,14,15})
	exp_bb_bp := make_bitboard([]uint{48,50,51,36,53,54,55})

	if piece_moved != 0 || cap_piece != 0 {
		t.Errorf("Normal capture failed. \nPiece moved should be 0 (%d) Piece captured should be 0 (%d)", piece_moved, cap_piece)
	}
	if *cb.WhitePawns != *exp_bb_wp {
		t.Errorf("Normal capture failed (expected bb followed by result)")
		exp_bb_wp.Print()
		cb.WhitePawns.Print()
	}
	if *cb.BlackPawns != *exp_bb_bp {
		t.Errorf("Normal capture failed (expected bb followed by result)")
		exp_bb_bp.Print()
		cb.BlackPawns.Print()
	}

	// enpassent capture -----------------------------------------------------
	basic_move = 0b0101_101100_100011
	piece_moved, cap_piece = cb.Move(basic_move, true)

		// test correct
	exp_bb_wp = make_bitboard([]uint{49,9,10,44,12,13,14,15})
	exp_bb_bp = make_bitboard([]uint{48,50,51,53,54,55})

	if piece_moved != 0 || cap_piece != 0 {
		t.Errorf("Enpassent capture failed. \nPiece moved should be 0 (%d) Piece captured should be 0 (%d)", piece_moved, cap_piece)
	}
	if *cb.WhitePawns != *exp_bb_wp {
		t.Errorf("Enpassent capture failed (expected bb followed by result)")
		exp_bb_wp.Print()
		cb.WhitePawns.Print()
	}
	if *cb.BlackPawns != *exp_bb_bp {
		t.Errorf("Enpassent capture failed (expected bb followed by result)")
		exp_bb_bp.Print()
		cb.BlackPawns.Print()
	}
}

// test castle move
func Test_castle_move(t *testing.T) {
	
	// create a board with all the pieces in their initial positions
	cb := make_inital_cb()
	cb.WhiteBishops = make_bitboard([]uint{})
	cb.WhiteKnights = make_bitboard([]uint{})
	cb.WhiteQueens = make_bitboard([]uint{})

	var castle_move uint

	// king side castle ------------------------------------------------------
	castle_move = 0b0010_000110_000100
	piece_moved, cap_piece := cb.Move(castle_move, true)

		// test correct
	//exp_bb_wk := make_bitboard([]uint{6})
	exp_bb_wr := make_bitboard([]uint{0,5})

	if piece_moved != 5 || cap_piece != 6 {
		t.Errorf("King side castle failed. \nPiece moved should be 5 (%d) Piece captured should be 6 (%d)", piece_moved, cap_piece)
	}
	if *cb.WhiteRooks != *exp_bb_wr {
		t.Errorf("King side castle failed (expected bb followed by result)")
		exp_bb_wr.Print()
		cb.WhiteRooks.Print()
	}

	// queen side castle -----------------------------------------------------
	cb.WhiteKing = make_bitboard([]uint{4})
	cb.WhiteRooks = make_bitboard([]uint{0,7})
	castle_move = 0b0011_000010_000100
	piece_moved, cap_piece = cb.Move(castle_move, true)

		// test correct
	//exp_bb_wk = make_bitboard([]uint{2})
	exp_bb_wr = make_bitboard([]uint{3,7})

	if piece_moved != 5 || cap_piece != 6 {
		t.Errorf("Queen side castle failed. \nPiece moved should be 5 (%d) Piece captured should be 6 (%d)", piece_moved, cap_piece)
	}
	if *cb.WhiteRooks != *exp_bb_wr {
		t.Errorf("Queen side castle failed (expected bb followed by result)")
		exp_bb_wr.Print()
		cb.WhiteRooks.Print()
	}
}


// test promotion move (including capture)
func Test_promotion_move(t *testing.T) {

	var piece_moved uint
	var cap_piece uint

	
	// promotion vars
	var knight_promo uint = 0b1000
	var bishop_promo uint = 0b1001
	var rook_promo uint = 0b1010
	var queen_promo uint = 0b1011
	promo_list := []uint{knight_promo, bishop_promo, rook_promo, queen_promo}

	var queen_promo_cap uint = 0b1111
	// promotion -------------------------------------------------------------

	var demo_move uint = 0b_010000_001000
	for i, promo := range promo_list {
		// create a board with all the pieces in their initial positions
		cb := make_inital_cb()
		cb.WhiteKnights = make_bitboard([]uint{})
		cb.WhiteBishops = make_bitboard([]uint{})
		cb.WhiteRooks = make_bitboard([]uint{})
		cb.WhiteQueens = make_bitboard([]uint{})

		promo_move := demo_move | (promo << 12)

		piece_moved, cap_piece = cb.Move(promo_move, true)

		// test correct
		bb_set := cb.ListBB(true)
		exp_bb := make_bitboard([]uint{16})
		exp_pawn_bb := make_bitboard([]uint{9,10,11,12,13,14,15})

		if piece_moved != 0 || cap_piece != 6 {
			t.Errorf("Promotion failed. \nPiece moved should be 0 (%d) Piece captured should be 6 (%d)", piece_moved, cap_piece)
		}
		if *bb_set[i+1] != *exp_bb {
			t.Errorf("Promotion failed - Promo piece BB (expected bb followed by result)")
			exp_bb.Print()
			bb_set[i+1].Print()
		}
		if *cb.WhitePawns != *exp_pawn_bb {
			t.Errorf("Promotion failed - Pawn BB (expected bb followed by result)")
			exp_pawn_bb.Print()
			cb.WhitePawns.Print()
		}
	}		

	// promotion capture ------------------------------------------------------

	// create a board with all the pieces in their initial positions
	cb := make_inital_cb()
	cb.BlackQueens = make_bitboard([]uint{16})

	promo_move := demo_move | (queen_promo_cap << 12)

	piece_moved, cap_piece = cb.Move(promo_move, true)

		// test correct
	exp_bb := make_bitboard([]uint{16,3})
	exp_pawn_bb := make_bitboard([]uint{9,10,11,12,13,14,15})
	exp_opp_q_bb := make_bitboard([]uint{})

	if piece_moved != 0 || cap_piece != 4 {
		t.Errorf("Promotion capture failed. \nPiece moved should be 0 (%d) Piece captured should be 4 (%d)", piece_moved, cap_piece)
	}
	if *cb.WhiteQueens != *exp_bb {
		t.Errorf("Promotion capture failed - Queen BB (expected bb followed by result)")
		exp_bb.Print()
		cb.WhiteQueens.Print()
	}
	if *cb.BlackQueens != *exp_opp_q_bb {
		t.Errorf("Promotion capture failed - Opp Queen BB (expected bb followed by result)")
		exp_opp_q_bb.Print()
		cb.BlackQueens.Print()
	}
	if *cb.WhitePawns != *exp_pawn_bb {
		t.Errorf("Promotion capture failed - Pawn BB (expected bb followed by result)")
		exp_pawn_bb.Print()
		cb.WhitePawns.Print()
	}

}



// ========================================================================
// helper function for testing

func make_inital_cb() *ChessBoard {
	// create a board with all the pieces in their initial positions
	var cb ChessBoard

	cb.WhitePawns = make_bitboard([]uint{8,9,10,11,12,13,14,15})
	cb.WhiteKnights = make_bitboard([]uint{1,6})
	cb.WhiteBishops = make_bitboard([]uint{2,5})
	cb.WhiteRooks = make_bitboard([]uint{0,7})
	cb.WhiteQueens = make_bitboard([]uint{3})
	cb.WhiteKing = make_bitboard([]uint{4})

	cb.BlackPawns = make_bitboard([]uint{48,49,50,51,52,53,54,55})
	cb.BlackKnights = make_bitboard([]uint{57,62})
	cb.BlackBishops = make_bitboard([]uint{58,61})
	cb.BlackRooks = make_bitboard([]uint{56,63})
	cb.BlackQueens = make_bitboard([]uint{59})
	cb.BlackKing = make_bitboard([]uint{60})

	cb.UpdateSideBB(true)
	cb.UpdateSideBB(false)

	return &cb
}

func make_bitboard(ind []uint) *Bitboard {
	var BB Bitboard
	for _, i := range ind {
		BB.flip(i)
	}
	return &BB
}