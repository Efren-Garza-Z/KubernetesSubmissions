package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	logData, err := os.ReadFile("/usr/src/app/files/log.txt")
	if err != nil {
		http.Error(w, "Log file not ready yet", http.StatusNotFound)
		return
	}

	pingPongs, err := os.ReadFile("/usr/src/app/shared/pingpongs.txt")
	if err != nil {
		pingPongs = []byte("0") // Default if ping-pong hasn't been accessed yet
	}

	fmt.Fprintf(w, "%s.\nPing / Pongs: %s", string(logData), string(pingPongs))
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
