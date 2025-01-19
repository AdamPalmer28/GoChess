// main.go
package main

import (
	"chess/src/server"

	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)



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
	gh := server.StartGameHost()


	// listen for requests for game data
	server.ChessGameEndpoints(router, gh) 

	// listen for requests for AI data
	server.ChessAIEndpoints(router, gh)

	// listen for requests for analysis data
	server.AnalysisEndpoints(router, gh)
	
	

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")	
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}