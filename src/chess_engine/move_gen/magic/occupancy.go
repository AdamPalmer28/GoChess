package magic

import "chess/src/chess_engine/board"

// ============================================================================
// 		Occupancy

// calcs the inner occupancy of a square (both diagonals and straights)
func innerOccupancy(ind uint, diag bool) board.Bitboard {

	row := int(ind / 8)
	col := int(ind % 8)

	// directions of move
	inner_occ := Fullrays(ind, diag)

	if row != 0 {
		// remove the bottom row
		inner_occ &= ^board.Rank1
	} 
	if row != 7 {
		// remove the top row
		inner_occ &= ^board.Rank8
	}
	if col != 0 {
		// remove the left column
		inner_occ &= ^board.FileA
	} 
	if col != 7 {
		// remove the right column
		inner_occ &= ^board.FileH
	}

	return inner_occ
}

// full rays from a square
func Fullrays(ind uint, diag bool) board.Bitboard {

	row := int(ind / 8)
	col := int(ind % 8)

	// directions of move
	var directions [4]int
	var max_dir [4]int

	if diag {
		directions = [4]int{9, -9, 7, -7}
		max_dir = [4]int{min(7 - col, 7 - row), min(col, row),
					min(7 - row, col), min(row, 7 - col)}
	} else {
		directions = [4]int{1, -1, 8, -8}
		max_dir = [4]int{7 - col, col, 7 - row, row}
	}
	rays := board.Bitboard(0)

	for i, dir := range directions {
		
		for j := 1; j <= max_dir[i]; j++ {

			sq := int(ind) + dir*(j)
			rays |= 1 << sq

		}
	}

	return rays
}


// ============================================================================
// Helper function
// ============================================================================

func combineRays(rays []board.Bitboard, 
			new_rays []board.Bitboard) []board.Bitboard {

	if len(new_rays) == 0 {
		return rays
	}
	
	comb_rays := []board.Bitboard{}
	for _, ray := range rays {
		for _, new_ray := range new_rays {
			comb_rays = append(comb_rays, (ray | new_ray))
		}
	}

	return comb_rays
}