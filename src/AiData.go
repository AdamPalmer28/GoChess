package src

import (
	"chess/src/chess_bot"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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

}
