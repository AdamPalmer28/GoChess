package cli_engine

import (
	"strconv"
)

func move_to_index(cord string) int {
	// convert a chess cord to an index

	var ind int

	cord = cord[0:2]
	
	letter := cord[0]
	number := cord[1]

	ind = int(number-'1')*8 + int(letter-'a')

	return ind
}

func index_to_move(ind int) string {

	var cord string

	letter := rune(ind%8) + 'a'
	rank := ind/8 + 1
	cord = string(letter) + strconv.Itoa(rank)

	return cord
}