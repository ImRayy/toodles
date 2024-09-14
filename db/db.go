package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"toodles/data/sql/migrations"

	"github.com/pressly/goose/v3"
	"github.com/pressly/goose/v3/database"
	_ "modernc.org/sqlite"
)

var Sqlite *sql.DB

func init() {
	log.SetFlags(0)
	ctx := context.Background()

	// User(s) home dir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := fmt.Sprintf("%s/.toodles/app.db", homeDir)

	var dbOpenErr error
	Sqlite, dbOpenErr = sql.Open("sqlite", dbPath)
	if dbOpenErr != nil {
		log.Fatal(dbOpenErr)
	}

	// Ensure the .toodles directory exists
	dbDir := fmt.Sprintf("%s/.toodles", homeDir)
	if err = os.MkdirAll(dbDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Create goose provider and up db
	provider, err := goose.NewProvider(database.DialectSQLite3, Sqlite, migrations.Embed)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = provider.Up(ctx); err != nil {
		log.Fatal(err)
	}
}
