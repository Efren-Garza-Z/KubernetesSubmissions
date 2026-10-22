package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"time"
)

var randomString string

// Generate the string once
func init() {
	b := make([]byte, 16)
	rand.Read(b)
	randomString = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// Keep printing to logs every 5 seconds
func printLogs() {
	for {
		fmt.Printf("%s: %s\n", time.Now().UTC().Format("2006-01-02T15:04:05.000Z"), randomString)
		time.Sleep(5 * time.Second)
	}
}

// Handle HTTP requests
func statusHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	fmt.Fprintf(w, "%s: %s", timestamp, randomString)
}

func main() {
	// Start the logging in the background
	go printLogs()

	// Start the HTTP server
	http.HandleFunc("/", statusHandler)
	port := "8080"
	fmt.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
