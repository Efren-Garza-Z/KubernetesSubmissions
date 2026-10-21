package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Read the PORT environment variable, default to 8080 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Simple handler so the server actually responds to requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Todo App is running")
	})

	// Log the exact message required by the exercise
	fmt.Printf("Server started in port %s\n", port)

	// Start the server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
