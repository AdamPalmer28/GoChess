package src

// ============================================================================
// Data structure

type ChessData struct {
	Message string `json:"message"`
	GameState GameStateData `json:"gamestate"`
}

type GameStateData struct { // Info about the gamestate
	MoveList [][3]uint `json:"movelist"`
	MoveHistory [][3]uint `json:"movehistory"`
	Board [64]uint `json:"board"`
	State struct {
		WtoMove bool `json:"w_move"`
		MoveNo uint `json:"moveno"`
		Incheck bool `json:"in_check"`
		GameOver bool `json:"game_over"`
		HalfMoveNo uint `json:"half_move_no"`
		CastleRights [2]uint `json:"castle_rights"`
	} `json:"state"`
}

type Eval struct {
	Total float64 `json:"eval"`
	White float64 `json:"white"`
	Black float64 `json:"black"`

	
} 

type BotData struct { // Info of bot analysis of gamestate

	Level uint `json:"level"`



	BestMove [3]uint `json:"bestmove"`
	BestLine [][3]uint `json:"bestline"`
		
	Depth struct {
		Depth uint `json:"depth"`
		Nodes uint `json:"nodes"`
		Pruned uint `json:"pruned"`
		TT_hits uint `json:"tt_hits"`
		TT_success uint `json:"tt_success"`
	} `json:"depth"`
}

// ============================================================================

func CreateData(gh *GameHost) ChessData {
	// Sends ChessData to client

	moveList := gh.GameState.MoveList

	// simple move lists
	SimpleMoveList := ExportMoveList(moveList)
	SimplePrevMoveList := ExportMoveList(gh.GameState.History.PrevMoves)

	// create the data
	gamedata := ChessData{
		Message: "Chess Game!",

		GameState: GameStateData{
			MoveList: SimpleMoveList,
			MoveHistory: SimplePrevMoveList,
			Board: gh.GameState.Board.ServerBoard(),
			State: struct {
				WtoMove bool `json:"w_move"`
				MoveNo uint `json:"moveno"`
				Incheck bool `json:"in_check"`
				GameOver bool `json:"game_over"`
				HalfMoveNo uint `json:"half_move_no"`
				CastleRights [2]uint `json:"castle_rights"`
			}{
				WtoMove: gh.GameState.White_to_move,
				MoveNo: gh.GameState.MoveNo,
				Incheck: gh.GameState.InCheck,
				GameOver: gh.GameState.GameOver,
				HalfMoveNo: gh.GameState.HalfMoveNo,
				CastleRights: [2]uint{gh.GameState.WhiteCastle, gh.GameState.BlackCastle},	
			},
			
		},
	
		
	}
	return gamedata
}

// ============================================================================
// Helper Functions

// converts each moveNo in list a list of [special, startSq, endSq] elements - for sending to client
func ExportMoveList(ml []uint) [][3]uint {

	var SimpleMoveList [][3]uint

	for _, move := range ml {
		var simpleMove [3]uint
		simpleMove[0] = move >> 12
		simpleMove[1] = move & 0b111111
		simpleMove[2] = (move >> 6) & 0b111111
		SimpleMoveList = append(SimpleMoveList, simpleMove)
	}

	return SimpleMoveList
}