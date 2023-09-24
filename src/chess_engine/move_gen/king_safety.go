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
func CheckKingSafety(kingSafety KingSafetyRelBB, knight_ray board.Bitboard,
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
func CheckDetails(kingSafety KingSafetyRelBB, knight_ray board.Bitboard, pawn_caps board.Bitboard,
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
	opp_pawn_bb = pawn_caps & kingSafety.Opp_pawn_bb

	if opp_pawn_bb != 0 {
		no_of_checks++
		threat_sq = opp_pawn_bb.Index()[0]
		threats = append(threats, threat_sq)
	}

	return threats, threat_path
}



// path is either +/- 1,7,8,9

// get the path between sq and end_sq
func get_ray_paths(sq int, end_sq int) board.Bitboard {

	diff := end_sq - sq
	var step int = 1 // step for creating bb

	
	if (diff % 8 == 0) {
		step = 8
	} else if (diff % 9 == 0) { // must be before 7
		step = 9
	} else if (diff % 7 == 0) && ( (end_sq/8 - sq/8) != 0) { // must be before 1
		step = 7
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


// ============================================================================
// - used in gs.GetCheck() to identify check and pinned pieces
// ============================================================================
 
// returns if king is *in check* and *pinned pieces*
func BoardKingAnalysis(kingSafety KingSafetyRelBB, knight_ray board.Bitboard,
	magic_str_sq *magic.Magicsq, magic_diag_sq *magic.Magicsq) (bool, map[uint][]uint) {

	var result bool = true
	var pinned_pieces map[uint][]uint = make(map[uint][]uint) // pinned piece -> list of path_sqs


	// check diag attacks (bishop and queen) - potential pin
	diag_safe, diag_pin_map := check_rays(magic_diag_sq, kingSafety.Team_bb, kingSafety.Opp_bb,
			(kingSafety.Opp_bishop_bb | kingSafety.Opp_queen_bb))

	// check straight attacks (rook and queen)
	straight_safe, straight_pin_map := check_rays(magic_str_sq, kingSafety.Team_bb, kingSafety.Opp_bb,
			(kingSafety.Opp_rook_bb | kingSafety.Opp_queen_bb))
		
	if !(diag_safe && straight_safe) {
		result = false
	}

	// combine pin maps
	for key, value := range diag_pin_map {
		pinned_pieces[key] = value
	}
	for key, value := range straight_pin_map {
		pinned_pieces[key] = value
	}


	// check knight attacks
	if (knight_ray & kingSafety.Opp_knight_bb) != 0 {
		result = false
	}

	if result {
		
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
			result = false
		}
	}

	return result, pinned_pieces
	}

// check if associated king rays are safe - and get potential pinned pieces
func check_rays(magic_sq *magic.Magicsq, 
	team_bb board.Bitboard, opp_bb board.Bitboard, 
	rel_attacker_bb board.Bitboard) (bool, map[uint][]uint) {

	var safe bool = true
	var pinned map[uint][]uint = make(map[uint][]uint) // pinned piece -> list of path_sqs

	occ_bb := team_bb | opp_bb
	// get the column
	rays := magic.Get_magic_rays(*magic_sq, occ_bb)

	if (rays & rel_attacker_bb) != 0 {
		safe = false
	}

	// -----------------------------------------------------------------------
	// potential pins
	king_sq := magic_sq.Index

	potential_pin := rays & team_bb
	new_occ := occ_bb ^ potential_pin

	rays_without_pieces  := magic.Get_magic_rays(*magic_sq, new_occ) // rays without pieces
	pinning_threats := (rays_without_pieces & rel_attacker_bb) 


	if (pinning_threats) != 0 {
		var pinned_paths board.Bitboard // path of pinned pieces
		var confirmed_pin board.Bitboard // confirmed pin

		for _, threat := range pinning_threats.Index() { 
			// loop through identified threats
			
			pinned_paths = get_ray_paths(int(king_sq), int(threat))

			
			confirmed_pin = (pinned_paths & potential_pin)
			if confirmed_pin != 0 {
				// if potential pin piece is on path
				pinned[confirmed_pin.Index()[0]] = append(pinned_paths.Index(), threat)
				
			}

		}
	}

	return safe, pinned
}