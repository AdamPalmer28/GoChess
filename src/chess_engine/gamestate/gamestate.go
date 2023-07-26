package gamestate

import "chess/chess_engine/board"

type GameState struct {
	Board board.ChessBoard

	WhiteCastle [2]bool
	BlackCastle [2]bool

	White_to_move bool
	Enpassent_ind uint

	MoveList []uint
	MoveHumanList []string // CLI referencing

	Moveno   uint
	PrevMoves []uint // previous moves (0000 000000 000000 form)
	Cap_pieces [][2]int // [move number][piece type]

}

var PieceValLookup = map[int]string{
	0: "P",
	1: "K",
	2: "B",
	3: "R",
	4: "Q",
	5: "K",
}

func (gs *GameState) Make_move(move uint)  {

	// get the start and finish squares

	special := (move >> 12) & 0xF

	// update piece bitboards (gs.Board)
	piece_moved, cap_piece := gs.Board.Move(move, gs.White_to_move)

	piece_moved_str := PieceValLookup[int(piece_moved)]
	cap_piece_str := PieceValLookup[int(cap_piece)]

	println(piece_moved_str, cap_piece_str)


	// update castling rights


	// update enpassent index

	// capture piece
	if (special & 4) > 0 {
		piece := 0 // placeholder
		gs.Cap_pieces = append(gs.Cap_pieces, [2]int{ int(gs.Moveno), piece})
	}

	// update previous moves
	gs.PrevMoves = append(gs.PrevMoves, move)
	
	// update move number
	gs.Moveno++

	
	// change move color
	gs.White_to_move = !gs.White_to_move




}