package cli_engine

import (
	"fmt"
	"strconv"
)

func Move_to_index(cord string) uint {
	// convert a chess cord to an index

	var ind uint

	cord = cord[0:2]
	
	letter := cord[0]
	number := cord[1]

	ind = uint(number-'1')*8 + uint(letter-'a')

	return ind
}

func Index_to_move(ind uint) string {

	var cord string

	letter := rune(ind%8) + 'a'
	rank := ind/8 + 1
	cord = string(letter) + strconv.Itoa( int(rank) )

	return cord
}


// CLI move input - determin if move is valid and determine move number
func (cfg *Config) move_input(user_input string) bool {

	// check if the input is valid
	if len(user_input) > 5 || len(user_input) < 4 {
		fmt.Println("Invalid input length")
		return false
	}

	start := user_input[0:2]

	end := user_input[2:] // string + 1 for promotion
	end_sq := end[0:2] 
	
	s := Move_to_index(start)
	e := Move_to_index(end_sq)
	
	var move_num uint
	move_num =  (e << 6) | s
	
	var matched_moves []uint 
	// check if the move is valid
	for _, move := range cfg.gs.MoveList {
		// check 
		mv_sq := move & 0b111111_111111
		if mv_sq == move_num {
			matched_moves = append(matched_moves, move)
		}
	}

	if len(matched_moves) == 0 {
		fmt.Println("\nInvalid move - available moves: ")
		GetMoves(cfg.gs.MoveList)
		return false

	} else if len(matched_moves) > 1 {
		// promotion handling
		if len(user_input) != 5 {
			fmt.Println("Move should be promotion - 5 characters")

		} else {
			var special uint
			promotion := end[2]

			if promotion == 'q' {
				special = 0b1011
			} else if promotion == 'r' {
				special = 0b1010
			} else if promotion == 'b' {
				special = 0b1001
			} else if promotion == 'n' {
				special = 0b1000
			} else {
				fmt.Println("Invalid promotion")
				return false
			}

			piece := special & 0b0011
			for _, move := range matched_moves {
				promo_piece := (move >> 12) & 0b0011

				if promo_piece  == piece {
					move_num = move
					break
				}
			}

		}
	} else {
		// only one move matched
		move_num = matched_moves[0]
	}


	cfg.gs.Make_move(move_num)
	cfg.gs.Board.Print()

	return true
}

