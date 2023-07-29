package cli_debug

import (
	"chess/chess_engine/board"
	"chess/chess_engine/gamestate"
)

func Bitboard_cli(gs gamestate.GameState, inputs []string) {
	// print bitboard of current gamestate
	
	pieceBitboard := map[string]*board.Bitboard{
		"wp": gs.Board.WhitePawns,
		"bp": gs.Board.BlackPawns,
		"wkn": gs.Board.WhiteKnights,
		"bkn": gs.Board.BlackKnights,
		"wb": gs.Board.WhiteBishops,
		"bb": gs.Board.BlackBishops,
		"wr": gs.Board.WhiteRooks,
		"br": gs.Board.BlackRooks,
		"wq": gs.Board.WhiteQueens,
		"bq": gs.Board.BlackQueens,
		"wk": gs.Board.WhiteKing,
		"bk": gs.Board.BlackKing,

		"w": &gs.Board.White,
		"b": &gs.Board.Black,
	}

	bb_type := inputs[1]

	// check value in map
	bb := pieceBitboard[bb_type]
	if bb == nil {
		println("Invalid bitboard type ", bb_type)
		return
	}

	// print bitboard
	bb.Print()
	
}