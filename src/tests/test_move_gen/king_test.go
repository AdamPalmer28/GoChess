package test_move_gen

import (
	"chess/src/chess_engine"
	"testing"
)

/*
Tests must include:
	- Basic moves
	- Opp king interference
	- King safety:
		- rook
		- bishop
		- queen
		- knight
		- pawn
*/

func TestBasicKing(t *testing.T) {
	// Test basic moves
	var actual_king_moves []uint
	var expected []uint
	var moves []uint
	var result bool
	// ------------------------------------------------------------------------
	// 1. No obstruction
	fen := "4q3/4rb2/4p1k1/8/8/5K2/2P3p1/2Q5 w - - 0 1"
	gs := chess_engine.CreateGameFen(fen)

	// get all the king moves
	actual_king_moves = get_piece_moves(gs.MoveList, *gs.Board.WhiteKing)

	// expected moves
	expected = []uint{}
	moves = create_moves([]string{"f3"}, [][]string{{"e2","e3","e4","f2","f4","g4","g3"}}, 0)
	expected = append(expected, moves...)
	moves = create_moves([]string{"f3"}, [][]string{{"g2"}}, 0b0100) // capture
	expected = append(expected, moves...)

	result = check_moves(actual_king_moves, expected)
	if !result {
		error_str := MoveListErrorMsg(actual_king_moves, expected)
		t.Errorf("1. King moves failed :\n%s", error_str)
	}
	

	// ------------------------------------------------------------------------
	// 2. Queen, pawn, knight obstruction
	fen = "4k3/8/1q6/8/4p3/8/3K2n1/2Q5 w - - 0 1"
	gs = chess_engine.CreateGameFen(fen)

	// get all the king moves
	actual_king_moves = get_piece_moves(gs.MoveList, *gs.Board.WhiteKing)

	// expected moves
	expected = []uint{}
	moves = create_moves([]string{"d2"}, [][]string{{"e2","d1","c2","c3"}}, 0)
	expected = append(expected, moves...)

	result = check_moves(actual_king_moves, expected)
	if !result {
		error_str := MoveListErrorMsg(actual_king_moves, expected)
		t.Errorf("2. King moves failed :\n%s", error_str)
	}

	// ------------------------------------------------------------------------
	// 3. knight, bishop, rook obstruction
	fen = "1q6/2r2b2/8/6k1/3Kp3/8/2P3n1/2Q5 w - - 0 1"
	gs = chess_engine.CreateGameFen(fen)


	// get all the king moves
	actual_king_moves = get_piece_moves(gs.MoveList, *gs.Board.WhiteKing)

	// expected moves
	expected = []uint{}
	moves = create_moves([]string{"d4"}, [][]string{{"e5"}}, 0)
	expected = append(expected, moves...)
	moves = create_moves([]string{"d4"}, [][]string{{"e4"}}, 0b0100) // capture
	expected = append(expected, moves...)

	result = check_moves(actual_king_moves, expected)
	if !result {
		error_str := MoveListErrorMsg(actual_king_moves, expected)
		t.Errorf("3. King moves failed :\n%s", error_str)
	}

	// ------------------------------------------------------------------------
	// 4. knight, bishop, rook, opp_king obstruction
	fen = "1q6/2r2b2/8/5k2/3Kp3/8/2P3n1/2Q5 w - - 0 1"
	gs = chess_engine.CreateGameFen(fen)

	// get all the king moves
	actual_king_moves = get_piece_moves(gs.MoveList, *gs.Board.WhiteKing)

	// expected moves
	expected = []uint{}
	// no moves

	result = check_moves(actual_king_moves, expected)
	if !result {
		error_str := MoveListErrorMsg(actual_king_moves, expected)
		t.Errorf("4. King moves failed :\n%s", error_str)
	}


	// ------------------------------------------------------------------------
	// 5. in check moves w/o protected piece
	fen = "5r2/5b2/6k1/8/4q3/3K4/1P4n1/Q7 w - - 0 1"
	gs = chess_engine.CreateGameFen(fen)

	// get all the king moves
	actual_king_moves = get_piece_moves(gs.MoveList, *gs.Board.WhiteKing)

	// expected moves
	expected = []uint{}
	moves = create_moves([]string{"d3"}, [][]string{{"d2","c3"}}, 0)
	expected = append(expected, moves...)
	moves = create_moves([]string{"d3"}, [][]string{{"e4"}}, 0b0100) // capture
	expected = append(expected, moves...)

	result = check_moves(actual_king_moves, expected)
	if !result {
		error_str := MoveListErrorMsg(actual_king_moves, expected)
		t.Errorf("5. King moves failed :\n%s", error_str)
	}


	// ------------------------------------------------------------------------
	// 6. in check moves w/ protected piece
	fen = "8/4rb2/6k1/8/4q3/3K4/1P4n1/Q7 w - - 0 1"
	gs = chess_engine.CreateGameFen(fen)

	// get all the king moves
	actual_king_moves = get_piece_moves(gs.MoveList, *gs.Board.WhiteKing)

	// expected moves
	expected = []uint{}
	moves = create_moves([]string{"d3"}, [][]string{{"d2","c3"}}, 0)
	expected = append(expected, moves...)

	result = check_moves(actual_king_moves, expected)
	if !result {
		error_str := MoveListErrorMsg(actual_king_moves, expected)
		t.Errorf("6. King moves failed :\n%s", error_str)
	}
}


