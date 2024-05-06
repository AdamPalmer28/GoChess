package move_gen

import (
	"chess/src/chess_engine/board"
	"sort"
)

/*
Moves are represented as a 16 bit integer
0000 000000 000000
special | finish | start

0000 - special moves (info below)
000000 - index of square

special moves
0000 - quite move
0001 - double pawn push
0010 - king side castle
0011 - queen side castle
0100 - capture
0101 - enpassent capture
1000 - promotion knight
1001 - promotion bishop
1010 - promotion rook
1011 - promotion queen
1100 - promotion knight capture
1101 - promotion bishop capture
1110 - promotion rook capture
1111 - promotion queen capture

*/

type MoveList []uint

// move list for sorting [move, score] uint, float
type ScoreMoveList []struct {
	Move uint
	Score float64
}

type BoardPerpective struct {
	// friendly pieces
	Pawn_bb board.Bitboard
	Knight_bb board.Bitboard
	Bishop_bb board.Bitboard
	Rook_bb board.Bitboard
	Queen_bb board.Bitboard
	King_bb board.Bitboard
	// opp pieces
	Opp_pawn_bb board.Bitboard
	Opp_knight_bb board.Bitboard
	Opp_bishop_bb board.Bitboard
	Opp_rook_bb board.Bitboard
	Opp_queen_bb board.Bitboard
	Opp_king_bb board.Bitboard

	Team_bb board.Bitboard
	Opp_bb board.Bitboard

	Fwd int
	P_start_row uint
	Castle_rights uint
	Enpass_ind uint
}



func (moves MoveList) GetMoveScore(bp BoardPerpective, ks KingSafetyRelBB) ScoreMoveList {
	// Give each move a quick estimate of score value - which will be used to sort moves

	sorted_moves := make(ScoreMoveList, len(moves))

	for i, move := range moves {
		sorted_moves[i].Move = move
		score := 0.0

		// move details
		special_move := move >> 12
		end_sq := (move >> 6) & 0b111111
		start_sq := (move) & 0b111111

		piece_moved := move_getPiece(start_sq, [6]*board.Bitboard{&bp.Pawn_bb, &bp.Knight_bb, &bp.Bishop_bb, &bp.Rook_bb, &bp.Queen_bb, &bp.King_bb})
		
		// promotion 
		if special_move > 0b1000 { 

			if special_move & 0b0011 == 0b0011 { // queen promotion
				score += 9.0
			} else if special_move & 0b0011 == 0b0000 { // knight promotion
				score += 5.0
			} else { // rook or bishop promotion - not worth searching much
				score += 1.5 
			} 
		}
		// capture
		if special_move >= 0b0100 { 
			score += 1.0

			// player piece, opp piece
			piece_captured := move_getPiece(end_sq, [6]*board.Bitboard{&bp.Opp_pawn_bb, &bp.Opp_knight_bb, &bp.Opp_bishop_bb, &bp.Opp_rook_bb, &bp.Opp_queen_bb, &bp.Opp_king_bb})
			piece_value := [6]float64{1, 3, 3, 5, 9, 100}

			score += piece_value[piece_captured] - piece_value[piece_moved]
		}

		// king safety

		sorted_moves[i].Score = score
	}

	return sorted_moves
}


func (move *ScoreMoveList) SortMoves() MoveList {
	// Sort MoveList based on moveScore

	sort.Slice(*move, func(i, j int) bool {
		return (*move)[i].Score > (*move)[j].Score
	})

	// make move list
	movelist := make(MoveList, len(*move))
	for i, m := range *move {
		movelist[i] = m.Move
	}

	return movelist
}

func (moves *ScoreMoveList) ImportantMoves() MoveList {
	// only look at moves with a score above a certain threshold

	important_moves := make(MoveList, 0)

	for _, m := range *moves {
		if m.Score > 2 {
			important_moves = append(important_moves, m.Move)
		}
	}
	return important_moves
}

// ----------------------------------------------------------------------------
// helper functions

// Move to board
func move_getPiece(sq uint, BB_list [6]*board.Bitboard) uint {

	// update piece bitboards
	var piece_moved uint
	for ind, BB := range BB_list {

		// ? would this be faster with BB != 0 
		if (*BB & (1 << sq) != 0) {
			piece_moved = uint(ind)
			break
		}
	}
	return piece_moved
}