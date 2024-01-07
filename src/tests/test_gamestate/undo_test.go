package test_gamestate

import (
	"chess/src/chess_engine/board"
	"chess/src/chess_engine/gamestate"
	"chess/src/tests/test_move_gen"
	"testing"
)

/*
Moves and undo to test:
	- simple quiet move
	- simple capture
	- enpassant
	- promotion
	- castle
*/

func TestUndo(t *testing.T) {

	var move uint
	var error_msg string
	// ------------------------------------------------------------------------
	// 1. Standard move

	gs := gamestate.MakeGameState()
	gs.Init()

	cb := gs.Board.Copy()
	original_move_list := list_copy(gs.MoveList)

	// make a move
	move = gs.MoveList[0]
	gs.Make_move(move)

	// undo the move
	gs.Undo()

	// check if the board is the same
	if !cb.Identical(gs.Board) {
		t.Errorf("1. Board not the same after undo")
		cb.Print()
		gs.Board.Print()
	}
	if !same_lists(original_move_list, gs.MoveList) {
		error_msg = test_move_gen.MoveListErrorMsg(gs.MoveList, original_move_list)
		t.Errorf("1. MoveList not the same after undo \n%s", error_msg)
	}

	// ------------------------------------------------------------------------
	// 2. Capture move
	fen := "r3kb1r/ppppqppp/n1b1p3/4P1n1/2NP4/2B2NB1/PPPQ1PPP/R3K2R w KQkq - 0 1"

	gs = gamestate.FEN_to_gs(fen)
	gs.Init()
	cb = gs.Board.Copy()
	original_move_list = list_copy(gs.MoveList)

	moves := create_moves([]string{"d2g5", "e7g5"}, 0b0100)
	
	// make move
	gs.Make_move(moves[0]) // queen capture knight
	if gs.History.Cap_pieces[0] != 1 {
		t.Errorf("2.1 Capture piece not recorded %d (expected 1)", gs.History.Cap_pieces[0])
	}
	gs.Make_move(moves[1]) // queen capture queen
	if gs.History.Cap_pieces[1] != 4 {
		t.Errorf("2.2 Capture piece not recorded %d (expected 4)", gs.History.Cap_pieces[1])
	}

	// undo the move
	gs.Undo()
	if len(gs.History.Cap_pieces) != 1 {
		t.Errorf("2.3 Capture piece not removed from history %d (expected 1)", len(gs.History.Cap_pieces))
	}
	gs.Undo()
	if len(gs.History.Cap_pieces) != 0 {
		t.Errorf("2.4 Capture piece not removed from history %d (expected 0)", len(gs.History.Cap_pieces))
	}

	// check if the board is the same
	if !cb.Identical(gs.Board) {
		t.Errorf("2. Board not the same after undo")
		cb.Print()
		gs.Board.Print()
	}
	if !same_lists(original_move_list, gs.MoveList) {
		error_msg = test_move_gen.MoveListErrorMsg(gs.MoveList, original_move_list)

		t.Errorf("2. MoveList not the same after undo \n%s", error_msg)
	}
	

		// -----------------
		// Enpassant Capture
	fen = "r3kb1r/ppp1qppp/n1b1p3/3pP1n1/2NP4/2B2NBP/PPPQ1PP1/R3K2R w KQkq d6 0 1"

	gs = gamestate.FEN_to_gs(fen)
	gs.Init()
	cb = gs.Board.Copy()
	exp_enpass_ind := board.Move_to_index("d6")

	if gs.Enpass_ind != exp_enpass_ind {
		t.Errorf("2. Enpassant square did not initialize correctly - check FEN_to_gs")
	}

	moves = create_moves([]string{"e5d6"}, 0b0101)

	// make move
	gs.Make_move(moves[0]) // pawn capture pawn
	if gs.History.Cap_pieces[0] != 0 {
		t.Errorf("2.5 Capture piece not recorded %d (expected 0)", gs.History.Cap_pieces[0])
	}
	if gs.History.EnPassHist[0] != exp_enpass_ind {
		t.Errorf("2.6 Enpassant square not recorded %d (expected %d)", gs.History.EnPassHist[0], exp_enpass_ind)
	}


	// undo the move
	gs.Undo()

	if len(gs.History.EnPassHist) != 0 {
		t.Errorf("2.7 Enpassant square not removed from history %d (expected 0)", len(gs.History.EnPassHist))
	}
	if gs.Enpass_ind != exp_enpass_ind {
		t.Errorf("2.8 Enpassant square not restored %d (expected %d)", gs.Enpass_ind, exp_enpass_ind)
	}

	// check if the board is the same
	if !cb.Identical(gs.Board) {
		t.Errorf("2.EP. Board not the same after undo")
		cb.Print()
		gs.Board.Print()
	}

	// ------------------------------------------------------------------------
	// 3. Promotion
	fen = "6n1/2k2P2/8/8/8/4K3/8/8 w - - 0 1"

	gs = gamestate.FEN_to_gs(fen)
	gs.Init()
	cb = gs.Board.Copy()

	moves = create_moves([]string{"f7f8"}, 0b0000)
	cap_moves := create_moves([]string{"f7g8"}, 0b0100)

	test_promotion_undo(t, moves[0], gs, cb)
	test_promotion_undo(t, cap_moves[0], gs, cb)
	

	// ------------------------------------------------------------------------
	// 4. Castle
	fen = "r3k2r/ppppqppp/n1b1p1n1/4P3/1bNP2N1/2B3B1/PPPQ1PPP/R3K2R w KQkq - 0 1"

	gs = gamestate.FEN_to_gs(fen)
	gs.Init()
	cb = gs.Board.Copy()

	var w_queen_castle uint = 0b0011_000010_000100
	var b_king_castle uint = 0b0010_111110_111100

	w_castle_rights := gs.WhiteCastle
	b_castle_rights := gs.BlackCastle

	// white queen castle
	gs.Make_move(w_queen_castle)
	cb2 := gs.Board.Copy()

	// black king castle
	gs.Make_move(b_king_castle)

	// undo the black move
	gs.Undo()

	// check if the board is the same
	if !cb2.Identical(gs.Board) {
		t.Errorf("4. Board not the same after undo 1")
		cb.Print()
		gs.Board.Print()
	}

	// check castle rights 
	if (gs.BlackCastle != b_castle_rights) || (gs.WhiteCastle != 0) {
		t.Errorf("4. Castle rights error %b, %b (expected 0, 11)", gs.WhiteCastle, gs.BlackCastle)
	}

	// undo the white move
	gs.Undo()

	// check if the board is the same
	if !cb.Identical(gs.Board) {
		t.Errorf("4. Board not the same after undo 2")
		cb.Print()
		gs.Board.Print()
	}

	// check castle rights
	if (gs.BlackCastle != b_castle_rights) || (gs.WhiteCastle != w_castle_rights) {
		t.Errorf("4. Castle rights error %b, %b (expected 11, 11)", gs.WhiteCastle, gs.BlackCastle)
	}


	// ------------------------------------------------------------------------
	// 5. Gamestate data
	fen = "r3k2r/ppppqppp/n1b1p1n1/4P3/1bNP2N1/2B3B1/PPPQ1PPP/R3K2R w KQkq - 0 1"

	gs = gamestate.FEN_to_gs(fen)
	gs.Init()
	cb = gs.Board.Copy()

	move_hist := []uint{}
	// make a move
	moves = create_moves([]string{"a2a3"}, 0b0000)
	move_hist = append(move_hist, moves[0])
	gs.Make_move(moves[0])
	moves = create_moves([]string{"b3a4"}, 0b0000)
	move_hist = append(move_hist, moves[0])
	gs.Make_move(moves[0])
	moves = create_moves([]string{"b2b4"}, 0b0001)
	move_hist = append(move_hist, moves[0])
	gs.Make_move(moves[0])
	moves = create_moves([]string{"a4b3"}, 0b0100)
	gs.Make_move(moves[0])
	move_hist = append(move_hist, moves[0])

	// check postion is as expected
	if gs.Moveno != 5 || !gs.White_to_move {
		t.Errorf("5. Gamestate data not correct %d, %t (expected 5, true)", gs.Moveno, gs.White_to_move)
	}
	if !same_lists(gs.History.PrevMoves, move_hist) {
		error_msg = test_move_gen.MoveListErrorMsg(gs.History.PrevMoves, move_hist)
		t.Errorf("5. MoveList not the same after undo \n%s", error_msg)
	}

	// undo the move
	gs.Undo()

	// move number, half move number, white to move
	if gs.Moveno != 4 || gs.White_to_move {
		t.Errorf("5. Gamestate data not correct %d, %t (expected 4, false)", gs.Moveno, gs.White_to_move)
	}
	if !same_lists(gs.History.PrevMoves, move_hist[:3]) {
		error_msg = test_move_gen.MoveListErrorMsg(gs.History.PrevMoves, move_hist[:len(move_hist)-1])
		t.Errorf("5. MoveList not the same after undo \n%s", error_msg)
	}

	// ------------------------------------------------------------------------
	// 6. InCheck
	fen = "r3kb1r/ppppqppp/n1b1p3/4P3/2NP2N1/2B2nB1/PPPQ1PPP/R3K2R w KQkq - 0 1"

	gs = gamestate.FEN_to_gs(fen)
	gs.Init()
	original_move_list = list_copy(gs.MoveList)
	cb = gs.Board.Copy()

	moves = create_moves([]string{"g2f3"}, 0b0100)

	// make move
	gs.Make_move(moves[0]) // pawn capture knight

	// undo the move
	gs.Undo()

	// check if the board is the same
	if !cb.Identical(gs.Board) {
		t.Errorf("6. Board not the same after undo")
		cb.Print()
		gs.Board.Print()
	}
	if !gs.InCheck {
		t.Errorf("6. InCheck not restored after undo")
	}
	if !same_lists(original_move_list, gs.MoveList) {
		error_msg = test_move_gen.MoveListErrorMsg(gs.MoveList, original_move_list)
		t.Errorf("6. MoveList not the same after undo \n%s", error_msg)
	}
}


