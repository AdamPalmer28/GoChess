package move_gen

import (
	"chess/chess_engine/board"
	"chess/chess_engine/move_gen/magic"
)


func DefenderMoves(threat_sq uint, threat_paths []uint, 
	Player BoardPerpective, knight_rays *[64]board.Bitboard,
	magic_str_sqs *[64]magic.Magicsq, magic_diag_sqs *[64]magic.Magicsq) []uint {
		
	var movelist []uint
	var moveno uint
	var defenders_bb board.Bitboard
	
	var opp_knight_bb board.Bitboard
	var attack_ray board.Bitboard
	var diag_rays board.Bitboard

	var dirConstant int = (Player.Fwd/8)
	
	// ------------------------------------------------------------------------
	// moves
	for _, sq := range threat_paths {
		// get the king safety struct

		// get the knight attacks
		opp_knight_bb = knight_rays[sq] & Player.Knight_bb
		
		// straight attacks
		attack_ray = magic.Get_magic_rays(magic_str_sqs[sq],
			(Player.Team_bb | Player.Opp_bb))
		attack_ray &= Player.Rook_bb | Player.Queen_bb

		// diag attacks
		diag_rays = magic.Get_magic_rays(magic_diag_sqs[sq],
			(Player.Team_bb | Player.Opp_bb))
		diag_rays &= Player.Bishop_bb | Player.Queen_bb
		
		defenders_bb |= (opp_knight_bb | attack_ray | diag_rays)
		for _, ind := range(defenders_bb.Index()) {
			movelist = append(movelist, (sq << 6) | ind)
		}

		// pawn moves
		double := uint(int(sq) - 2 * Player.Fwd) // double pawn push
		single := uint(int(sq) - Player.Fwd) // single pawn push

		if (Player.P_start_row == double / 8) && (Player.Pawn_bb & (1 << double) != 0) {
			// double pawn push
			movelist = append(movelist, (1 << 12 | sq << 6 | double))

		} else if (Player.Pawn_bb & (1 << single) != 0) { 
			// single pawn push
			moveno = (sq << 6 | single)

			if (7 - Player.P_start_row) == single / 8 { // promotion
				moveno |= 0b1000 << 12
				promo_list := promotion(moveno)
				movelist = append(movelist, promo_list[:]...)
			} else {
				movelist = append(movelist, moveno)
			}
		}
	}

	// -------------------------------------------------------------------------
	// captures moves (capture threats)

	// knights caps
	opp_knight_bb = knight_rays[threat_sq] & Player.Opp_bb

	// straight caps
	attack_ray = magic.Get_magic_rays(magic_str_sqs[threat_sq],
		(Player.Team_bb | Player.Opp_bb))
	attack_ray &= Player.Rook_bb | Player.Queen_bb

	// diag caps
	diag_rays = magic.Get_magic_rays(magic_diag_sqs[threat_sq],
		(Player.Team_bb | Player.Opp_bb))
	diag_rays &= Player.Bishop_bb | Player.Queen_bb

	defenders_bb = (opp_knight_bb | attack_ray | diag_rays)
	for _, ind := range(defenders_bb.Index()) { // capture moves
		movelist = append(movelist, 0b0100 << 12 | (threat_sq << 6) | ind)
	}

	// pawn caps
	// ? does this work???
	if int(threat_sq / 8) * (dirConstant) > int(Player.P_start_row) * (dirConstant){

		cap_bb := get_pawn_attack(threat_sq, -Player.Fwd)
		cap_bb &= Player.Pawn_bb
		cap_sq := cap_bb.Index()
		
	
		for _, sq := range cap_sq {
			moveno = (0b0100 << 12 | (threat_sq << 6) | sq)
	
			if (7 - Player.P_start_row) == sq / 8 { // promotion
				moveno |= 0b1000 << 12
				promo_list := promotion(moveno)
				movelist = append(movelist, promo_list[:]...)
			} else {
				movelist = append(movelist, moveno)
			}
		}
		
		if Player.Enpass_ind < 64 { // enpassent capture
			cap_bb := get_pawn_attack(Player.Enpass_ind, -Player.Fwd)
			cap_bb &= Player.Pawn_bb
			cap_sq := cap_bb.Index()
	
			for _, sq := range cap_sq {
				if Player.Enpass_ind == sq {
					moveno = (0b0101 << 12 | (threat_sq << 6) | sq)
					movelist = append(movelist, moveno)
				}
			}		
		}
	}
	



	return movelist
}