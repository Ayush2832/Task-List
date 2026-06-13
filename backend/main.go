package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := os.Getenv("ALLOWED_ORIGIN")
		if origin == "" {
			origin = "*"
		}

		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-User-ID")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("X-User-ID")

	if userID == "" {
		http.Error(w, "missing user id", http.StatusBadRequest)
		return
	}

	rows, err := db.Query(`
		SELECT id, title, created_at
		FROM tasks
		WHERE user_id = $1
		ORDER BY id DESC
	`, userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	tasks := make([]Task, 0)

	for rows.Next() {
		var task Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.CreatedAt,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("X-User-ID")

	if userID == "" {
		http.Error(w, "missing user id", http.StatusBadRequest)
		return
	}

	var input struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if input.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	var task Task

	err := db.QueryRow(`
		INSERT INTO tasks(user_id, title)
		VALUES($1, $2)
		RETURNING id, title, created_at
	`,
		userID,
		input.Title,
	).Scan(
		&task.ID,
		&task.Title,
		&task.CreatedAt,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(task)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("X-User-ID")

	if userID == "" {
		http.Error(w, "missing user id", http.StatusBadRequest)
		return
	}

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(
		"DELETE FROM tasks WHERE id = $1 AND user_id = $2",
		id,
		userID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment")
	}

	connectDB()
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", corsMiddleware(healthHandler))

	mux.HandleFunc("GET /api/tasks", corsMiddleware(getTasksHandler))
	mux.HandleFunc("POST /api/tasks", corsMiddleware(createTaskHandler))
	mux.HandleFunc("DELETE /api/tasks/{id}", corsMiddleware(deleteTaskHandler))

	log.Printf("Server starting on port %s", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
