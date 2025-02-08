package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"SkatCRM-Tiny/internal/backend/database/entities"
)

type Database struct {
	db        *sql.DB
	clientEnt entities.ClientEntity
}

func Open() *Database {
	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("sqlite3", dbUrl)

	if err != nil {
		log.Fatal(err)
	}

	ctx := &Database{
		db:        db,
		clientEnt: *entities.NewClientEntity(db),
	}

	return ctx
}

func Close(s *Database) error {
	return s.db.Close()
}
