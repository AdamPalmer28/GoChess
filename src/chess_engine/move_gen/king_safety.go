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
func check_king_safety(end_sq uint, kingSafety KingSafetyRelBB, knight_ray board.Bitboard,
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
func check_safe_rays(magic_sqs *magic.Magicsq, 
				occ_bb board.Bitboard, rel_attacker_bb board.Bitboard) bool {

	// get the column
	rays := magic.Get_magic_rays(*magic_sqs, occ_bb)

	if (rays & rel_attacker_bb) != 0 {
		return false
	}

	return true
}


// identify check attacks
func check_details(){}

