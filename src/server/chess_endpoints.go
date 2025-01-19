package server

import (
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type clientData struct {
	// Define the structure of your 'move' object
	Move [3]uint `json:"move"`
}

// Get Chess data from the GameHost
func ChessGameEndpoints(router *mux.Router, gh *GameHost) {

	// ---------------------------------------------------------------------
	// normal client request for game data
	router.HandleFunc("/chessgame", func(w http.ResponseWriter, r *http.Request) {
		gh.PackageChessData(w, r)
	}).Methods("GET")

	// ---------------------------------------------------------------------
	// undo request from client 
	router.HandleFunc("/undo", func(w http.ResponseWriter, r *http.Request) {
		if gh.GameState.MoveNo == 1 {
			http.Error(w, "No moves to undo", http.StatusBadRequest)
			return
		}
		gh.GameState.Undo()
		gh.PackageChessData(w, r)
	})

	// ---------------------------------------------------------------------
	// reset request from client
	router.HandleFunc("/newgame", func(w http.ResponseWriter, r *http.Request) {
		gh = StartGameHost()
		gh.PackageChessData(w, r)
	})

	// ---------------------------------------------------------------------
	// move request from client
	router.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {

		// Decode JSON request
		var moveRequest clientData
		
		// print body or request
		
		err := json.NewDecoder(r.Body).Decode(&moveRequest)
		
		// Decode JSON request
		fmt.Print(moveRequest)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		start := moveRequest.Move[0]
		end := moveRequest.Move[1]
		special := moveRequest.Move[2]
		
		fmt.Printf("Received move request: %d %d %d\n", start, end, special)
		

		// make the movenumber
		move_num := special << 12 | end << 6 | start

		// check if the move is valid
		var final_move uint 
		for _, gs_move := range gh.GameState.MoveList {
			// check 
			mv_sq := gs_move & (special << 12 | 0b111111_111111)
			if mv_sq == move_num {
				final_move = gs_move
				break
			}
		}

		if final_move == 0 {
			// invalid move
			fmt.Println("Invalid move")
			http.Error(w, "Invalid move", http.StatusBadRequest)

		} else {// make the move
			gh.GameState.Make_move(final_move)
		}

		// gh.GameState.Board.Print()
		// gh.GameState.Board.WhitePawns.Print()
		gh.PackageChessData(w, r)
	
	})

	// ---------------------------------------------------------------------
}