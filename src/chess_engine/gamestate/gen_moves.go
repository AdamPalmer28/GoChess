package gamestate

import (
	"chess/chess_engine/move_gen"
	"chess/chess_engine/move_gen/magic"
)

// generate moves when king not in check
func (gs *GameState) GenMoves() {

	// reset the moves
	MoveList := move_gen.MoveList{}

	var player move_gen.BoardPerpective = gs.PlayerBoard

	// generate pawn moves
	pawn_moves := move_gen.GenPawnMoves(player.Pawn_bb, gs.White_to_move, 
				gs.Enpass_ind, player.Team_bb, player.Opp_bb)

	MoveList = append(MoveList, pawn_moves...)


	// generate knight moves
	knight_moves := move_gen.GenKnightMoves(player.Knight_bb, 
				&gs.MoveRays.KnightRays, player.Team_bb, player.Opp_bb)

	MoveList = append(MoveList, knight_moves...)


	// generate rook moves
	rook_moves := magic.GenMagicMoves(player.Rook_bb, &gs.MoveRays.Magic.RookMagic,
				player.Team_bb, player.Opp_bb)

	MoveList = append(MoveList, rook_moves...)


	// generate bishop moves
	bishop_moves := magic.GenMagicMoves(player.Bishop_bb, &gs.MoveRays.Magic.BishopMagic,
				player.Team_bb, player.Opp_bb)
	
	MoveList = append(MoveList, bishop_moves...)


	// generate queen moves
	queen_moves := magic.GenMagicMoves(player.Queen_bb, &gs.MoveRays.Magic.BishopMagic,
				player.Team_bb, player.Opp_bb)
	MoveList = append(MoveList, queen_moves...)

	queen_moves = magic.GenMagicMoves(player.Queen_bb, &gs.MoveRays.Magic.RookMagic,
				player.Team_bb, player.Opp_bb)
	MoveList = append(MoveList, queen_moves...)



	// generate king moves
	king_moves := move_gen.GenKingMoves(gs.PlayerKingSaftey, player.Castle_rights,
				&gs.MoveRays.Magic.RookMagic, &gs.MoveRays.Magic.BishopMagic,
				&gs.MoveRays.KnightRays, &gs.MoveRays.KingRays)
	MoveList = append(MoveList, king_moves...)

	
	gs.MoveList = MoveList
}

// ----------------------------------------------------------------------------
// generate moves when king in check
func (gs *GameState) GenCheckMoves() {

	var MoveList []uint

	king_sq := gs.PlayerKingSaftey.King_sq
	// get threat details
	threats, threat_path := move_gen.CheckDetails(gs.PlayerKingSaftey, gs.MoveRays.KnightRays[king_sq], 
				&gs.MoveRays.Magic.RookMagic[king_sq],&gs.MoveRays.Magic.BishopMagic[king_sq])
	
	// threats - must be captures
	// threat_path - must be blocks (i.e moved into)

	if len(threats) == 1 { // calculate blockers

		threat_sq := threats[0]

		defender_moves := move_gen.DefenderMoves(threat_sq, threat_path, 
			gs.PlayerBoard, &gs.MoveRays.KnightRays,
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
func remove_illegal_moves(gs *GameState, pinned_pieces map[uint][]uint) {

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