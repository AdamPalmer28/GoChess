package cli_engine

import "fmt"

func GetMoves(moves []uint) {

	for _, move := range moves {

		start_sq := move & 0b111111
		end_sq := (move >> 6) & 0x3f
		special := (move >> 12) & 0xf

		start := Index_to_move(int(start_sq))
		end := Index_to_move(int(end_sq))

		fmt.Println(start, end, " special: ", special)
	}

}

