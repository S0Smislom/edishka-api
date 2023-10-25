package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq" // ...
)

// initDB ...
func InitDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Init test db
func InitTestDB(dbURL string) (*sql.DB, error) {
	db, err := InitDB(dbURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}

func TeardownTestDB(db *sql.DB, tables ...string) error {
	if len(tables) > 0 {
		db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
	}

	db.Close()
	return nil
}
