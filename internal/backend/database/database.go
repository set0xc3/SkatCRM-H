// Database:
//
// 	[driver] test.db
// 		[entity] users
//
// 	[driver] crm.db: open, close
// 		[entity] clients
// 		[entity] products
//		[entity] calls

package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"SkatCRM-Tiny/internal/backend/database/entities"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db         *sql.DB
	clientsEnt *entities.ClientEntity
	// productsEnt *entities.ProductEntity
	// callsEnt    *entities.CallEntity
}

var (
	instance *Database
	once     sync.Once
)

func GetInstance() *Database {
	once.Do(func() {
		instance = &Database{}
	})
	return instance
}

func (db *Database) Init() error {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return fmt.Errorf("DB_URL environment variable is not set")
	}

	var err error
	db.db, err = sql.Open("sqlite3", dbURL)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.db.Ping(); err != nil {
		db.db.Close()
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	db.clientsEnt = entities.NewClientEntity(db.db, dbURL)
	// db.productsEnt = entities.NewProductEntity(db.db, dbURL)
	// db.callsEnt = entities.NewCallEntity(db.db, dbURL)

	log.Println("Database initialized successfully")
	return nil
}

func (db *Database) Close() error {
	if db.db != nil {
		return db.db.Close()
	}
	return nil
}

func (db *Database) GetClients() *entities.ClientEntity {
	return db.clientsEnt
}
