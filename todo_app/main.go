package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const imagePath = "/usr/src/app/files/image.jpg"

func downloadImage() error {
	resp, err := http.Get("https://picsum.photos/1200")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func checkAndFetchImage() {
	info, err := os.Stat(imagePath)
	// If file doesn't exist or is older than 10 minutes, download a new one
	if os.IsNotExist(err) || time.Since(info.ModTime()) > 10*time.Minute {
		err := downloadImage()
		if err != nil {
			log.Printf("Error downloading image: %v", err)
		}
	}
}

func main() {
	// Ensure the shared directory exists
	os.MkdirAll("/usr/src/app/files", os.ModePerm)

	// Route for the main HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		checkAndFetchImage()

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html>
			<head><title>Todo App</title></head>
			<body style="text-align: center; font-family: sans-serif;">
				<h1>Todo App</h1>
				<img src="/image.jpg" alt="Cached Image" style="max-width: 400px; height: auto; border-radius: 8px;"/>
				
				<div style="margin-top: 20px;">
					<form action="#" method="POST">
						<input type="text" name="todo" maxlength="140" placeholder="Enter a new todo (max 140 characters)" style="width: 300px; padding: 5px;"/>
						<button type="submit" style="padding: 5px 10px;">Send</button>
					</form>
				</div>

				<div style="margin-top: 20px; text-align: left; display: inline-block; width: 400px;">
					<h3>Todos</h3>
					<ul>
						<li>Learn Kubernetes basics</li>
						<li>Deploy application to cluster</li>
						<li>Configure persistent volumes</li>
					</ul>
				</div>
			</body>
			</html>
		`)
	})

	// Route to serve the actual image file
	http.HandleFunc("/image.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, imagePath)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Todo Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
