package magic

import (
	"chess/chess_engine/board"
	"math/rand"
)

type Magicsq struct {
	index uint
	magic board.Bitboard 
	shift int 
	diag bool
	
	// attack rays - must be calculated
	attack_rays []board.Bitboard
	
	occ_mask board.Bitboard // inner occupancy 
}

// ============================================================================

// Generation functions // ! NEEDS TO BE REDESIGNED
func Gen_all_magics(diag bool) [64]Magicsq {
	
	var magics [64]Magicsq
	var magics_num board.Bitboard
	var default_shift int

	for i := 0; i < 64; i++ {

		// create magic square
		magics[i] = create_magicsq(uint(i), diag)
		default_shift = 64 - len(magics[i].occ_mask.Index())

		// gen magic
		if magics[i].shift == default_shift {
			magics_num = gen_magic( &magics[i] )
			magics[i].magic = magics_num
		} 

		break

	}
	return magics
}


// generate magic square for a given square // ! NEEDS TO BE REDESIGNED
func gen_magic(msq *Magicsq) board.Bitboard {

	var magics_tried uint = 0
	var magics_found uint = 0

	var best_magic board.Bitboard

	// generate all occupancy bitboards
	all_occ := allOccupancy(msq.index, msq.diag)

	
	for (magics_found < 1) || 
		( (magics_tried < 10_000) && (magics_found < 5) ){ 

		// generate random magic
		magic_bb := board.Bitboard(rand.Uint64())
		
		// check if magic is valid

		if check_magic(magic_bb, msq, &all_occ) {

			magics_found++


			best_magic = magic_bb

			println("Magic found: ", magics_found, " Magic tried: ", magics_tried, "\n")
			magic_bb.Print()
		}

		magics_tried++
	}

	return best_magic

}




// ============================================================================
// Staging


func allOccupancy(ind uint, diag bool) []board.Bitboard {
	// all possible bit combinations of fullrays
	
	all_occ := []board.Bitboard{board.Bitboard(0)}
	
	full_rays := fullrays(ind, diag)
	full_rays &= innerOccupancy(ind)
	
	index := full_rays.Index()
	
	var bb board.Bitboard
	var bb_comb []board.Bitboard // not the best way to do this
	
	
	for _, ind  := range index {
		bb = board.Bitboard(0) | (1 << ind)
		
		// combine all the rays
		bb_comb = []board.Bitboard{bb, 0}
		
		all_occ = combineRays(all_occ, bb_comb)
		
	}
	
	return all_occ
}


