package magic

import (
	"chess/chess_engine/board"
	"math"
	"math/rand"
)

type Magicsq struct {
	index uint
	magic board.Bitboard 
	shift int 

	default_shift int
	mapsize int
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
	var bits int

	magics = load_magic(diag) // load existing magic squares data

	// // generate empty magic squares
	// for i := 0; i < 64; i++ {

	// 	inner_occ := innerOccupancy(uint(i), diag)
	// 	magics[i] = Magicsq{
	// 		index: uint(i),
	// 		diag: diag,
	// 		occ_mask: inner_occ, // inner occupancy
	// 		shift: len(inner_occ.Index()),
	// 		default_shift: len(inner_occ.Index()),
	// 		}
	// }


	var valid_magic bool
	var mapsize int
	for i, magic_sq := range magics {

		magics_num, bits, mapsize, valid_magic = gen_magic(&magic_sq)

		if valid_magic {
			magics[i].magic = magics_num
			magics[i].shift = bits
			magics[i].mapsize = mapsize

			// save magic squares
			export_all_magic(magics, diag)
		}
	}
}


// generate magic square for a given square 
func gen_magic(msq *Magicsq) (board.Bitboard, int, int, bool) {

	var magics_tried uint = 0
	var magics_found uint = 0
	var target_shift int = msq.shift

	var new_magic board.Bitboard

	var valid_map bool = false
	var mapsize int = 0
	
	// pre-load all occupancy 
	all_occ := allOccupancy(msq.index, msq.diag)

	if msq.magic != 0 { // if magic number already exists
		//try to improve the magic number
		
		if msq.mapsize < int(math.Pow(2, float64(msq.shift - 1))) {
			println("Ind:", msq.index, " MN - Bit reduction")

			// reduce bits
			target_shift  = target_shift - 1
			for (magics_tried < 10_000_000) { 
				
				new_magic, valid_map, mapsize = find_single_magic(*msq, target_shift, &all_occ)
				magics_tried++

				if (valid_map) { 
					// if magic number is valid
					magics_found++
					println("Ind:", msq.index, "-  (Bit reduction) Magic Found! - Magic tried: ", magics_tried, "// - Mapsize: ", mapsize, " - Old mapsize: ", msq.mapsize)
					break
				}
			}
		} else {
			println("Ind:", msq.index, " MN - Improve")
			// reduce map size 
			for (magics_tried < 1_000_000) { 
				
				new_magic, valid_map, mapsize = find_single_magic(*msq, target_shift, &all_occ)
				magics_tried++

				if (valid_map) && (mapsize < msq.mapsize){ 
					// if magic number is valid
					magics_found++
					println("Ind:", msq.index, "-  (Map reduction) Magic Found! - Magic tried: ", magics_tried, "// - Mapsize: ", mapsize, " - Old mapsize: ", msq.mapsize)
					break
				}
			}
		}


	} else { // find new magic number
		println("Ind:", msq.index, " MN - Search")
		for (magics_tried < 10_000_000) { 
	
			// generate random magic
			new_magic, valid_map, mapsize = find_single_magic(*msq, target_shift, &all_occ)
			
			magics_tried++
			
			if valid_map { 
				// if magic number is valid
				magics_found++
				println("Ind:", msq.index, "-  NEW - Magic Found! - Magic tried: ", magics_tried)

				break
			}
		}
	}

	
	valid_magic := false
	if magics_found != 0 {

		msq.magic = new_magic
		msq.mapsize = mapsize
		msq.shift = target_shift

		valid_magic = true
	}

	return new_magic, target_shift, mapsize, valid_magic

}

func find_single_magic(msq Magicsq, target_shift int, all_occ *[]board.Bitboard) (board.Bitboard, bool, int){

	new_magic := rand_rel_magic_num(msq.occ_mask, target_shift)

	magic_bb := board.Bitboard(new_magic)
	msq.magic = magic_bb
	valid_map, mapsize := check_magicnum(msq, all_occ, target_shift)

	return magic_bb, valid_map, mapsize
}

// ============================================================================
// Magic number generation

func rand_rel_magic_num(occ board.Bitboard, shift int) uint64 {

	all_index := occ.Index()

	// sample of number of bits
	default_bits := len(all_index)
	rand_ind := rand.Perm(default_bits)

	min := shift

	var magic_num uint64 = 0
	var rand_increase int
	for i, index := range all_index {
		dist_to_min := min - int(index)
		if dist_to_min < 0 {

			rand_increase = rand.Int() % (64 - int(index))
			dist_to_min = min
		} else {
			rand_increase = rand_ind[i]
		}

		// shift the bits
		magic_num |= 1 << (rand_increase + dist_to_min) 
	}

	magic_num |= rand.Uint64() | rand.Uint64()

	return magic_num
}

// check hashmap for magic number
func check_magicnum(msq Magicsq, all_occ *[]board.Bitboard, shift int) (bool, int) {

	var mapsize int

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
	for _, occ := range *all_occ {

		expected = exp_attack_ray(msq.index, occ) // expected attack ray - using basic sliding rays

		magic_index = (magic * occ) >> (64 - shift)

		// check if magic number is already in the hash table
		if val, ok := hashmap[magic_index]; ok {
			
			// check if the attack ray is the same
			if val != expected {
				//fmt.Println("Error: magic number is not unique", i)
				return false, 0
			}

		} else {
			hashmap[magic_index] = expected
			mapsize++
		}
	}

	return true, mapsize

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


