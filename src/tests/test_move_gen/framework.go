package test_move_gen

import (
	"chess/chess_engine/board"
	"strconv"
)

func create_moves(start []string, moves [][]string, special uint) []uint {

	var result []uint



	for ind, start_sq := range start {
		sq_moves := moves[ind]

		for _, end_sq := range sq_moves {
			start_sq_ind := board.Move_to_index(start_sq)
			end_sq_ind := board.Move_to_index(end_sq)

			move_num := (start_sq_ind | (end_sq_ind << 6) | (special << 12))

			result = append(result, move_num)
		}
	}

	return result
}


// get all the related moves to a piece_bb
func get_piece_moves(move_list []uint, piece_bb board.Bitboard) []uint {

	var result []uint

	for _, move := range move_list {

		start_sq := move & 0b111111

		if (piece_bb >> uint(start_sq)) & 1 == 1 {
			result = append(result, move)
		}
	}
	return result
}

// check if is in expected list has all the moves in moves
func check_moves(moves []uint, expected []uint) bool {

	// check if moves is the same length as expected
	if len(moves) != len(expected) {
		return false
	}


	// create map of expected moves
	expected_map := make(map[uint]bool)
	for _, exp := range expected {
		expected_map[exp] = true
	}

	// check if moves isn't in expected
	for _, move := range moves {
		if !expected_map[move] {
			return false
		}
	}
	return true
}



func useful_error_msg(move_list []uint, expected []uint) string {
	
	var expected_diff []uint
	var move_diff []uint

	if len(move_list) != len(expected) {
		return "Different move lists sizes:\n" +
			"Expected: " + move_list_to_string(expected_diff) + "\n" +
			"Got: " + move_list_to_string(move_diff)
	}


	// find the moves that are in expected but not in move_list
	for _, exp := range expected {
		if !contains(move_list, exp) {
			expected_diff = append(expected_diff, exp)
		}
	}

	// find the moves that are in move_list but not in expected
	for _, move := range move_list {
		if !contains(expected, move) {
			move_diff = append(move_diff, move)
		}
	}

	return "Expected: " + move_list_to_string(expected_diff) + "\n" +
			"Got: " + move_list_to_string(move_diff)




}

// ----------------------------------------------------------------------------
// HELPER FUNCTIONS
// ----------------------------------------------------------------------------

func move_list_to_string(move_list []uint) string {

	var result string

	for _, move := range move_list {
		start_sq := move & 0b111111
		end_sq := (move >> 6) & 0b111111
		special := (move >> 12) & 0b111111

		// convert uint to string
		special_str := strconv.FormatUint(uint64(special), 2)
		

		result += board.Index_to_move(start_sq) + 
					board.Index_to_move(end_sq) + 
					" " + string(special_str) + ", "
	}

	return result
}



func contains(list []uint, value uint) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}