package move_gen

import (
	"chess/chess_engine/board"
)

// horizontal and vertical rays
func slidingRays(ind uint, bb board.Bitboard)  board.Bitboard{

	// get row and column of index
	row := ind / 8
	col := ind % 8
	
	// directions of move
	directions := [4]int{1, -1, 8, -8}
	max_dir := [4]uint{7 - col, col, 7 - row, row}
	
	rays := board.Bitboard(0)

	for i, dir := range directions {

		for j := 1; uint(j) <= max_dir[i]; j++ {

			sq := int(ind) + dir * (j)
			rays |= 1 << sq

			// check if ray is blocked
			if bb & (1 << sq) != 0 {
				break
			}
		} 
	}		

	return rays
}


// diagonal rays
func diagonalRays(ind uint, bb board.Bitboard)  board.Bitboard{
	
	// get row and column of index
	row := ind / 8
	col := ind % 8
	
	// directions of move
	directions := [4]int{9, -9, 7, -7}
	max_dir := [4]uint{min(7 - col, 7 - row), min(col, row), 
				min(7 - row, col), min(row, 7 - col)}

	rays := board.Bitboard(0)

	for i, dir := range directions {

		for j := 1; uint(j) <= max_dir[i]; j++ {

			sq := (int(ind)) + dir * (j)
			rays |= 1 << sq

			// check if ray is blocked
			if bb & (1 << sq) != 0 {
				break
			}
		}
	}		

	return rays
}


func min(a, b uint) uint {
    if a < b {
        return a
    }
    return b
}

func max(a, b uint) uint {
    if a > b {
        return a
    }
    return b
}
