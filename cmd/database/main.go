package main

import (
	"SkatCRM-Tiny/internal/backend/database"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.Open()
}
