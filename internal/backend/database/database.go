package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"SkatCRM-Tiny/internal/backend/database/entities"
)

type Service struct {
	db         *sql.DB
	clientRepo entities.ClientRepository
}

var (
	dbUrl = os.Getenv("DB_URL")
)

func Open() *Service {
	db, err := sql.Open("sqlite3", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	s := &Service{
		db:         db,
		clientRepo: *entities.NewClientRepository(db),
	}

	return s
}

func Close(s *Service) error {
	log.Printf("Disconnected from database: %s", dbUrl)
	return s.db.Close()
}
