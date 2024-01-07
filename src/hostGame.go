package src

import (
	"chess/src/chess_engine"
	"chess/src/chess_engine/gamestate"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type GameHost struct {

	GameState *gamestate.GameState

	// 
}

type ChessData struct {
	Message string `json:"message"`
	MoveList [][3]uint `json:"movelist"`
	MoveHistory [][3]uint `json:"movehistory"`
	Board [64]uint `json:"board"`
	WtoMove bool `json:"w_move"`
}


func SendGameData(router *mux.Router, gh *GameHost) {
	// Sends ChessData to client

	moveList := gh.GameState.MoveList

	// simple move lists
	SimpleMoveList := ExportMoveList(moveList)
	SimplePrevMoveList := ExportMoveList(gh.GameState.History.PrevMoves)



	router.HandleFunc("/chessgame", func(w http.ResponseWriter, r *http.Request) {
		// Simulating data retrieval from a database or external source
		gamedata := ChessData{
			Message: "Chess Game!",
			MoveList: SimpleMoveList,
			MoveHistory: SimplePrevMoveList,
			WtoMove: gh.GameState.White_to_move,
			Board: gh.GameState.Board.ServerBoard(),
		}
	
		// Convert the data to JSON
		jsonData, err := json.Marshal(gamedata)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	
		// Set response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	
		// Send the JSON response
		w.Write(jsonData)
	})

}

// Server Host of the chess gamestate
func StartGameHost() *GameHost {

	gamestate.InitZobrist() // init zobrist keys

	// start the game
	gs := chess_engine.StartGame()
	
	gh := &GameHost{
		GameState: gs,
	}

	return gh
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