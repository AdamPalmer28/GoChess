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
	
}

// test castle move
func Test_castle_move(t *testing.T) {
	
}


// test promotion move (including capture)
func Test_promotion_move(t *testing.T) {

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