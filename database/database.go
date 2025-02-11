package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"hugo.mardbrink.se/internal/config"
)

func InitDB() *sql.DB {
    db, err := sql.Open("sqlite3", config.DatabaseFile)
    if err != nil {
        log.Fatal(err)
    }

    if err := initSchema(db); err != nil {
        log.Fatal(err)
    }

    return db
}

func initSchema(db *sql.DB) error {
    content, err := os.ReadFile(config.SeedFile)
    if err != nil {
        return err
    }

    _, err = db.Exec(string(content))
    if err != nil {
        return err
    }

    return nil
}

func CloseDB(db *sql.DB) {
    db.Close()
}


