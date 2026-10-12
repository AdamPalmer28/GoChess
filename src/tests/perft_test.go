package test

import (
	"chess/src/chess_engine"
	"chess/src/chess_engine/board"
	"chess/src/chess_engine/gamestate"
	"fmt"
	"sort"
	"strings"
	"testing"
)

const (
	perftPosition3 = "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1"
	perftPosition5 = "rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8"
)

// Reference node counts come from ChessProgramming Wiki's standard perft set:
// https://www.chessprogramming.org/Perft_Results
func TestPerftReferencePositions(t *testing.T) {
	tests := []struct {
		name  string
		fen   string
		depth uint
		want  uint64
	}{
		{
			name:  "initial position",
			fen:   "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			depth: 4,
			want:  197281,
		},
		{
			name:  "kiwipete castling checks and pins",
			fen:   "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
			depth: 3,
			want:  97862,
		},
		{
			name:  "en passant and king safety",
			fen:   perftPosition3,
			depth: 4,
			want:  43238,
		},
		{
			name:  "promotion and castling rights",
			fen:   "r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1",
			depth: 3,
			want:  9467,
		},
		{
			name:  "promotion and check evasion",
			fen:   perftPosition5,
			depth: 3,
			want:  62379,
		},
		{
			name:  "pins and checks",
			fen:   "r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10",
			depth: 3,
			want:  89890,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := chess_engine.CreateGameFen(tt.fen)
			got := perft(gs, tt.depth)
			if got != tt.want {
				divide := perftDivide(gs, tt.depth)
				t.Fatalf("perft(%d) = %d, want %d\nroot divide:\n%s", tt.depth, got, tt.want, formatDivide(divide))
			}
		})
	}
}

func TestEnPassantThatExposesRookCheckIsIllegal(t *testing.T) {
	gs := chess_engine.CreateGameFen(perftPosition3)
	playTestMove(t, gs, "g2g4")

	if hasLegalMove(gs, "f4g3") {
		t.Fatal("en passant should be illegal when removing both pawns exposes the black king to the rook on b4")
	}
}

func TestNonPawnOnPawnAttackSquareDoesNotGiveCheck(t *testing.T) {
	gs := chess_engine.CreateGameFen(perftPosition5)
	if gs.InCheck {
		t.Fatal("the black knight on f2 does not attack the white king on e1")
	}
}

func TestCheckEvasionDoublePushRequiresClearIntermediateSquare(t *testing.T) {
	gs := chess_engine.CreateGameFen(perftPosition5)
	playTestMove(t, gs, "c4b3")
	playTestMove(t, gs, "d8a5")

	if !gs.InCheck {
		t.Fatal("queen move to a5 should check the white king on e1")
	}
	if hasLegalMove(gs, "b2b4") {
		t.Fatal("b2-b4 cannot block the check because the white bishop occupies b3")
	}
}

func TestUndoRestoresOpponentCastlingRightsAfterRookCapture(t *testing.T) {
	gs := chess_engine.CreateGameFen("4k2r/8/8/8/8/8/8/4K2R w k - 0 1")
	playTestMove(t, gs, "h1h8")
	if gs.BlackCastle != 0 {
		t.Fatalf("capturing h8 rook should remove black's kingside castling right, got %b", gs.BlackCastle)
	}

	gs.Undo()
	if gs.BlackCastle != 1 {
		t.Fatalf("undo should restore black's kingside castling right, got %b", gs.BlackCastle)
	}
}

func playTestMove(t *testing.T, gs *gamestate.GameState, name string) {
	t.Helper()
	for _, move := range gs.MoveList {
		if perftMoveName(move) == name {
			gs.Make_move(move)
			return
		}
	}
	t.Fatalf("legal move %s not found", name)
}

func hasLegalMove(gs *gamestate.GameState, name string) bool {
	for _, move := range gs.MoveList {
		if perftMoveName(move) == name {
			return true
		}
	}
	return false
}

// perft counts the legal positions at exactly depth plies below gs. A leaf is
// counted without applying its moves; all earlier moves are made and undone.
func perft(gs *gamestate.GameState, depth uint) uint64 {
	if depth == 0 {
		return 1
	}
	if depth == 1 {
		return uint64(len(gs.MoveList))
	}

	// Make_move refreshes MoveList, so preserve the current node's moves first.
	moves := append([]uint(nil), gs.MoveList...)
	var nodes uint64
	for _, move := range moves {
		gs.Make_move(move)
		nodes += perft(gs, depth-1)
		gs.Undo()
	}
	return nodes
}

type perftDivideEntry struct {
	move  string
	nodes uint64
}

// perftDivide returns the leaf count beneath every legal root move, like the
// divide output used to locate a bad branch in a perft mismatch.
func perftDivide(gs *gamestate.GameState, depth uint) []perftDivideEntry {
	if depth == 0 {
		return nil
	}

	moves := append([]uint(nil), gs.MoveList...)
	divide := make([]perftDivideEntry, 0, len(moves))
	for _, move := range moves {
		gs.Make_move(move)
		nodes := perft(gs, depth-1)
		gs.Undo()
		divide = append(divide, perftDivideEntry{move: perftMoveName(move), nodes: nodes})
	}

	sort.Slice(divide, func(i, j int) bool {
		return divide[i].move < divide[j].move
	})
	return divide
}

func perftMoveName(move uint) string {
	start := board.Index_to_move(move & 0x3f)
	end := board.Index_to_move((move >> 6) & 0x3f)
	name := start + end

	special := (move >> 12) & 0xf
	if special&0x8 != 0 {
		promotion := [...]byte{'n', 'b', 'r', 'q'}
		name += string(promotion[special&0x3])
	}
	return name
}

func formatDivide(divide []perftDivideEntry) string {
	var output strings.Builder
	for _, entry := range divide {
		fmt.Fprintf(&output, "  %s: %d\n", entry.move, entry.nodes)
	}
	return output.String()
}
