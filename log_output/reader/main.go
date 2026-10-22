package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Read from the shared file
	data, err := os.ReadFile("/usr/src/app/files/log.txt")
	if err != nil {
		http.Error(w, "Log file not ready yet", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s", string(data))
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	fmt.Printf("Reader server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
