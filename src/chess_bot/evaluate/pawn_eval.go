package evaluate

import (
	"chess/src/chess_engine/board"
	"math"
)

/*
Pawn eval:
	- centre control
	- pomotion potential
		- blocked
		- past pawns
	- pawn structure
		- chain
		- doubled
		- isolated
		- backward
*/

const (
	centre_sq_val float64 = 0.07

	//blocked_pawn_val float64 = -0.1
	past_pawn_val float64 = 0.05

	// pawn structure
	chain_val float64 = 0.04
	doubled_val float64 = -0.04
	isolated_val float64 = -0.08
	backward_val float64 = -0.05

)


func PawnEval(cb board.ChessBoard, pawn_cap_rays *[2][64]board.Bitboard) [2]float64 { 

	score := [2]float64{0.0, 0.0}
	white_cap_rays := pawn_cap_rays[0]
	black_cap_rays := pawn_cap_rays[1]

	wp := *cb.WhitePawns
	bp := *cb.BlackPawns

	var blocked bool
	var row int
	var col int
	var rows_away int
	var adj_files board.Bitboard
	var adj_files_past board.Bitboard

	// ------------------------------------------------------------------------
	// white
	var white_attacks board.Bitboard
	for _, ind := range wp.Index() {
		// pawn info
		if (cb.Black & (1 << ind + 8)) != 0 {blocked = true} else {blocked = false}
		row = int(ind / 8)
		col = int(ind % 8)
		rows_away = 7 - row
		
		white_attacks |= white_cap_rays[ind] // get all pawn attacks
		
		adj_files = Adjacent_files(ind) // get all adjacent files
		adj_files_past = (adj_files << (8 * (row+1))) // past pawns adjacent files

		// past pawn
		if (bp & (adj_files_past)) == 0 {
			if blocked {
				score[0] += past_pawn_val * math.Pow( float64(6 - rows_away), 2) / 2
			} else {
				score[0] += past_pawn_val * math.Pow( float64(6 - rows_away), 2)
			}
		}

		// doubled
		if (wp & board.FileA << col) != 0 {
			score[0] += doubled_val
		}
		// isolated
		if (wp & adj_files) == 0 {
			score[0] += isolated_val
		}
		// backward
		if (wp & (adj_files & ^adj_files_past)).Count() == 1 {
			score[0] += backward_val
		}

	}
	// chain
	chains := wp & white_attacks
	score[0] += chain_val * float64(chains.Count())

	
	// ------------------------------------------------------------------------
	// black

	var black_attacks board.Bitboard
	for _, ind := range bp.Index() {
		// pawn info
		if (cb.White & (1 << ind - 8)) != 0 {blocked = true} else {blocked = false}
		row = int(ind / 8)
		col = int(ind % 8)
		rows_away = row

		black_attacks |= black_cap_rays[ind] // get all pawn attacks

		adj_files = Adjacent_files(ind) // get all adjacent files
		adj_files_past = (adj_files >> (8 * (7 - row + 1))) // past pawns adjacent files 

		// past pawn
		if (wp & (adj_files_past)) == 0 {
			if blocked {
				score[1] += past_pawn_val * math.Pow( float64(6 - rows_away), 2) / 2
			} else {
				score[1] += past_pawn_val * math.Pow( float64(6 - rows_away), 2)
			}
		}

		// doubled
		if (bp & board.FileA << col) != 0 {
			score[1] += doubled_val
		}
		// isolated
		if (bp & adj_files) == 0 {
			score[1] += isolated_val
		}
		// backward
		if (bp & (adj_files & ^adj_files_past)).Count() == 1 {
			score[1] += backward_val
		}
	}
	// chain
	chains = bp & black_attacks
	score[1] += chain_val * float64(chains.Count())

	// ------------------------------------------------------------------------
	PawnCenterScore := pawn_centre_control(wp, white_attacks, bp, black_attacks)
	score[0] += PawnCenterScore[0]
	score[1] += PawnCenterScore[1]

	return score

}

// eval function for pawn centre control
func pawn_centre_control(wp board.Bitboard, wp_cap board.Bitboard,
						bp board.Bitboard, bp_cap board.Bitboard) [2]float64 {
	score := [2]float64{0.0, 0.0}

	// white
		// postion 
	pawns_on_centre := wp & (Mid4Centre & ^board.Rank3)
	n := pawns_on_centre.Count()
	score[0] += float64(n) * centre_sq_val / 2

		// attacks
	attacks_on_centre := wp_cap & (Mid4Centre & ^board.Rank3)
	n = attacks_on_centre.Count()
	score[0] += float64(n) * centre_sq_val

	// black
		// postion
	pawns_on_centre = bp & (Mid4Centre & ^board.Rank6)
	n = pawns_on_centre.Count()
	score[1] += float64(n) * centre_sq_val / 2

		// attacks
	attacks_on_centre = bp_cap & (Mid4Centre & board.Rank6)
	n = attacks_on_centre.Count()
	score[1] += float64(n) * centre_sq_val
				
	return score
}


// =================================================================================================
// Helper function


// returns bitboard of adjacent files to the given index
func Adjacent_files(ind uint) board.Bitboard {

	file_index := int(ind % 8)

	// middle, left, right
	adja_file :=  (board.FileA << file_index) | (board.FileA << max(0, file_index - 1)) | (board.FileA << min(7, file_index + 1))

	return adja_file
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b 
	}
	return a
}