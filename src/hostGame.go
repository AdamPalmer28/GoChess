package src

import (
	"chess/src/chess_engine"
	"chess/src/chess_engine/gamestate"
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)
type GameHost struct {

	GameState *gamestate.GameState

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

// PackageChessData
func (gh *GameHost) PackageChessData (w http.ResponseWriter, r *http.Request) {
	// Convert the data to JSON
	gamedata := CreateData(gh)
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
}

// 

// Get Chess data from the GameHost
func GetChessData(router *mux.Router, gh *GameHost) {

	// normal client request for game data
	router.HandleFunc("/chessgame", gh.PackageChessData).Methods("GET")

	// -------------------------------------------------------------------------
	// move request from client
	router.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {

		// client sends move in the form of a list (start, end, promotion)
		// start and end are the square numbers

		move := r.URL.Query().Get("move")
		fmt.Println("Move received: ", move)

		gh.PackageChessData(w, r)

	}).Methods("GET")
}