package benchmark

import (
	"chess/src/chess_engine"
	"chess/src/chess_engine/move_gen"
	"chess/src/chess_engine/move_gen/magic"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var Fen_positions_short = [][2]string{
	{"Starting pos", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"},
	{"End of opening", "r3k2r/pppb1ppp/1qnbpn2/3P4/3P1N2/2N3P1/PP1BPPBP/R2QK2R w KQkq - 0 1"},
	{"Middle 2", "2rq1rk1/1p1b1pp1/p4n1p/2bp4/3N3B/2N1P3/PPQ2PPP/R4RK1 w - - 0 16"},
	{"InCheck 1", "r2bnk1r/ppp3pp/1qn1p1p1/8/1B1P4/P1N3P1/1P2PPBP/R2Q1RK1 b - - 0 1"},
	{"InCheck 3", "1b4rk/1p2qp1p/p1n5/2p1p2P/P1NpP3/1P1P1N2/2PQ1Kr1/R5R1 w - - 0 29"},
	{"End 2", "2k5/4rp1p/p1n5/Ppp1p1NP/3pP3/1P1P4/2PK1R2/8 w - b6 0 29"},
	{"End 4", "2b5/3PnP1k/5ppP/KP6/P2p2r1/2p3P1/2Rb4/8 b - - 0 44"},
}

func Benchmark_Next_move(b *testing.B) {
	// Next_move function within gamestate
	//	- this to prepare gamestate for next move (including move generation)
	b.ReportAllocs()

	for pos_ind, fen := range Fen_positions_short {

		// Setup gamestate

		str_ind := strconv.Itoa(pos_ind)
		name, fen := fen[0], fen[1]
		gs := chess_engine.CreateGameFen(fen)

		fmt.Println("\n\nPostion: ", name, " - Incheck:", gs.InCheck)
		fmt.Println("Fen: ", fen)

		// --------------------------------------------------------------------

		// Make_BP function
		b.Run(str_ind+"__Make_BP__", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gs.Make_BP()
			}
		})
		// GetCheck function
		b.Run(str_ind+"__GetCheck__", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gs.GetCheck()
			}
		})		
		pinned_pieces := gs.GetCheck() // get check status
		// Benchmark remove_illegal_moves function
		b.Run(str_ind+"__RM_IllegalMoves__", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gs.RM_IllegalMoves(pinned_pieces)
			}
		})
		// MoveGen function
		var moveGenFn func()
		if gs.InCheck {
			moveGenFn = gs.GenCheckMoves
			} else {
				moveGenFn = gs.GenMoves
			}
			b.Run(str_ind+"__MoveGen_Total__", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					moveGenFn()
				}
			})
			

		
		
		// MOVE GENERATION
		// --------------------------------------------------------------------
		move_gen_prefix := "__MoveGen_"
		// check if InCheck is in name
		if strings.Contains(name, "InCheck") {
			continue
		}
		var player move_gen.BoardPerpective = gs.PlayerBoard

		// pawn moves
		b.Run(str_ind + move_gen_prefix + "_PawnMoves__", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var pawn_caps_ind uint
				if gs.White_to_move {pawn_caps_ind = 0} else {pawn_caps_ind = 1}
				move_gen.GenPawnMoves(player.Pawn_bb, gs.White_to_move, gs.Enpass_ind, 
										&gs.MoveRays.PawnCapRays[pawn_caps_ind],
										player.Team_bb, player.Opp_bb)
			}
		})
		// knight moves
		b.Run(str_ind + move_gen_prefix + "_KnightMoves__", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				move_gen.GenKnightMoves(player.Knight_bb, 
										&gs.MoveRays.KnightRays, player.Team_bb, player.Opp_bb)
			}
		})
		// rook moves
		b.Run(str_ind + move_gen_prefix + "_RookMoves__", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				magic.GenMagicMoves((player.Rook_bb | player.Queen_bb), &gs.MoveRays.Magic.RookMagic,
									player.Team_bb, player.Opp_bb)
			}
		})
		// bishop moves
		b.Run(str_ind + move_gen_prefix + "_BishopMoves__", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				magic.GenMagicMoves((player.Bishop_bb | player.Queen_bb), &gs.MoveRays.Magic.BishopMagic,
									player.Team_bb, player.Opp_bb)
			}
		})
		// // queen moves
		// b.Run(str_ind + move_gen_prefix + "__QueenMoves__", func(b *testing.B) {
		// 	for i := 0; i < b.N; i++ {
		// 		magic.GenMagicMoves(player.Queen_bb, &gs.MoveRays.Magic.BishopMagic,
		// 							player.Team_bb, player.Opp_bb)
		// 		magic.GenMagicMoves(player.Queen_bb, &gs.MoveRays.Magic.RookMagic,
		// 							player.Team_bb, player.Opp_bb)
		// 	}
		// })

		// king moves
		b.Run(str_ind + move_gen_prefix + "_KingMoves__", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				move_gen.GenKingMoves(gs.PlayerKingSaftey, player.Castle_rights,
									&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.Magic.BishopMagic,
									&gs.MoveRays.KnightRays, &gs.MoveRays.KingRays)
			}
		})


	}
}



