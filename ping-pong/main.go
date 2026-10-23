package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var counter int
var mutex sync.Mutex

func pingpongHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	c := counter
	counter++
	mutex.Unlock()
	fmt.Fprintf(w, "pong %d", c)
}

func pingsHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	c := counter
	mutex.Unlock()
	fmt.Fprintf(w, "%d", c)
}

func main() {
	http.HandleFunc("/pingpong", pingpongHandler)
	http.HandleFunc("/pings", pingsHandler)
	port := "8080"
	fmt.Printf("Ping-pong server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
