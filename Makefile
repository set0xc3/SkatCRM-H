.PHONY: all pre-build build run-cli run-server clean

all: build

pre-build:
	@mkdir -p build
	@tailwindcss -o internal/frontend/assets/css/tw.css --minify

build: pre-build
	@templ generate
	@go build -o build/cli cmd/cli/main.go
	@go build -o build/server cmd/server/main.go

run-cli: build
	@go run cmd/cli/main.go

run-server: build
	@go run cmd/server/main.go

clean:
	@rm -rf build
