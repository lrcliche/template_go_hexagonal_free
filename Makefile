APP_NAME := template-go-hexagonal
MAIN_PACKAGE := ./cmd/api
BUILD_DIR := ./bin
BINARY := $(BUILD_DIR)/api

.PHONY: help up run dev test test-cover fmt tidy lint-safe architecture-check build clean

help:
	@echo "Available commands:"
	@echo "  make up                 - bootstrap local dependencies"
	@echo "  make run                - run API once (no live reload)"
	@echo "  make dev                - run API with Air live reload"
	@echo "  make test               - run all tests"
	@echo "  make test-cover         - run tests with coverage report"
	@echo "  make fmt                - format all Go files"
	@echo "  make tidy               - tidy go.mod and go.sum"
	@echo "  make lint-safe          - run safe static checks (fmt/vet/test)"
	@echo "  make architecture-check - run hexagonal architecture guardrails"
	@echo "  make build              - build binary into ./bin"
	@echo "  make clean              - remove build artifacts"

up:
	@test -f .env || cp .env.example .env
	@echo "Local environment ready (.env created if missing)."

run:
	@./scripts/run.sh

dev:
	@./scripts/dev.sh

test:
	@./scripts/test.sh

test-cover:
	@./scripts/test.sh -coverprofile=coverage.out
	@go tool cover -func=coverage.out

fmt:
	@./scripts/fmt.sh

tidy:
	@go mod tidy

lint-safe:
	@./scripts/lint_safe.sh

architecture-check:
	@./scripts/check_architecture.sh

build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BINARY) $(MAIN_PACKAGE)

clean:
	@rm -rf $(BUILD_DIR) ./tmp coverage.out
