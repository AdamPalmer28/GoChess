package cli_engine

import (
	"strconv"
)

func Move_to_index(cord string) int {
	// convert a chess cord to an index

	var ind int

	cord = cord[0:2]
	
	letter := cord[0]
	number := cord[1]

	ind = int(number-'1')*8 + int(letter-'a')

	return ind
}

func Index_to_move(ind int) string {

	var cord string

	letter := rune(ind%8) + 'a'
	rank := ind/8 + 1
	cord = string(letter) + strconv.Itoa(rank)

	return cord
}

func (cfg *Config) move_input(user_input string) bool {

	start := user_input[0:2]
	end := user_input[2:] // string + 1 for promotion
	end_sq := end[0:2] 

	s := Move_to_index(start)
	e := Move_to_index(end_sq)

	
	move_num :=  (e << 6) | s

	//fmt.Printf("Move: %b\n", move_num)

	cfg.gs.Make_move(uint(move_num))
	cfg.gs.Board.Print()

	return true
}

func MoveNum_readable(move_num uint) string {

	var move string

	start := move_num & 0x3f
	end := (move_num >> 6) & 0x3f
	//special := (move_num >> 12) & 0xf

	move = Index_to_move(int(start)) + Index_to_move(int(end))

	return move
}