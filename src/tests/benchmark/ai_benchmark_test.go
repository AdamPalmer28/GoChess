package benchmark

// Benchmarks the time to generate moves from a set of fen positions

import (
	"chess/src/chess_bot"
	"chess/src/chess_engine"
	"fmt"
	"strconv"
	"testing"
)

var Fen_positions = [][2]string{
	{"Starting pos","rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"},
	{"End of opening","r3k2r/pppb1ppp/1qnbpn2/3P4/3P1N2/2N3P1/PP1BPPBP/R2QK2R w KQkq - 0 1"},
	{"Middle 1", "r3k2r/pppb2pp/1qn1pnp1/8/1B1P4/P1N3P1/1P2PPBP/R2Q1RK1 b kq - 0 1"},
	{"Middle 2", "2rq1rk1/1p1b1pp1/p4n1p/2bp4/3N3B/2N1P3/PPQ2PPP/R4RK1 w - - 0 16"},
	{"Middle 3", "3r1rk1/p4q1n/1pp1b1pp/5p2/N1P1p3/PP2P2P/2Q1BPPB/R5K1 b - - 0 25"},
	{"InCheck 1", "r2bnk1r/ppp3pp/1qn1p1p1/8/1B1P4/P1N3P1/1P2PPBP/R2Q1RK1 b - - 0 1"},
	{"InCheck 2", "6k1/ppp1q1b1/2n4p/4pQp1/2P5/P3PNPP/1P2B1P1/3r2K1 w - - 0 22"},
	{"InCheck 3", "1b4rk/1p2qp1p/p1n5/2p1p2P/P1NpP3/1P1P1N2/2PQ1Kr1/R5R1 w - - 0 29"},
	{"End 1", "6rk/1p3p1p/p1n5/2p1p2P/P2pP3/1P1P1N2/2P2KR1/R7 w - - 0 29"},
	{"End 2", "2k5/4rp1p/p1n5/Ppp1p1NP/3pP3/1P1P4/2PK1R2/8 w - b6 0 29"},
	{"End 3", "8/7k/5ppP/K7/P5r1/2p5/2Rb4/8 w - - 0 44"},
	{"End 4", "2b5/3PnP1k/5ppP/KP6/P2p2r1/2p3P1/2Rb4/8 b - - 0 44 w - - 0 1"},
}



func Benchmark_Search(b *testing.B) {
	// Benchmark the AI Search process
	//	- this function is used to Evaluate the gamestate for search (bot move selection) 
	depth := []uint{1,  4, 6}
	b.ReportAllocs()

	for pos_ind, fen := range Fen_positions {
		str_ind := strconv.Itoa(pos_ind) 

		name, fen := fen[0], fen[1]

		fmt.Println("\nPostion: ", name)

		gs := chess_engine.CreateGameFen(fen)

		// Benchmark the Evaluate function
		b.Run(str_ind + "__Evaluate__", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				chess_bot.Evaluate(gs)
			}
		})

		// Benchmark the Next_move function
		b.Run(str_ind + "__Next_move__", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gs.Next_move()
			}
		})

		// Benchmark the FindBestMove function
		for _, d := range depth {

			b.Run(str_ind + "__FindBestMove__" + "-Depth-" + strconv.Itoa(int(d)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					chess_bot.FindBestMove(gs, d, false)
				}
			})
		}
		   
			
	}
}

