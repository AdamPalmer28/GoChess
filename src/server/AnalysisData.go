package server

import (
	"chess/src/chess_bot"
	"chess/src/chess_engine/board"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// ============================================================================
// Data structure

type BitboardData struct {
	WhitePawns board.Bitboard `json:"wp"`
	BlackPawns board.Bitboard `json:"bp"`
}

type analysisData struct {
	
	Eval chess_bot.EvalScore `json:"evalScore"`

	BB BitboardData `json:"bitboard"`

}

// ============================================================================
// Endpoint


// Get Chess data from the GameHost
func AnalysisEndpoints(router *mux.Router, gh *GameHost) {
	// Analysis data of current gamestate (not considering depth)
	
	// ---------------------------------------------------------------------
	router.HandleFunc("/analysis", func(w http.ResponseWriter, r *http.Request) {
		
		evalData := chess_bot.EvalScoreData(gh.GameState)

		data := analysisData{
			Eval: evalData,

			BB: BitboardData{
				WhitePawns: *gh.GameState.Board.WhitePawns,
				BlackPawns: *gh.GameState.Board.BlackPawns,
			},
		}


		// export data
		jsonData, err := json.Marshal(data)
		

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