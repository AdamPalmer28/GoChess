package test_move_gen

import (
	"chess/src/chess_engine"
	"testing"
)

func Test_PinnedPieces_Moves(t *testing.T) {

	var actual_moves []uint
	var expected []uint
	var moves []uint
	var result bool


	fen := "3kr3/1b5b/5pP1/3PP3/q1Q1KPRr/4N3/2r1r1P1/7b w - - 0 1"
	gs := chess_engine.CreateGameFen(fen)

	//gs.Board.Print()

	// 1. Pawn moves ------------------------------------------------
	actual_moves = get_piece_moves(gs.MoveList, *gs.Board.WhitePawns)
	
	// expected moves
	expected = []uint{}
	moves = create_moves([]string{"d5","e5","f4","g6","g2"}, 
	[][]string{{},{"e6"},{"f5"}, {},{}}, 0)
	expected = append(expected, moves...)
	// capture
	moves = create_moves([]string{"g6"}, [][]string{{"h7"}}, 0b0100)
	expected = append(expected, moves...)
	
	result = check_moves(actual_moves, expected)
	if !result {
		error_str := MoveListErrorMsg(actual_moves, expected)
		t.Errorf("1. Pawn moves failed :\n%s", error_str)
	}
	
	// 2. Queen moves ------------------------------------------------
	actual_moves = get_piece_moves(gs.MoveList, *gs.Board.WhiteQueens)

	// expected moves
	expected = []uint{}
	moves = create_moves([]string{"c4"}, [][]string{{"b4","d4"}}, 0)
	expected = append(expected, moves...)
	moves = create_moves([]string{"c4"}, [][]string{{"a4"}}, 0b0100) // capture
	expected = append(expected, moves...)

	result = check_moves(actual_moves, expected)
	if !result {
		error_str := MoveListErrorMsg(actual_moves, expected)
		t.Errorf("2. Queen moves failed :\n%s", error_str)
	}

	// 3. Knight moves ------------------------------------------------
	actual_moves = get_piece_moves(gs.MoveList, *gs.Board.WhiteKnights)

	// expected moves
	expected = []uint{}
	moves = create_moves([]string{"e3"}, [][]string{{}}, 0)
	expected = append(expected, moves...)

	result = check_moves(actual_moves, expected)
	if !result {
		error_str := MoveListErrorMsg(actual_moves, expected)
		t.Errorf("3. Knight moves failed :\n%s", error_str)
	}

}