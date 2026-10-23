package main

import (
	"fmt"
	"io"
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

	// Make an HTTP request to the ping-pong service via Kubernetes internal DNS
	resp, err := http.Get("http://ping-pong-svc:2345/pings")
	pingPongs := "0"
	if err == nil {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			pingPongs = string(body)
		}
	}

	fmt.Fprintf(w, "%s.\nPing / Pongs: %s", string(logData), pingPongs)
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
