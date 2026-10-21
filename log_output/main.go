package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

func main() {
	// Generate the random string once on startup
	b := make([]byte, 16)
	rand.Read(b)
	randomString := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	// Loop to print the string with a timestamp every 5 seconds
	for {
		fmt.Printf("%s: %s\n", time.Now().UTC().Format("2006-01-02T15:04:05.000Z"), randomString)
		time.Sleep(5 * time.Second)
	}
}
