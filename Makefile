APP_NAME := template-go-hexagonal
MAIN_PACKAGE := ./cmd/api
BUILD_DIR := ./bin
BINARY := $(BUILD_DIR)/api

.PHONY: help up run test fmt tidy build clean

help:
	@echo "Available commands:"
	@echo "  make up    - bootstrap local .env"
	@echo "  make run   - run API"
	@echo "  make test  - run tests"
	@echo "  make fmt   - format all Go files"
	@echo "  make tidy  - tidy go.mod and go.sum"
	@echo "  make build - build binary into ./bin"
	@echo "  make clean - remove build artifacts"

up:
	@test -f .env || cp .env.example .env
	@echo "Local environment ready (.env created if missing)."

run:
	@go run $(MAIN_PACKAGE)

test:
	@go test ./...

fmt:
	@go fmt ./...

tidy:
	@go mod tidy

build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BINARY) $(MAIN_PACKAGE)

clean:
	@rm -rf $(BUILD_DIR) coverage.out
