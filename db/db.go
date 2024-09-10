package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var Sqlite *sql.DB

func init() {
	var err error
	Sqlite, err = sql.Open("sqlite", "./app.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Optionally, ping to ensure the connection is valid
	if err = Sqlite.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}
