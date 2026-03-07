# template_go_hexagonal_v1

Professional Go backend starter kit using hexagonal architecture, prepared for fast local onboarding and production-friendly evolution.

## Purpose

This template provides a clean foundation for REST APIs in Go with strict boundaries between domain, application, infrastructure, and presentation layers.

## Architecture Summary

- **domain**: entities and ports (business contracts)
- **application**: use cases/services orchestrating business rules
- **infrastructure**: adapters (database/repository implementations)
- **presentation**: HTTP handlers, middleware, responses, routing, DI container

The architecture guard script (`scripts/check_architecture.sh`) enforces these boundaries.

## Prerequisites

- Go 1.22+
- PostgreSQL (for runtime persistence)
- [Air](https://github.com/air-verse/air) for live reload in dev mode

Install Air:

```bash
go install github.com/air-verse/air@latest
```

## Environment Setup

1. Copy environment template:

```bash
make up
```

2. Adjust `.env` values if needed (especially `POSTGRES_DSN`).

## Local Development Flow

### Run in development mode (live reload)

```bash
make dev
```

Or directly:

```bash
air
```

### Run normally (no live reload)

```bash
make run
```

### Run tests

```bash
make test
```

### Run architecture checks

```bash
make architecture-check
```

### Run safe lint checks

```bash
make lint-safe
```

## Main Developer Commands

- `make help` — list all commands
- `make up` — bootstrap `.env`
- `make dev` — live reload with Air
- `make run` — run API once
- `make test` — run all tests
- `make test-cover` — test coverage summary
- `make fmt` — format Go code
- `make tidy` — tidy dependencies
- `make lint-safe` — gofmt/go vet/go test checks
- `make architecture-check` — hexagonal boundary checks
- `make build` — build binary in `./bin`
- `make clean` — remove generated artifacts

## Healthcheck

The API exposes:

- `GET /health`

Expected response:

```json
{
  "errors": [],
  "data": {
    "status": "ok"
  }
}
```
