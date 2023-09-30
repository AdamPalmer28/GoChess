package evaluate

/*
Useful board segments


*/
import (
	"chess/chess_engine/board"
	"chess/chess_engine/gamestate"
)

const (
	// files
	MidFile board.Bitboard = board.FileD | board.FileE
	Mid4File board.Bitboard = board.FileC | board.FileD | board.FileE | board.FileF

	// ranks
	MidRows board.Bitboard = board.Rank4 | board.Rank5
	Mid4Rows board.Bitboard = board.Rank3 | board.Rank4 | board.Rank5 | board.Rank6
	
	// centre
	MidCentre board.Bitboard = MidFile & MidRows
	Mid4Centre board.Bitboard = Mid4File & Mid4Rows 
	
	// sides
	WhiteSide board.Bitboard = board.Rank1 | board.Rank2 | board.Rank3 | board.Rank4
	BlackSide board.Bitboard = board.Rank5 | board.Rank6 | board.Rank7 | board.Rank8

	// cross
	Cross4 board.Bitboard = Mid4Rows | Mid4File
	Corners board.Bitboard = ^Cross4

	// edge
	Edge board.Bitboard = board.FileA | board.FileH | board.Rank1 | board.Rank8
)

/*
Get areas of the board which are of high activity / importance this is used for
significant mappings in evaluation functions
*/

type BoardActivity struct {

	// high activity pieces
	wHighActivty [64]float64
	bHighActivty [64]float64

	// king critical & threats
	wKingThreats [64]float64
	bKingThreats [64]float64

	// critical sqaures for knights to use
	wCriticalAtx [64]float64
	bCriticalAtx [64]float64

	// knight outposts
	wKnightOutposts [64]float64
	bKnightOutposts [64]float64	 

}


func getBoardActivity(gs *gamestate.GameState, mv_ray EvalMoveRays) BoardActivity {

	var ba BoardActivity
	
	return ba
}