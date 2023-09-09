package test_move_gen

import (
	"chess/chess_engine"
	"testing"
)

/*
Pawn tests:

- standard moves
- capture
- en passant
- promotion
- promotion capture

*/


func TestPawnStandardMove(t *testing.T) {

	var expected []uint

	fen := "rnbqkb1r/1p1p1p2/2p1p1n1/8/4B3/1N3N2/PP2PP2/R1BQK2R w KQkq - 0 1"
	gs := chess_engine.CreateGameFen(fen)

	// get all the pawn moves
	pawn_moves := get_piece_moves(gs.MoveList, *gs.Board.WhitePawns)

	// expected moves
	single_push := create_moves([]string{"a2", "e2"}, [][]string{{"a3"}, {"e3"}}, 0)
	expected = append(expected, single_push...)

	double_push := create_moves([]string{"a2"}, [][]string{{"a4"}}, 1)
	expected = append(expected, double_push...)

	// check moves
	result := check_moves(pawn_moves, expected)
	if !result {
		error_str := useful_error_msg(pawn_moves, expected)
		t.Errorf("Standard_moves failed :\n%s", error_str)
	}

	// make move - now black to move
	gs.Make_move(expected[0])

	// get all moves
	pawn_moves = get_piece_moves(gs.MoveList, *gs.Board.BlackPawns)

	// expected moves
	single_push = create_moves([]string{"b7", "c6", "d7", "e6", "f7"}, 
							[][]string{{"b6"}, {"c5"}, {"d6"}, {"e5"}, {"f6"}}, 0)
	expected = single_push

	// double push
	double_push = create_moves([]string{"b7", "d7", "f7"}, [][]string{{"b5"}, {"d5"}, {"f5"}}, 1)
	expected = append(expected, double_push...)

	// check moves
	result = check_moves(pawn_moves, expected)
	if !result {
		error_str := useful_error_msg(pawn_moves, expected)
		t.Errorf("Standard_moves failed :\n%s", error_str)
	}
}


func TestPawnCapture(t *testing.T) {

	var expected []uint

	fen := "rnbqkb1r/3p4/6n1/2p1Pp2/pP1NP3/PP5p/3B3P/R1BQK2R w KQkq - 0 1"
	gs := chess_engine.CreateGameFen(fen)

	// get all the pawn moves
	pawn_moves := get_piece_moves(gs.MoveList, *gs.Board.WhitePawns)

	// expected moves
	capture := create_moves([]string{"e4", "b3", "b4"}, 
						[][]string{{"f5"}, {"a4"}, {"c5"}}, 0b0100)
	expected = append(expected, capture...)

	normal_moves := create_moves([]string{"e5", "b4"}, 
							[][]string{{"e6"}, {"b5"}}, 0)
	expected = append(expected, normal_moves...)

	// check moves
	result := check_moves(pawn_moves, expected)
	if !result {
		error_str := useful_error_msg(pawn_moves, expected)
		t.Errorf("Pawn Capture failed :\n%s", error_str)
	}
}

func TestEnpassCapture(t *testing.T) {

	var expected []uint

	fen := "rnbqkb1r/3p4/6n1/2p1Pp2/pP1NP3/PP5p/3B3P/R1BQK2R b KQkq - 0 1"
	gs := chess_engine.CreateGameFen(fen)
	
	// make double pawn move d7d5
	var move_num uint = 0b0001 << 12 | (3 + 32) << 6 | (3 + 48)
	if !contains(gs.MoveList, move_num) {
		t.Errorf("Move %b not found in move list", move_num)
	}
	gs.Make_move(move_num)


	// get all the pawn moves
	pawn_moves := get_piece_moves(gs.MoveList, *gs.Board.WhitePawns)

	// get normal pawn moves
	capture := create_moves([]string{"e4", "b3", "b4", "e4"}, 
		[][]string{{"f5"}, {"a4"}, {"c5"}, {"d5"}}, 0b0100)
	expected = append(expected, capture...)

	normal_moves := create_moves([]string{"e5", "b4"}, 
			[][]string{{"e6"}, {"b5"}}, 0)
	expected = append(expected, normal_moves...)

	// en - passant capture
	en_passant := create_moves([]string{"e5"}, [][]string{{"d6"}}, 0b0101)
	expected = append(expected, en_passant...)

	// check moves
	result := check_moves(pawn_moves, expected)
	if !result {
		error_str := useful_error_msg(pawn_moves, expected)
		t.Errorf("Enpassent test failed :\n%s", error_str)
	}
}


func TestPawnPromotion(t *testing.T) {

	var expected []uint

	fen := "4n3/1P1P3k/8/8/8/8/8/3K4 w - - 0 1"
	gs := chess_engine.CreateGameFen(fen)

	// promotion b7b8, d7d8, d7e8

	moves := promo_movenum([]string{"b7", "d7"}, [][]string{{"b8"}, {"d8"}}, false)
	expected = append(expected, moves...)

	cap_moves := promo_movenum([]string{"d7"}, [][]string{{"e8"}}, true)
	expected = append(expected, cap_moves...)

	// get all the pawn moves
	pawn_moves := get_piece_moves(gs.MoveList, *gs.Board.WhitePawns)

	// check moves
	result := check_moves(pawn_moves, expected)
	if !result {
		error_str := useful_error_msg(pawn_moves, expected)
		t.Errorf("Promotion test failed :\n%s", error_str)
	}
}


// helper function for creating promotion moves
func promo_movenum(start []string, moves [][]string, capture bool) []uint {

	var expected []uint
	var special [4]uint

	if capture {
		special = [4]uint{0b1100, 0b1101, 0b1110, 0b1111}
	} else {
		special = [4]uint{0b1000, 0b1001, 0b1010, 0b1011}
	}

	for _, spec := range special {
		expected = append(expected, create_moves(start, moves, spec)...)
	}
	

	return expected
}



