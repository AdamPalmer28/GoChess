package test_move_gen

import (
	"chess/chess_engine"
	"testing"
)

func TestKnightMoves(t *testing.T) {

	var expected []uint

	fen := "rnbqkbnr/ppp2pp1/3p4/3Bp2p/1NN4Q/5N2/PPPPPPPP/R1B1K2R w KQkq - 0 1"
	gs := chess_engine.CreateGameFen(fen)

	eng_moves := &gs.MoveList

	// get all the knight moves
	knight_moves := get_piece_moves(*eng_moves, *gs.Board.WhiteKnights)

	// expected moves
	noncaps := create_moves([]string{"b4", "c4", "f3"}, 
			[][]string{{"d3","a6","c6"}, {"a3","a5","b6","e3"}, {"g1","g5","d4"}}, 0)
	expected = append(expected, noncaps...)

	cap := create_moves([]string{"c4", "f3"},
			[][]string{{"d6", "e5"}, {"e5"}}, 0b0100)
	expected = append(expected, cap...)

	// check moves
	result := check_moves(knight_moves, expected)
	if !result {
		error_str := MoveListErrorMsg(knight_moves, expected)
		t.Errorf("Standard_moves failed :\n%s", error_str)
	}
}