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


func special_move(move uint) uint {
	return (move >> 12)
}

func (moves *MoveList) SortMoves() {

	// priority order: promotion, capture, castle, double pawn push
	
	// they are already in numerical value
	// so we just can sort them by their value (highest to lowest)

	sort.Slice(*moves, func(i, j int) bool {

		return (special_move((*moves)[i]) > special_move((*moves)[j]))
		
	})


}

func (moves *MoveList) ImportantMoves() MoveList {

	// priority order: promotion, capture, castle
	
	// they are already in numerical value
	// so we just can sort them by their value (highest to lowest)

	important_moves := make(MoveList, 0)

	for _, move := range *moves {
		if special_move(move) > 1 {
			important_moves = append(important_moves, move)
		}
	}
	return important_moves
}

