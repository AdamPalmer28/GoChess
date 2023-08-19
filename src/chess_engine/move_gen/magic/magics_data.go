package magic

import (
	"chess/chess_engine/board"
	"encoding/json"
	"fmt"
	"os"
)

type magicdata struct {
	index uint
	magic board.Bitboard
	shift int
	diag bool
}

// Transform magic data to magicsq  // ? NEEDS TO BE REDESIGNED
func create_magicsq(data magicdata) Magicsq {

	msq := Magicsq{
		index: data.index, 
		diag: data.diag,
		magic: data.magic,
		shift: data.shift,
		}

	msq.occ_mask = innerOccupancy(data.index) // inner occupancy

	magic_attack_rays(&msq)

	return msq
}

// create hash table for magic square
func magic_attack_rays(msq *Magicsq) bool{

	index := msq.index
	diag := msq.diag
	magic_num := msq.magic
	shift := msq.shift

	var attack_rays []board.Bitboard // attack rays for magic square

	// all occupancy bitboards
	all_occ := allOccupancy(msq.index, msq.diag)


	var exp_attack_ray_func func(uint, board.Bitboard) board.Bitboard
	if diag {
		exp_attack_ray_func = SlidingRays
	} else {
		exp_attack_ray_func = DiagonalRays
	}

	var expected board.Bitboard

	// create hash table
	for _, occ := range all_occ {

		expected = exp_attack_ray_func(index, occ) // expected attack ray - using basic sliding rays

		magic_index := (occ * magic_num) >> shift

		// add attack ray to hash table
		attack_rays[magic_index] = expected

	}
	msq.attack_rays = attack_rays
}
// ============================================================================
// Load magic data



func load_magic(diag bool) [64]Magicsq{

	var file_name string
	if diag {
		file_name = "diag_magics.json"
	} else {
		file_name = "magics.json"
	}

	
	// open file
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	decoding := json.NewDecoder(file)
	
	// decode json
	data := [64]magicdata{}

	err = decoding.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding:", err)
	}

	// create magic squares
	var magicsquares [64]Magicsq

	for i := 0; i < 64; i++ {
		magicsquares[i] = create_magicsq(data[i])
	}
	
	return magicsquares
}

// ============================================================================
// Export magic data

func export_all_magic(msq [64]Magicsq, diag bool) {

	var file_name string
	if diag {
		file_name = "diag_magics.json"
	} else {
		file_name = "magics.json"
	}

	file, err := os.Create(file_name)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	encoding := json.NewEncoder(file)
	
	var export_data magicdata
	for i := 0; i < 64; i++ {

		export_data.index = msq[i].index
		export_data.magic = msq[i].magic
		export_data.shift = msq[i].shift
		export_data.diag = msq[i].diag

		// export magic data to json
		
		err = encoding.Encode(export_data)
		if err != nil {
			fmt.Println("Error encoding:", err)
			return
		}
	}
	fmt.Println("Exported magic data to", file_name)
}