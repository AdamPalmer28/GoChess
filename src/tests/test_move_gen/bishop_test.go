package test_move_gen

import (
	"chess/chess_engine"
	"testing"
)

func TestBishop(t *testing.T) {

	var expected []uint

	fen := "rn1qkbnr/1pppp1pp/p2b1p2/7R/2B2B2/2P3Q1/PPR1P3/1N2K1N1 w kq - 0 1"
	gs := chess_engine.CreateGameFen(fen)

	eng_moves := &gs.MoveList

	// get all the queen moves
	moves := get_piece_moves(*eng_moves, *gs.Board.WhiteBishops)

	// expected moves
	non_caps := create_moves([]string{"c4", "f4"},
		[][]string{{"b3","d3","b5","d5","e6","f7"}, {"e3","d2","c1","e5","g5","h6"}}, 0)
	expected = append(expected, non_caps...)

	caps := create_moves([]string{"c4", "f4"}, [][]string{{"g8","a6"},{"d6"}}, 0b0100)
	expected = append(expected, caps...)

	// check moves
	result := check_moves(moves, expected)
	if !result {
		error_str := useful_error_msg(moves, expected)
		t.Errorf("Queen moves failed :\n%s", error_str)
	}
}