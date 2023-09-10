package test_gamestate

// test game state and

import (
	"chess/chess_engine/gamestate"
	"testing"
)

// test if state is: normal, check, checkmate, stalemate
func Test_Gamestate_state(t *testing.T) {

	// normal state
	normal_state_fen := []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"rnbqk2r/pp2bp1p/2p1pnp1/3p4/5P2/2P5/PP1PP1PP/RNBQKBNR w KQkq - 0 1",
		"r1bqk2r/pp2bp1p/2p1pnp1/2np4/3P1P2/N1P1PN2/PP4PP/R1BQKB1R w KQkq - 0 1",
		"r1b5/ppk2p1p/2pb1npr/1n1p4/3PpP1q/1PP1P1N1/PBN3PP/R2QKB1R b KQa - 0 1",
		"3k4/8/8/8/3K4/4p3/8/8 w - - 0 1",
	}

	// check state
	check_state_fen := []string{
		"r1bqk2r/pp2bp1p/2p1pnp1/3p4/3P1P2/N1P1PN2/PPn3PP/R1BQKB1R w KQkq - 0 1", // knight 
		"r1bqk2r/pp3p1p/2p1pnp1/bn1p4/2PP1P2/N3PN2/PP4PP/R1BQKB1R w KQkq - 0 1", // bishop
		"r1b1k2r/pp3p1p/2pbpnp1/1n1p4/3P1P1q/N1P1PN2/PP4PP/R1BQKB1R w KQkq - 0 1", // queen
		"r1b1k3/pp3p1p/2pbpnp1/1n1p4/3P1P1q/N1PKPN2/PP1r2PP/R1BQ1B1R w HAq - 0 1", // rook
		"r1b1k3/pp3p1p/2pb1npr/1n1p4/3PpP1q/N1PKPN2/PP4PP/R1BQ1B1R w HAq - 0 1", // pawn
		"r1b1k3/pp3pNp/2pb1npr/1n1p4/3PpP1q/N1P1P3/PP4PP/R1BQKB1R b KQq - 0 1", // b knight
		"r1b1k3/pp1P1p1p/2pb1npr/1n1p4/3PpP1q/N1P1P1N1/P5PP/R1BQKB1R b KQq - 0 1", // b pawn
		"1rb1R3/ppk2Q1p/2pb1npr/Pn1p4/3PpP1q/1PP1P1N1/1BN3PP/4KB1R b K - 0 1", // b queen
		"6rk/pp5p/b5p1/P7/1B3P2/1Ppn4/2NrBQPP/4KR2 w - - 0 1", // capture kn
	}


	// checkmate state
	checkmate_state_fen := []string{
		"1rb1R3/ppk1Q2p/2p1N1pr/Pn1p4/3PpP2/1PP1P3/1BN3PP/4KB1R b K - 0 1", 
		"5Q1k/pp1r3p/b4Np1/PB6/4pP2/BP2P3/2N3PP/4K2R b K - 0 1",
		"1rb1R3/ppk1Q2p/2p1N1pr/Pn1p4/3PpP2/1PP1P3/1BN3PP/4KB1R b K - 0 1", // double check
		"4k3/pp1rQ2p/b5pr/PB1N4/4pP2/1P2P3/1BN3PP/4K2R b K - 0 1", // pinned pieces
		"7k/pp5p/b5p1/PB6/1B3P2/1Pp1r3/2N1NQPP/1r2K2R w K - 0 1", // pinned pieces
		"7k/pp5p/b5p1/P7/1B3P2/1Ppnr3/2NrBQPP/4KR2 w - - 0 1", // pinned pieces
	}


	// stalemate state
	stalemate_state_fen := []string{
		"6k1/1p5p/pB3QpR/P5P1/1p6/1P6/2N5/6K1 b - - 0 1",
		"7k/1pR4p/pB4pN/P5P1/1p6/1P3Q2/8/6K1 b - - 0 1",
		"4R3/1p5p/pB4p1/P5P1/1p3Q2/1P5k/2N5/6K1 b - - 0 1", // king bubble
		"4Rbk1/1p5p/pB3QpB/P5P1/1p6/1P6/2N5/4K3 b - - 0 1", // pinned pieces
	}

	exp_InCheck := [4]bool{false, true, true, false}
	exp_GameOver := [4]bool{false, false, true, true}
	diff_states := [4][]string{normal_state_fen, check_state_fen, checkmate_state_fen, stalemate_state_fen}

	debug_states_text := [4]string{"Normal", "Check", "Checkmate", "Stalemate"}

	for ind, fens := range(diff_states) {

		for i, fen := range(fens) {
			state_str := debug_states_text[ind]

			gs := gamestate.FEN_to_gs(fen)
			gs.Init()

			if gs.InCheck != exp_InCheck[ind] {
				t.Errorf("InCheck error: fen index %v (expected state: %v)", i, state_str)
			}

			if gs.GameOver != exp_GameOver[ind] {
				t.Errorf("GameOver error: fen index %v (expected state: %v)", i, state_str)
			}
		}
	}

}

