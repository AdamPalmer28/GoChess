package make_game

import "chess/board"

type GameState struct {
	Board board.ChessBoard

	WhiteCastle [2]bool
	BlackCastle [2]bool

	White_to_move bool
	Enpassent_ind uint

	MoveList []uint16
	MoveHumanList []string // CLI referencing

	Moveno   uint
	PrevMoves []uint16 // previous moves (0000 000000 000000 form)
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

func Make_move(move uint16, gs *GameState) {

	// get the start and finish squares
	//start := move & (0x3F)
	//finish := (move >> 6) & 0x3F

	special := (move >> 12) & 0xF

	// update piece bitboards (gs.Board)



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




}