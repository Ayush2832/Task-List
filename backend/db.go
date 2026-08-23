package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func connectDB() {
	connStr := os.Getenv("DATABASE_URL")

	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error in making sql connection", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			user_id VARCHAR(255) NOT NULL,
			title TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
	log.Println("Database migrated")
}
