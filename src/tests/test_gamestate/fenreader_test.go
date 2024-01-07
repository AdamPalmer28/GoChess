package test_gamestate

import (
	"chess/src/chess_engine/board"
	"chess/src/chess_engine/gamestate"
	"testing"
)

func Test_FEN_to_gs(t *testing.T) {

	fen := "rnb1k2r/1pp2ppp/2q2nb1/p1Ppp3/3PP2P/1B1N1B2/PP3PP1/R2QK1NR b KQkq e3 1 32"
	gs := *(gamestate.FEN_to_gs(fen))

	// player to move
	if gs.White_to_move {
		t.Error("Should be black to move")
	}
	// castle rights
	if gs.WhiteCastle != 0b11 {
		t.Error("White castle rights incorrect")
	}
	if gs.BlackCastle != 0b11 {
		t.Error("Black castle rights incorrect")
	}

	// enpassant
	if gs.Enpass_ind != 20 {
		t.Errorf("Enpassant square incorrect, got %d, expected 20", gs.Enpass_ind)
	}

	// move counts
	if gs.Moveno != 32 {
		t.Errorf("Move number incorrect, got %d, expected 32", gs.Moveno)
	}
	if gs.HalfMoveNo != 1 {
		t.Errorf("Half move number incorrect, got %d, expected 1", gs.HalfMoveNo)
	}

	// check the board
	
	// white pieces
	wp := board.Make_bitboard([]uint{8, 9, 34, 27, 28, 13, 14, 31})
	wn := board.Make_bitboard([]uint{19, 6})
	wb := board.Make_bitboard([]uint{17, 21})
	wr := board.Make_bitboard([]uint{0, 7})
	wq := board.Make_bitboard([]uint{3})
	wk := board.Make_bitboard([]uint{4})

	// black pieces
	bp := board.Make_bitboard([]uint{32, 49, 50, 35, 36, 53, 54, 55})
	bn := board.Make_bitboard([]uint{57, 45})
	bb := board.Make_bitboard([]uint{58, 46})
	br := board.Make_bitboard([]uint{56, 63})
	bq := board.Make_bitboard([]uint{42})
	bk := board.Make_bitboard([]uint{60})

	// create chessboard
	exp_cb := board.ChessBoard{
		WhitePawns:   wp,
		WhiteKnights: wn,
		WhiteBishops: wb,
		WhiteRooks:   wr,
		WhiteQueens:  wq,
		WhiteKing:    wk,

		BlackPawns:   bp,
		BlackKnights: bn,
		BlackBishops: bb,
		BlackRooks:   br,
		BlackQueens:  bq,
		BlackKing:    bk,

		White:        *wp | *wn | *wb | *wr | *wq | *wk,
		Black:        *bp | *bn | *bb | *br | *bq | *bk,
	}

	// check the board
	if gs.Board.Identical(exp_cb) == false {
		t.Error("Board incorrect")
	}
}
