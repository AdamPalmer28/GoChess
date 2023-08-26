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
func check_king_safety(square uint, team_bb board.Bitboard, fwd int,
			magic_str_sqs *[64]magic.Magicsq, magic_diag_sqs *[64]magic.Magicsq, 
			knight_rays *[64]board.Bitboard, king_rays *[64]board.Bitboard,
			opp_pawn board.Bitboard, opp_knight board.Bitboard, opp_bishop board.Bitboard,
			opp_rook board.Bitboard, opp_queen board.Bitboard, opp_king board.Bitboard) bool {

	// get the column
	col := square % 8
	

	// check diag attacks (bishop and queen)
	if !check_safe_rays(square, magic_diag_sqs, team_bb, opp_bishop | opp_queen) {
		return false
	}

	// check straight attacks (rook and queen)
	if !check_safe_rays(square, magic_str_sqs, team_bb, opp_rook | opp_queen) {
		return false
	}

	// check pawn attacks
	pawn_attack_bb := board.Bitboard(0)
	if col != 0 { // left capture
		en_sq := int(square) + fwd - 1
		pawn_attack_bb |= 1 << en_sq
	}
	if col != 7 { // right capture
		en_sq := int(square) + fwd + 1
		pawn_attack_bb |= 1 << en_sq
	}

	if pawn_attack_bb & opp_pawn != 0 {
		return false
	}
	
	// check knight attacks
	knight_attack_bb := knight_rays[square]
	if knight_attack_bb & opp_knight != 0 {
		return false
	}

	return true
}

// check if associated king rays are safe
func check_safe_rays(square uint, magic_sqs *[64]magic.Magicsq, 
				team_bb board.Bitboard, rel_attacker_bb board.Bitboard) bool {


	// get the column
	rays := magic.Get_magic_rays(magic_sqs[square], team_bb)

	if rays & rel_attacker_bb != 0 {
		return false
	}

	return true
}
