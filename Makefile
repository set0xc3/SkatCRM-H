include .env

.PHONY: all tailwind-build migrations pre-build build-all run-cli run-server clean

all: build

tailwind-build:
	@tailwindcss -o internal/frontend/assets/css/tw.css --minify

migrations:
	@goose -dir ./migrations/ sqlite3 ${DB_URL} up

pre-build:
	@mkdir -p build
	@tailwindcss -o internal/frontend/assets/css/tw.css --minify
	@templ generate

build-all: pre-build
	@go build -o build/cli cmd/cli/main.go
	@go build -o build/server cmd/server/main.go

run-cli: pre-build
	@mkdir -p build
	@templ generate
	@go run cmd/cli/main.go

run-server:
	@mkdir -p build
	@templ generate
	@go run cmd/server/main.go

clean:
	@rm -rf build
