package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// ... (keep downloadImage and checkAndFetchImage functions exactly the same) ...
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
	if os.IsNotExist(err) || time.Since(info.ModTime()) > 10*time.Minute {
		downloadImage()
	}
}

func main() {
	os.MkdirAll("/usr/src/app/files", os.ModePerm)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		checkAndFetchImage()

		// Fetch todos from backend
		var todos []string
		resp, err := http.Get("http://todo-backend-svc:2345/todos")
		if err == nil {
			defer resp.Body.Close()
			json.NewDecoder(resp.Body).Decode(&todos)
		}

		// Build the HTML list
		var todosHtml strings.Builder
		for _, todo := range todos {
			todosHtml.WriteString(fmt.Sprintf("<li>%s</li>", todo))
		}

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html>
			<head><title>Todo App</title></head>
			<body style="text-align: center; font-family: sans-serif;">
				<h1>Todo App</h1>
				<img src="/image.jpg" alt="Cached Image" style="max-width: 400px; height: auto; border-radius: 8px;"/>
				
				<div style="margin-top: 20px;">
					<!-- Form POSTs to /todos handled by Ingress -->
					<form action="/todos" method="POST">
						<input type="text" name="todo" maxlength="140" placeholder="Enter a new todo (max 140 characters)" style="width: 300px; padding: 5px;"/>
						<button type="submit" style="padding: 5px 10px;">Send</button>
					</form>
				</div>

				<div style="margin-top: 20px; text-align: left; display: inline-block; width: 400px;">
					<h3>Todos</h3>
					<ul>%s</ul>
				</div>
			</body>
			</html>
		`, todosHtml.String())
	})

	http.HandleFunc("/image.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, imagePath)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
