package server

import (
	"chess/src/chess_engine"
	"chess/src/chess_engine/gamestate"

	"encoding/json"
	"net/http"
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


// Converts the Current GameHost data to JSON for outgoing server data
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

