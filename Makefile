TEMPL_VERSION ?= latest

.PHONY: all build run-cli run-server clean

all: build

templ-install:
	@command -v templ >/dev/null || (echo "Installing templ..." && \
		go install github.com/a-h/templ/cmd/templ@$(TEMPL_VERSION))

build: templ-install
	@mkdir -p build
	@templ generate
	@go build -o build/cli cmd/cli/main.go
	@go build -o build/server cmd/server/main.go

run-cli:
	@go run cmd/cli/main.go

run-server:
	@go run cmd/server/main.go

clean:
	@rm -rf build
