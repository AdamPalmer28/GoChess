package move_gen

import (
	"chess/chess_engine/board"
	"chess/chess_engine/move_gen/magic"
)

/*
Possible double check combinations:
	- knight and bishop (diag)
	- knight and rook (straight)
	- rook (straight) and bishop (diag)
	- pawn and rook (diag)
*/

// check if king check
func Check_king_safety(kingSafety KingSafetyRelBB, knight_ray board.Bitboard,
			magic_str_sq *magic.Magicsq, magic_diag_sq *magic.Magicsq) bool {
					
	// check diag attacks (bishop and queen)
	if !check_safe_rays(magic_diag_sq, (kingSafety.Team_bb_no_king | kingSafety.Opp_bb), 
						(kingSafety.Opp_bishop_bb | kingSafety.Opp_queen_bb)) {
		return false
	}
	
	// check straight attacks (rook and queen)
	if !check_safe_rays(magic_str_sq, (kingSafety.Team_bb_no_king | kingSafety.Opp_bb), 
						(kingSafety.Opp_rook_bb | kingSafety.Opp_queen_bb)) {
		return false
	}

	// check knight attacks
	if (knight_ray & kingSafety.Opp_knight_bb) != 0 {
		return false
	}
	
	// check pawn attacks
	pawn_attack_bb := board.Bitboard(0)
	col := kingSafety.King_sq % 8
	if col != 0 { // left capture
		en_sq := int(kingSafety.King_sq) + kingSafety.Fwd - 1
		pawn_attack_bb |= 1 << en_sq
	}
	if col != 7 { // right capture
		en_sq := int(kingSafety.King_sq) + kingSafety.Fwd + 1
		pawn_attack_bb |= 1 << en_sq
	}

	if pawn_attack_bb & kingSafety.Opp_bb != 0 {
		return false
	}
	

	return true
}

// check if associated king rays are safe
func check_safe_rays(magic_sq *magic.Magicsq, 
				occ_bb board.Bitboard, rel_attacker_bb board.Bitboard) bool {

	// get the column
	rays := magic.Get_magic_rays(*magic_sq, occ_bb)

	if (rays & rel_attacker_bb) != 0 {
		return false
	}

	return true
}

// ============================================================================
// detailed check
// ============================================================================

// identify check attacks and squares associated with check
func CheckDetails(kingSafety KingSafetyRelBB, knight_ray board.Bitboard,
			magic_str_sq *magic.Magicsq, magic_diag_sq *magic.Magicsq) ([]uint, []uint){

	var no_of_checks uint
	var threats []uint
	var threat_sq uint
	var threat_path []uint
	var path board.Bitboard
	var attack_ray board.Bitboard

	var opp_straight_bb board.Bitboard
	var opp_diag_bb board.Bitboard
	var opp_knight_bb board.Bitboard
	var opp_pawn_bb board.Bitboard

	// get the king safety struct


	// get the knight attacks
	opp_knight_bb = knight_ray & kingSafety.Opp_knight_bb

	if opp_knight_bb != 0 {
		no_of_checks++
		threat_sq = opp_knight_bb.Index()[0]
		threats = append(threats, threat_sq)
	}

	// get the diag attacks
	diag_rays := magic.Get_magic_rays(*magic_diag_sq, 
		(kingSafety.Team_bb_no_king | kingSafety.Opp_bb))
	opp_diag_bb = diag_rays & (kingSafety.Opp_bishop_bb | kingSafety.Opp_queen_bb)

	
	if opp_diag_bb != 0 {
		no_of_checks++
		threat_sq = opp_diag_bb.Index()[0]
		threats = append(threats, threat_sq)
		
		path = get_ray_paths(int(kingSafety.King_sq), int(threat_sq))
		threat_path = append(threat_path, path.Index()...)

		if no_of_checks == 2 { // early return
			return threats, threat_path
		}
	}

	// get the straight attacks
	attack_ray = magic.Get_magic_rays(*magic_str_sq, 
		(kingSafety.Team_bb_no_king | kingSafety.Opp_bb))
	opp_straight_bb = attack_ray & (kingSafety.Opp_rook_bb | kingSafety.Opp_queen_bb)
	
	if opp_straight_bb != 0 {
		no_of_checks++
		threat_sq = opp_straight_bb.Index()[0]
		threats = append(threats, threat_sq)
		
		path = get_ray_paths(int(kingSafety.King_sq), int(threat_sq))
		threat_path = append(threat_path, path.Index()...)

		if no_of_checks == 2 { // early return
			return threats, threat_path
		}
	}


	// get the pawn attacks
	pawn_attack_bb := get_pawn_attack(kingSafety.King_sq, kingSafety.Fwd)
	opp_pawn_bb = pawn_attack_bb & kingSafety.Opp_pawn_bb

	if opp_pawn_bb != 0 {
		no_of_checks++
		threat_sq = opp_pawn_bb.Index()[0]
		threats = append(threats, threat_sq)
	}

	return threats, threat_path
}



// path is either +/- 1,7,8,9

// get the path between sq and end_sq
func get_ray_paths(end_sq int, sq int) board.Bitboard {

	diff := end_sq - sq
	var step int = 1 // step for creating bb
	
	if (diff % 7 == 0) && (sq % 8 != 0) {
		step = 7
	} else if (diff % 8 == 0) {
		step = 8
	} else if (diff % 9 == 0) {
		step = 9
	} 
	// else step = 1
	
	if diff < 0 {
		step *= -1 
	} 

	var ray_bb board.Bitboard = 0
	var curr_sq int = sq + step // set as first square

	for curr_sq != end_sq {

		ray_bb |= 1 << curr_sq
		curr_sq += step
	} 
	
	return ray_bb
}