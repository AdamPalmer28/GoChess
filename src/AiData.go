package src

import (
	"chess/src/chess_bot"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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

// Get Chess data from the GameHost
func ChessAIEndpoints(router *mux.Router, gh *GameHost) {


	// ---------------------------------------------------------------------
	// undo request from client 
	router.HandleFunc("/EvalScoreData", func(w http.ResponseWriter, r *http.Request) {
		
		EvalScore := chess_bot.EvalScoreData(gh.GameState)

		jsonData, err := json.Marshal(EvalScore)
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


	router.HandleFunc("/BitBoardData", func(w http.ResponseWriter, r *http.Request) {

		var BitBoard BitboardData

		// Set the data
		BitBoard.WhitePawns = *gh.GameState.Board.WhitePawns

		jsonData, err := json.Marshal(BitBoard)
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
