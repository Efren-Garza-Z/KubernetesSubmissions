package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var counter int
var mutex sync.Mutex

func pingpongHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	c := counter
	counter++
	// Write the current count to the shared volume
	os.WriteFile("/usr/src/app/shared/pingpongs.txt", []byte(fmt.Sprintf("%d", counter)), 0644)
	mutex.Unlock()
	fmt.Fprintf(w, "pong %d", c)
}

func main() {
	os.MkdirAll("/usr/src/app/shared", os.ModePerm)
	http.HandleFunc("/pingpong", pingpongHandler)
	port := "8080"
	fmt.Printf("Ping-pong server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
