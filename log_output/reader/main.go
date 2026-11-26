package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	logData, err := os.ReadFile("/usr/src/app/files/log.txt")
	if err != nil {
		http.Error(w, "Log file not ready yet", http.StatusNotFound)
		return
	}

	infoData, err := os.ReadFile("/etc/config/information.txt")
	fileContent := ""
	if err == nil {
		fileContent = strings.TrimRight(string(infoData), "\r\n")
	}

	message := os.Getenv("MESSAGE")

	// Make an HTTP request to the ping-pong service via Kubernetes internal DNS
	resp, err := http.Get("http://ping-pong-svc:2345/pings")
	pingPongs := "0"
	if err == nil {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			pingPongs = strings.TrimSpace(string(body))
		}
	}

	fmt.Fprintf(
		w,
		"file content: %s\nenv variable: MESSAGE=%s\n%s.\nPing / Pongs: %s",
		fileContent,
		message,
		strings.TrimSpace(string(logData)),
		pingPongs,
	)
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
