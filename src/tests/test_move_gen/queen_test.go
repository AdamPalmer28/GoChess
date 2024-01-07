package test_move_gen

import (
	"chess/src/chess_engine"
	"testing"
)

func TestQueen(t *testing.T) {

	var expected []uint

	fen := "rn1qkbnr/1pppp1pp/p2b1p2/7R/2B2B2/2P3Q1/PPR1P3/1N2K1N1 w kq - 0 1"
	gs := chess_engine.CreateGameFen(fen)

	eng_moves := &gs.MoveList

	// get all the queen moves
	moves := get_piece_moves(*eng_moves, *gs.Board.WhiteQueens)

	// expected moves
	non_caps := create_moves([]string{"g3"},
		[][]string{{"f3","e3","d3","h3","g2","g4","g5","g6","f2","h2","h4"}}, 0)
	expected = append(expected, non_caps...)

	caps := create_moves([]string{"g3"}, [][]string{{"g7"}}, 0b0100)
	expected = append(expected, caps...)

	// check moves
	result := check_moves(moves, expected)
	if !result {
		error_str := MoveListErrorMsg(moves, expected)
		t.Errorf("Queen moves failed:\n%s", error_str)
	}
}