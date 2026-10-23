package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var todos = []string{"Learn Kubernetes basics", "Deploy application to cluster", "Configure persistent volumes"}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		newTodo := r.FormValue("todo")
		if newTodo != "" {
			todos = append(todos, newTodo)
		}
		// Redirect the browser back to the main page after form submission
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	http.HandleFunc("/todos", todosHandler)
	port := "8080"
	fmt.Printf("Todo Backend starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
