package move_gen

import "chess/chess_engine/board"

// ============================================================================
// 		All attack rays

// get all possible rook attack rays - for a given index
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

func Gen_attack_rays(diag bool) [64][]board.Bitboard {

	var rays [64][]board.Bitboard

	for i := 0; i < 64; i++ {
		rays[i] = attack_rays(uint(i), diag)
	}
	return rays
}


// ============================================================================
// 		Occupancy 

func allOccupancy(ind uint, diag bool) []board.Bitboard {

	all_occ := []board.Bitboard{board.Bitboard(0)}

	row := int(ind / 8)
	col := int(ind % 8)

	maxrays := fullrays(ind, diag)

	// all possible bit combinations of fullrays


	return all_occ



func fullrays(ind uint, diag bool) board.Bitboard {

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

func 

