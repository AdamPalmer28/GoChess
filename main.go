// main.go
package main

import (
	"chess/src"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type ChessData struct {
	Message string `json:"message"`
	MoveList [][3]uint `json:"movelist"`
}

func main() {
	router := mux.NewRouter()

	// Enable CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}), // Adjust this based on your requirements
	)
	// Use CORS middleware for all routes
	router.Use(corsHandler)


	// -------------------------------------------------------------------------

	// start game host
	gh := src.StartGameHost()

	moveList := gh.GameState.MoveList
	// convert []uint to []string for json
	var SimpleMoveList [][3]uint // [special, startSq, endSq]
	for _, move := range moveList {
		var simpleMove [3]uint
		simpleMove[0] = move >> 12
		simpleMove[1] = move & 0b111111
		simpleMove[2] = (move >> 6) & 0b111111
		SimpleMoveList = append(SimpleMoveList, simpleMove)
	}

	router.HandleFunc("/chessgame", func(w http.ResponseWriter, r *http.Request) {
		// Simulating data retrieval from a database or external source
		myData := ChessData{
			Message: "Chess Game!",
			MoveList: SimpleMoveList,
		}
	
		// Convert the data to JSON
		jsonData, err := json.Marshal(myData)
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
	

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")	
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}