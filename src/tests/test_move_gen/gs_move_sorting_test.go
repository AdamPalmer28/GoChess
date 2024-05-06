package test_move_gen

// import (
// 	"chess/src/chess_engine/gamestate"
// 	"testing"
// )

// func Test_MoveSort(t *testing.T) {

// 	fen := "2rq1rk1/1p1bPpp1/p4n1p/2bp1N2/1P5B/8/2QP2P1/R6K w - - 0 16"

// 	// Setup gamestate
// 	gs := gamestate.FEN_to_gs(fen)
// 	gs.Init()

// 	ml := gs.MoveList
// 	//sml := gs.ScoreMoveList

// 	// sort the move list

// 	exp_moveorder := []uint{
// 		0b0000_000000_000000, // e7e8q
// 	}

// 	for i, move := range ml {
// 		exp := exp_moveorder[i]
// 		if move != exp {

// 			error_str := MoveListErrorMsg(ml, exp_moveorder)
// 			t.Errorf("Move order failed: %s", error_str)
// 		}
// 	}

// }