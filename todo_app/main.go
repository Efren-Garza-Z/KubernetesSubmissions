package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

var (
	imagePath  = getEnv("IMAGE_PATH", "/usr/src/app/files/image.jpg")
	imageURL   = getEnv("IMAGE_URL", "https://picsum.photos/1200")
	backendURL = getEnv("BACKEND_URL", "http://todo-backend-svc:2345/todos")
	ttlMinutes = getEnv("IMAGE_TTL_MINUTES", "10")
)

func downloadImage() error {
	resp, err := http.Get(imageURL)
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
	mins, err := strconv.Atoi(ttlMinutes)
	if err != nil {
		mins = 10
	}
	info, err := os.Stat(imagePath)
	if os.IsNotExist(err) || time.Since(info.ModTime()) > time.Duration(mins)*time.Minute {
		downloadImage()
	}
}

func main() {
	os.MkdirAll(filepath.Dir(imagePath), os.ModePerm)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		checkAndFetchImage()

		var todos []string
		resp, err := http.Get(backendURL)
		if err == nil {
			defer resp.Body.Close()
			json.NewDecoder(resp.Body).Decode(&todos)
		}

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

	port := getEnv("PORT", "8080")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