const (
	// movenums of castle moves
	wQueenCastle = 0b0011_000010_000100
	wKingCastle = 0b0010_000110_000100
	bQueenCastle = 0b0011_111010_111100
	bKingCastle = 0b0010_111110_111100
)
func TestCastleGen(t *testing.T) {
	// Test castle moves

	castle_fen_map := map[string][]uint{
		"r3k2r/8/8/8/8/8/8/4K2R w K - 0 1": {wKingCastle},
		"r3k2r/8/8/8/8/8/8/R3K3 w Q - 0 1": {wQueenCastle},
		"r3k2r/8/8/8/8/8/8/R3K2R w KQ - 0 1": {wQueenCastle, wKingCastle},
		"r3k2r/8/8/8/8/8/8/R3K2R b KQkq - 0 1": {bQueenCastle, bKingCastle},
		"r3k2r/8/8/8/8/8/5P2/1Q2KQ1R b kq - 0 1": {bQueenCastle, bKingCastle},
		"r3k2r/8/8/8/8/8/5P2/1Q2K1QR b kq - 0 1": {bQueenCastle},
		"r3k2r/8/8/8/1Q6/8/5P2/4K1QR b kq - 0 1": {bQueenCastle},
		"r3k2r/8/6n1/1q3b2/1Q6/6Q1/5P2/R3K2R b KQkq - 0 1": {bQueenCastle},
		"r3k2r/8/8/8/1Qq5/6Q1/5P2/4K2R b Kkq - 0 1": {},
		"r3k2r/8/8/1q6/1Q4b1/6Q1/5P2/R3K2R b KQkq - 0 1": {},
		"r3k2r/8/8/1q6/1Q6/4n1Q1/5P2/R3K2R b KQkq - 0 1": {},
		"r3k2r/8/8/1q6/1Q6/6Q1/2n2P2/R3K2R b KQkq - 0 1": {}, // in check
	}

	ind := 0
	for fen, expected := range(castle_fen_map) {
		gs := chess_engine.CreateGameFen(fen)

		if !contains_list(gs.MoveList, expected) {

			t.Errorf("Castle moves not found:\n%s\n", fen)
		}
		ind++
	}
}


// helper
func contains_list(list []uint, sublist []uint) bool {

	for _, item := range(sublist) {
		if !contains(list, item) {
			return false
		}
	}
	return true
}

