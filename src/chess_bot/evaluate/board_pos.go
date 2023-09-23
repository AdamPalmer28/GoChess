package evaluate

/*
Useful board segments


*/
import (
	"chess/chess_engine/board"
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
