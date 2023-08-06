package cli_engine

import (
	"chess/chess_engine/board"
	"fmt"
)

func GetMoves(moves []uint) {

	for _, move := range moves {

		start_sq := move & 0b111111
		end_sq := (move >> 6) & 0x3f
		special := (move >> 12) & 0xf

		start := board.Index_to_move(start_sq)
		end := board.Index_to_move(end_sq)

		fmt.Println(start+end, " special: ", special)
	}

}

