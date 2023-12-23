// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Data struct {
	Message string `json:"message"`
    Timestamp time.Time `json:"timestamp"`
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

	router.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		// Simulating data retrieval from a database or external source
		myData := Data{
			Message: "Hello from the server!",
			Timestamp: time.Now(),
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