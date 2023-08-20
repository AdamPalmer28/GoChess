package magic

import (
	"chess/chess_engine/board"
	"fmt"
	"math/rand"
)

type Magicsq struct {
	index uint
	magic board.Bitboard 
	shift int 
	diag bool
	
	// attack rays - must be calculated
	attack_rays map[board.Bitboard]board.Bitboard
	
	occ_mask board.Bitboard // inner occupancy 
}

// ============================================================================

// Gen new magic numbers
func Gen_all_magics(diag bool) {
	
	var magics [64]Magicsq
	var magics_num board.Bitboard
	var default_shift int
	var bits int

	magics = load_magic(diag) // load existing magic squares data

	// // generate empty magic squares
	// for i := 0; i < 64; i++ {
	// 	magics[i] = Magicsq{
	// 		index: uint(i),
	// 		diag: diag,
	// 		occ_mask: innerOccupancy(uint(i), diag), // inner occupancy
	// 		}
	// }

	for i, magic_sq := range magics {

		default_shift = 64 - len(magic_sq.occ_mask.Index())
		

		magics_num, bits = gen_magic(&magic_sq)

		magics[i].magic = magics_num
		magics[i].shift = default_shift

	}

	
	// save magic squares
	//export_all_magic(magics, diag)

	if false {
		fmt.Println(default_shift, bits, magics_num)
	}

}


// generate magic square for a given square 
func gen_magic(msq *Magicsq) (board.Bitboard, int) {

	var magics_tried uint = 0
	var magics_found uint = 0

	var best_magic board.Bitboard
	
	var default_bits = len(msq.occ_mask.Index())


	for ((magics_found > 0) || (magics_tried < 1_000_000)) { 

		// generate random magic
		magic_bb := board.Bitboard(rand.Uint64())
		msq.magic = magic_bb
		
		magics_tried++
		// check if magic is valid
		if check_magicnum(msq) {

			magics_found++

			best_magic = magic_bb
		}
		
	}
	println("Ind:", msq.index, "Magic found: ", best_magic, " Magic tried: ", magics_tried, "\n")
	
	used_bits := default_bits

	return best_magic, used_bits

}

// check hasmap for magic number
func check_magicnum(msq *Magicsq) bool {

	// generate all occupancy bitboards
	all_occ := allOccupancy(msq.index, msq.diag)

	magic := msq.magic

	hashmap := make(map[board.Bitboard]board.Bitboard)

	var exp_attack_ray func(uint, board.Bitboard) board.Bitboard
	if msq.diag {
		exp_attack_ray = DiagonalRays
		} else {
		exp_attack_ray = SlidingRays
	}

	// hash table checking variables
	var expected board.Bitboard
	var magic_index board.Bitboard

	// create hash table
	for _, occ := range all_occ {

		expected = exp_attack_ray(msq.index, occ) // expected attack ray - using basic sliding rays

		magic_index = (magic * occ) >> msq.shift

		// check if magic number is already in the hash table
		if val, ok := hashmap[magic_index]; ok {
			
			// check if the attack ray is the same
			if val != expected {
				//fmt.Println("Error: magic number is not unique", i)
				return false
			}

		} else {
			hashmap[magic_index] = expected
		}
	}

	return true

}

// ============================================================================
// Staging


func allOccupancy(ind uint, diag bool) []board.Bitboard {
	// all possible bit combinations of fullrays
	
	all_occ := []board.Bitboard{board.Bitboard(0)}
	
	full_rays := fullrays(ind, diag)
	full_rays &= innerOccupancy(ind, diag)
	
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


