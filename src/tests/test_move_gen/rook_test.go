package test_move_gen

import (
	"chess/chess_engine"
	"testing"
)

func TestRook(t *testing.T) {

	var expected []uint

	fen := "rn1qkbnr/1pppp1pp/p2b1p2/7R/2B2B2/2P3Q1/PPR1P3/1N2K1N1 w kq - 0 1"
	gs := chess_engine.CreateGameFen(fen)

	eng_moves := &gs.MoveList

	// get all the queen moves
	moves := get_piece_moves(*eng_moves, *gs.Board.WhiteRooks)

	// expected moves
	non_caps := create_moves([]string{"h5", "c2"},
		[][]string{{"g5","f5","e5","d5","c5","b5","a5","h1","h2","h3","h4","h6",}, {"c1","d2"}}, 0)
	expected = append(expected, non_caps...)

	caps := create_moves([]string{"h5"}, [][]string{{"h7"}}, 0b0100)
	expected = append(expected, caps...)

	// check moves
	result := check_moves(moves, expected)
	if !result {
		error_str := MoveListErrorMsg(moves, expected)
		t.Errorf("Rook moves failed :\n%s", error_str)
	}
}