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
}
