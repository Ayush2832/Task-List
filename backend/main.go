package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
}

var (
	tasks  []Task
	nextID = 1
	mu     sync.Mutex
)

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := os.Getenv("ALLOWED_ORIGIN")
		if origin == "" {
			origin = "*"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	if tasks == nil {
		json.NewEncoder(w).Encode([]Task{})
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Title == "" {
		http.Error(w, `{"error":"invalid title"}`, http.StatusBadRequest)
		return
	}

	mu.Lock()
	task := Task{
		ID:        nextID,
		Title:     input.Title,
		CreatedAt: time.Now().Format("2006-01-02 15:04"),
	}
	tasks = append(tasks, task)
	nextID++
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, `{"error":"task not found"}`, http.StatusNotFound)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
	// Load .env FIRST — before any os.Getenv calls
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Printf("ALLOWED_ORIGIN = %q", os.Getenv("ALLOWED_ORIGIN"))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", corsMiddleware(healthHandler))
	mux.HandleFunc("GET /api/tasks", corsMiddleware(getTasksHandler))
	mux.HandleFunc("POST /api/tasks", corsMiddleware(createTaskHandler))
	mux.HandleFunc("OPTIONS /api/tasks", corsMiddleware(healthHandler)) // preflight
	mux.HandleFunc("DELETE /api/tasks/{id}", corsMiddleware(deleteTaskHandler))
	mux.HandleFunc("OPTIONS /api/tasks/{id}", corsMiddleware(healthHandler)) // preflight

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
