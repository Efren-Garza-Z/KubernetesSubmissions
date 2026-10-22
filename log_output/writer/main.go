package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"
)

func main() {
	b := make([]byte, 16)
	rand.Read(b)
	randomString := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	// Ensure the shared directory exists
	os.MkdirAll("/usr/src/app/files", os.ModePerm)

	for {
		timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
		logLine := fmt.Sprintf("%s: %s\n", timestamp, randomString)

		// Write to the shared file
		os.WriteFile("/usr/src/app/files/log.txt", []byte(logLine), 0644)

		time.Sleep(5 * time.Second)
	}
}
