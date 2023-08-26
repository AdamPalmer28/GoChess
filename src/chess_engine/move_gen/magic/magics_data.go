package magic

import (
	"chess/chess_engine/board"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// Construct the relative file path from the current directory
const strt_magic_data_path = "data/magics.json"
const diag_magic_data_path = "data/diag_magics.json"

type magicdata struct {
	Index uint
	Magic board.Bitboard
	Shift int
	default_shift int
	Diag bool
}

// create hash table for magic square
func magic_attack_rays(msq *Magicsq) {

	index := msq.index
	diag := msq.diag
	magic_num := msq.magic
	shift := msq.shift

	// attack rays for magic square
	attack_rays := make(map[board.Bitboard]board.Bitboard)

	// all occupancy bitboards
	all_occ := allOccupancy(msq.index, msq.diag)


	var exp_attack_ray func(uint, board.Bitboard) board.Bitboard
	if diag {
		exp_attack_ray = DiagonalRays
		} else {
		exp_attack_ray = SlidingRays
	}

	var expected board.Bitboard

	// create hash table
	for _, occ := range all_occ {

		expected = exp_attack_ray(index, occ) // expected attack ray - using basic sliding rays
		
		// expected.Print()
		// occ.Print()
		// println(magic_num, shift)
		magic_index := (occ * magic_num) >> shift


		// add attack ray to hash table
		attack_rays[magic_index] = expected

	}
	msq.attack_rays = attack_rays // hash table for magic square
}


// ============================================================================
// Load magic data

// Transform magicdata to magicsq 
func create_magicsq(data magicdata) Magicsq {

	msq := Magicsq{
		index: data.Index, 
		diag: data.Diag,
		magic: data.Magic,
		shift: data.Shift, // 64 - data.Shift 
		//default_shift: data.default_shift,

		occ_mask: innerOccupancy(data.Index, data.Diag), // inner occupancy
		}

	magic_attack_rays(&msq)

	return msq
}

func Load_all_magicsq() ([64]Magicsq, [64]Magicsq) {


	strt := load_magic(false)
	diag := load_magic(true)

	return strt, diag
}

func load_magic(diag bool) [64]Magicsq{
	
	// Get the path of the current source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)


	var file_name string
	if diag {
		file_name = diag_magic_data_path
	} else {
		file_name = strt_magic_data_path
	}
	file_name = filepath.Join(dir, file_name)
	
	// open file
	file, err := os.Open(file_name)
	if err != nil {fmt.Println("Error opening file:", err)}

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
			file_name = diag_magic_data_path
			} else {
				file_name = strt_magic_data_path
			}
			
			file, err := os.Create(file_name)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			encoding := json.NewEncoder(file)
			
			// create magic data
			export_data := create_magicdata(msq)
	
			
			// export magic data to json
			err = encoding.Encode(export_data)
			if err != nil {
				fmt.Println("Error encoding:", err)
				return
			}
	}

			
	// convert Magicsq to magicdata
	func create_magicdata(all_msq [64]Magicsq) [64]magicdata {
	
		var data [64]magicdata
	
		for i := 0; i < 64; i++ {
			data[i].Index = all_msq[i].index
			data[i].Magic = all_msq[i].magic
			data[i].Shift = 64 - all_msq[i].shift
			data[i].default_shift = 64 - len(all_msq[i].occ_mask.Index()) // default shift
			data[i].Diag = all_msq[i].diag
		}
		return data
	}
