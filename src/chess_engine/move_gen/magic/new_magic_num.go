package magic

import (
	"chess/chess_engine/board"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
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

	var valid_magic bool
	var mapsize int

	var magic_new int = 0
	var magic_improve int = 0
	var NumMagics int = 0

	for i, magic_sq := range magics {

		magics_num, bits, mapsize, valid_magic = gen_magic(&magic_sq)

		if valid_magic {
			if magic_sq.magic == 0 {
				magic_new++
			} else {
				magic_improve++
			}

			// update magic square
			magics[i].magic = magics_num
			magics[i].shift = bits
			magics[i].mapsize = mapsize

			// save magic squares
			export_all_magic(magics, diag)
		}

		if magic_sq.magic != 0 {
			NumMagics++
		}
	}

	fmt.Printf("New magic numbers: %d, Improved magic numbers: %d\n", magic_new, magic_improve)
	fmt.Printf("==> Total magic numbers: %d <==\n", NumMagics)
}


// generate magic square for a given square 
func gen_magic(msq *Magicsq) (board.Bitboard, int, int, bool) {
	var wg sync.WaitGroup
	
	var magics_tried int64 = 0
	var magics_found int64 = 0
	var target_shift int = msq.shift

	var BestMagicNum board.Bitboard = 0
	var BestMagicMapsize int = msq.mapsize


	var goRoutineCount int = 12
	wg.Add(goRoutineCount)
	
	// pre-load all occupancy 
	all_occ := allOccupancy(msq.index, msq.diag)

	found_magics := []board.Bitboard{}
	found_magics_mapsize := []int{}

	if msq.magic != 0 { // if magic number already exists
		//try to improve the magic number
		
		if msq.mapsize < int(math.Pow(2, float64(msq.shift - 1))) {
			println("Ind:", msq.index, " MN - Bit reduction")
			target_shift  = target_shift - 1

			for i := 0; i < goRoutineCount; i++ {
				// reduce bits
				go func() {
					defer wg.Done()
					var newMagic board.Bitboard
					var mapsize int
					var valid_map bool = false
					for {
						if atomic.AddInt64(&magics_tried, 1) >= 5_000_000 {
							break
						}
						newMagic, valid_map, mapsize = findSingleMagic(*msq, target_shift, &all_occ)

						if (valid_map) { 
							// if magic number is valid
							atomic.AddInt64(&magics_found, 1)
							found_magics = append(found_magics, newMagic)
							found_magics_mapsize = append(found_magics_mapsize, mapsize)
							fmt.Printf("(Bit reduction) Magic Found! Magics tried: %d // Mapsize: %d // old mapsize: %d\n", magics_tried, mapsize, msq.mapsize)
							
							}
						}
				}()
			}
		} else {
			println("Ind:", msq.index, " MN - Improve")
			// reduce map size 
			for i := 0; i < goRoutineCount; i++ {
				go func() {
					defer wg.Done()
					var newMagic board.Bitboard
					var mapsize int
					var valid_map bool = false
					for {
						if atomic.AddInt64(&magics_tried, 1) >= 1_000_000 {
							break
						}
						
						newMagic, valid_map, mapsize = findSingleMagic(*msq, target_shift, &all_occ)

						if ((valid_map) && (mapsize < msq.mapsize)) { 
							
							atomic.AddInt64(&magics_found, 1)
        					found_magics = append(found_magics, newMagic)
							found_magics_mapsize = append(found_magics_mapsize, mapsize)
							fmt.Printf("(Map reduction) Magic Found! Magics tried: %d // Mapsize: %d // old mapsize: %d\n", magics_tried, mapsize, msq.mapsize)
							
						}
					}
				}()
			}
		}


	} else { // find new magic number
		println("Ind:", msq.index, " MN - Search")
		for i := 0; i < goRoutineCount; i++ {
			go func() {
				defer wg.Done()
				var newMagic board.Bitboard
				var mapsize int
				var valid_map bool = false
				for {
					if atomic.AddInt64(&magics_tried, 1) >= 10_000_000 {
						break
					}
					//msqCopy := *msq
					newMagic, valid_map, mapsize = findSingleMagic(*msq, target_shift, &all_occ)
				
					if valid_map {
						atomic.AddInt64(&magics_found, 1)
						found_magics = append(found_magics, newMagic)
						found_magics_mapsize = append(found_magics_mapsize, mapsize)
						fmt.Printf("New Magic Found! Magics tried: %d // Mapsize: %d // magic: %d\n", magics_tried, mapsize, newMagic)
						
					}
				}
			}()
		}
	}
	wg.Wait()

	
	valid_magic := false
	if magics_found != 0 {

		// find the best magic number
		var best_magics_index int = 0
		//println("Ind:", msq.index, " Magics Found!", len(found_magics), " ",len(found_magics_mapsize))

		
		for i, mapsize := range found_magics_mapsize {
			if mapsize < BestMagicMapsize {
				best_magics_index = i
				BestMagicMapsize = mapsize
			}
		}
		BestMagicNum = found_magics[best_magics_index]
		
		valid_magic = true
	}
	
	fmt.Printf("Ind: %d - Output: valid %t // magic: %d // mapsize: %d\n\n", msq.index, valid_magic, BestMagicNum, BestMagicMapsize)
	return BestMagicNum, target_shift, BestMagicMapsize, valid_magic
}


func findSingleMagic(msq Magicsq, target_shift int, all_occ *[]board.Bitboard) (board.Bitboard, bool, int){

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