// ============================================================================
// helper function

func test_promotion_undo(t *testing.T, move uint, gs *gamestate.GameState, 
					original_cb board.ChessBoard) {

	specials := []uint{0b1000, 0b1001, 0b1010, 0b1011} // knight, bishop, rook, queen

	original_move_list := list_copy(gs.MoveList)
	for _, spec := range specials {

		special_move := move | (spec << 12)

		// make move
		gs.Make_move(special_move)

		// undo the move
		gs.Undo()
		
		// check if the board is the same
		if !original_cb.Identical(gs.Board) {
			t.Errorf("3. Board not the same after undo")
			original_cb.Print()
			gs.Board.Print()
		}
		if !same_lists(original_move_list, gs.MoveList) {
			error_msg := test_move_gen.MoveListErrorMsg(gs.MoveList, original_move_list)
			t.Errorf("3 MoveList not the same after promotion\n%s", error_msg) 
		}
	}
}

func create_moves(moves []string, special uint) []uint {

	var result []uint

	for _, move_str := range moves {
		start_sq := move_str[0:2]
		end_sq := move_str[2:4]

		start_sq_ind := board.Move_to_index(start_sq)
		end_sq_ind := board.Move_to_index(end_sq)

		move_num := (start_sq_ind | (end_sq_ind << 6) | (special << 12))

		result = append(result, move_num)
	}

	return result
}

// ----------------------------------------------------------------------------
// LIST HELPERS

func list_copy(list []uint) []uint {

	result := []uint{}

	result = append(result, list...)

	return result
}

// list checking - used for comparing more lists
func same_lists(list []uint, list2 []uint) bool {
	
	if len(list) != len(list2) {
		return false
	}
	for _, v := range list {
		if !contains(list2, v) {
			return false
		}
	}
	return true
}

func contains(list []uint, value uint) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

