package test_move_gen

import (
	"chess/chess_engine"
	"testing"
)

func Test_InCheck_MoveGen(t *testing.T) {

	var actual_moves []uint
	var expected []uint
	var moves []uint
	var result bool

	// Single check 
	// ------------

	fen := "4k1r1/8/1q6/3NQ3/N6R/2B5/3PPK2/8 w - - 0 1"
	gs := chess_engine.CreateGameFen(fen)

	actual_moves = append(actual_moves, get_piece_moves(gs.MoveList, *gs.Board.WhiteKing)...)
	actual_moves = append(actual_moves, get_piece_moves(gs.MoveList, *gs.Board.WhitePawns)...)
	actual_moves = append(actual_moves, get_piece_moves(gs.MoveList, *gs.Board.WhiteQueens)...)
	actual_moves = append(actual_moves, get_piece_moves(gs.MoveList, *gs.Board.WhiteBishops)...)
	actual_moves = append(actual_moves, get_piece_moves(gs.MoveList, *gs.Board.WhiteRooks)...)
	actual_moves = append(actual_moves, get_piece_moves(gs.MoveList, *gs.Board.WhiteKnights)...)	

	// moves ------------------------------------------------

	// expected moves
	expected = []uint{}
	moves = create_moves([]string{"a4","c3","d5","e5","h4",
							"e2","f2"}, 
						[][]string{{"c5"},{"d4"},{"e3"},{"d4","e3"},{"d4"},
							{"e3"},{"f1","f3","e1"}}, 0)
	expected = append(expected, moves...)
	moves = create_moves([]string{"d2"}, [][]string{{"d4"}}, 0b0001) // double pawn push
	expected = append(expected, moves...)
	moves = create_moves([]string{"a4","d5"}, [][]string{{"b6"},{"b6"}}, 0b0100) // capture
	expected = append(expected, moves...)

	result = check_moves(actual_moves, expected)
	if !result {
		error_str := useful_error_msg(actual_moves, expected)
		t.Errorf("Single check: moves failed :\n%s", error_str)
	}
	
	// Double check
	// ------------

	fen = "4kr2/8/1q6/3NQ3/N6R/2B5/3PPK2/8 w - - 0 1"
	gs = chess_engine.CreateGameFen(fen)

	// moves ------------------------------------------------

	// expected moves
	expected = []uint{}
	moves = create_moves([]string{"f2"}, [][]string{{"e1","g3","g2"}}, 0)
	expected = append(expected, moves...)

	result = check_moves(gs.MoveList, expected)
	if !result {
		error_str := useful_error_msg(gs.MoveList, expected)
		t.Errorf("Double check: moves failed :\n%s", error_str)
	}

}