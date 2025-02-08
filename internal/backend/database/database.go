// Database:
//
// 	[sqlite3] test.db
// 		[entity] users
//
// 	[sqlite3] crm.db
// 		[entity] clients
// 		[entity] products
//		[entity] calls

package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"SkatCRM-Tiny/internal/backend/database/entities"
)

type Database struct {
	clientEnt entities.ClientEntity
}

func Init() *Database {
	dbUrl := os.Getenv("DB_URL")

	db, err := sql.Open("sqlite3", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	ctx := &Database{
		clientEnt: *entities.NewClientEntity(db, dbUrl),
	}

	return ctx
}

func (ctx *Database) GetClientInstance() *entities.ClientEntity {
	return &ctx.clientEnt
}
