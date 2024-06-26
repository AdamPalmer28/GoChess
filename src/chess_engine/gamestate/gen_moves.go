package gamestate

import (
	"chess/src/chess_engine/move_gen"
	"chess/src/chess_engine/move_gen/magic"
)

func (gs *GameState) GenMoves() {
	// Move generation when not in check

	var player move_gen.BoardPerpective = gs.PlayerBoard

	// generate pawn moves
	var pawn_caps_ind uint
	if gs.White_to_move {pawn_caps_ind = 0} else {pawn_caps_ind = 1}
	pawn_moves := move_gen.GenPawnMoves(player.Pawn_bb, gs.White_to_move, gs.Enpass_ind, 
				&gs.MoveRays.PawnCapRays[pawn_caps_ind],
				player.Team_bb, player.Opp_bb)

	// generate knight moves
	knight_moves := move_gen.GenKnightMoves(player.Knight_bb, 
				&gs.MoveRays.KnightRays, player.Team_bb, player.Opp_bb)

	// generate rook moves
	rook_moves := magic.GenMagicMoves(
				(player.Rook_bb | player.Queen_bb), 
				&gs.MoveRays.Magic.RookMagic,
				player.Team_bb, player.Opp_bb)
	// generate bishop moves
	bishop_moves := magic.GenMagicMoves(
				(player.Bishop_bb | player.Queen_bb), 
				&gs.MoveRays.Magic.BishopMagic,
				player.Team_bb, player.Opp_bb)

	// generate queen moves
	// queen_moves := magic.GenMagicMoves(player.Queen_bb, &gs.MoveRays.Magic.BishopMagic,
	// 			player.Team_bb, player.Opp_bb)

	// queen_moves = magic.GenMagicMoves(player.Queen_bb, &gs.MoveRays.Magic.RookMagic,
	// 			player.Team_bb, player.Opp_bb)
	

	// generate king moves
	king_moves := move_gen.GenKingMoves(gs.PlayerKingSaftey, player.Castle_rights,
				&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.Magic.BishopMagic,
				&gs.MoveRays.KnightRays, &gs.MoveRays.KingRays)

	// combine move lists together to make complete gamestate move list
	MoveList := append(pawn_moves, knight_moves...)
	MoveList = append(MoveList, rook_moves...)
	MoveList = append(MoveList, bishop_moves...)
	MoveList = append(MoveList, king_moves...)
	
	gs.MoveList = MoveList
}

// ----------------------------------------------------------------------------
func (gs *GameState) GenCheckMoves() {
	// Move generation when InCheck

	var MoveList []uint
	var pawn_caps_ind uint
	if gs.White_to_move {pawn_caps_ind = 0} else {pawn_caps_ind = 1}

	king_sq := gs.PlayerKingSaftey.King_sq
	// get threat details
	threats, threat_path := move_gen.CheckDetails(gs.PlayerKingSaftey, gs.MoveRays.KnightRays[king_sq], 
				gs.MoveRays.PawnCapRays[pawn_caps_ind][king_sq],	
				&gs.MoveRays.Magic.RookMagic[king_sq],&gs.MoveRays.Magic.BishopMagic[king_sq])
	
	// threats - must be captures
	// threat_path - must be blocks (i.e moved into)

	if len(threats) == 1 { // calculate blockers

		threat_sq := threats[0]
		opp_pawn_caps := gs.MoveRays.PawnCapRays[1 - pawn_caps_ind]

		defender_moves := move_gen.DefenderMoves(threat_sq, threat_path, 
			gs.PlayerBoard, &gs.MoveRays.KnightRays, &opp_pawn_caps,
			&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.Magic.BishopMagic)

		MoveList = append(MoveList, defender_moves...)

	} 

	king_moves := move_gen.GenKingMoves(gs.PlayerKingSaftey, 0,
		&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.Magic.BishopMagic,
		&gs.MoveRays.KnightRays, &gs.MoveRays.KingRays)
	MoveList = append(MoveList, king_moves...)

	gs.MoveList = MoveList
}


// ----------------------------------------------------------------------------

// remove illegal moves
func (gs *GameState) RM_IllegalMoves(pinned_pieces map[uint][]uint) {

	var legal_moves []uint

	var start_sq uint
	var end_sq uint

	for _, move := range gs.MoveList {

		// check if move is legal
		start_sq = move & 0b111111

		legal_path, exist := pinned_pieces[start_sq]
		if exist {
			// pinned
			end_sq = (move >> 6) & 0b111111
			for _, ind := range legal_path {
				if ind == end_sq {
					legal_moves = append(legal_moves, move)
					break
				}
			}
		} else {
			// not pinned
			legal_moves = append(legal_moves, move)
			continue
		}
	}
	gs.MoveList = legal_moves
}

// ----------------------------------------------------------------------------


func (gs *GameState) SortMoves() {
	// Sorts movelist - based on move score

	gs.ScoreMoveList = gs.MoveList.GetMoveScore(gs.PlayerBoard, gs.PlayerKingSaftey)
	gs.MoveList = gs.ScoreMoveList.SortMoves()

}