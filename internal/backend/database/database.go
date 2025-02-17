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

func (ctx *Database) Init() error {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return fmt.Errorf("DB_URL environment variable is not set")
	}

	var err error
	ctx.db, err = sql.Open("sqlite3", dbURL)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err := ctx.db.Ping(); err != nil {
		ctx.db.Close()
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	ctx.clientsEnt = entities.NewClientEntity(ctx.db, dbURL)
	// db.productsEnt = entities.NewProductEntity(db.db, dbURL)
	// db.callsEnt = entities.NewCallEntity(db.db, dbURL)

	log.Println("Database initialized successfully")
	return nil
}

func (ctx *Database) Close() error {
	if ctx.db != nil {
		return ctx.db.Close()
	}
	return nil
}

func (ctx *Database) GetClients() *entities.ClientEntity {
	return ctx.clientsEnt
}

func (ctx *Database) FetchMarks() ([]string, error) {
	query := `
      SELECT id, name
      FROM marks
  `
	rows, err := ctx.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var marks []string

	for rows.Next() {
		var id string
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		marks = append(marks, name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration: %w", err)
	}

	return marks, nil
}
