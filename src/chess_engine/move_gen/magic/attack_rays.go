package magic

import "chess/chess_engine/board"

// ! I THINK THIS FILE CAN BE DEPRECATED

// ============================================================================
// 		All attack rays

// gen all attack rays for all squares
func Gen_attack_rays(diag bool) [64][]board.Bitboard {

	var rays [64][]board.Bitboard

	for i := 0; i < 64; i++ {
		rays[i] = attack_rays(uint(i), diag)
	}
	return rays
}



// get all possible attack rays - for a given index (square)
func attack_rays(ind uint, diag bool) []board.Bitboard {

	all_rays := []board.Bitboard{board.Bitboard(0)}

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

	for i, dir_max := range max_dir {
		
		dir_rays := []board.Bitboard{}
		dir_bb := board.Bitboard(0)

		val := directions[i]

		for x := 1; x <= dir_max; x++ {
			dir_bb|= 1 << (int(ind) + val * x)
			
			dir_rays = append(dir_rays, dir_bb)
		}

		// combine all the rays
		all_rays = combineRays(all_rays, dir_rays)

	}
	return all_rays
}




