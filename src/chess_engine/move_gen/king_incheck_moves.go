package move_gen

import (
	"chess/chess_engine/board"
	"chess/chess_engine/move_gen/magic"
)


func DefenderMoves(threats uint, threat_paths []uint, 
	Player BoardPerpective, knight_rays *[64]board.Bitboard,
	magic_str_sqs *[64]magic.Magicsq, magic_diag_sqs *[64]magic.Magicsq) []uint {
	
	var movelist []uint
	var defenders_bb board.Bitboard
	
	// moves
	for _, sq := range threat_paths {
		// get the king safety struct

		// get the knight attacks
		opp_knight_bb := knight_rays[sq] & Player.Knight_bb

		// straight attacks
		attack_ray := magic.Get_magic_rays(magic_str_sqs[sq],
			(Player.Team_bb | Player.Opp_bb))
		attack_ray &= Player.Rook_bb | Player.Queen_bb

		// diag attacks
		diag_rays := magic.Get_magic_rays(magic_diag_sqs[sq],
			(Player.Team_bb | Player.Opp_bb))
		diag_rays &= Player.Bishop_bb | Player.Queen_bb
		
		defenders_bb |= (opp_knight_bb | attack_ray | diag_rays)
		for _, ind := range(defenders_bb.Index()) {
			movelist = append(movelist, (sq << 6) | ind)
		}

		// pawn moves
			// double fwd
		double := uint(int(sq) - 2 * Player.Fwd)
		if (Player.P_start_row == double / 8) && (Player.Pawn_bb & (1 << double) != 0) {
			movelist = append(movelist, (1 << 12 | sq << 6 | double))
		}
			// single fwd
		single := uint(int(sq) - Player.Fwd)
		if (Player.Pawn_bb & (1 << single) != 0) {
			movelist = append(movelist, (sq << 6 | single))
		}
	}

			



	return movelist
}