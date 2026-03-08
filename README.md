# Go Hexagonal Architecture Template (Free Community Version)

This repository is the **free basic/community edition** of a Go backend template based on Hexagonal Architecture.

It is designed to:

- demonstrate architecture quality and folder organization
- provide a clean starter skeleton for study and experimentation
- showcase the coding style used in the premium version

> The full production-ready implementation is available in the paid Gumroad version.

## What this free version includes

- clean hexagonal folder structure
- semantic English package names
- one sample domain entity (`Product`)
- one sample repository port (`ProductRepository`)
- one sample application service skeleton (`ProductService`)
- one sample HTTP handler skeleton (`ProductHandler`)
- one sample router with healthcheck + demo products route
- `.env.example`
- architecture overview document (`docs/architecture_overview.md`)

## What is intentionally simplified in this free version

To protect premium value, this repository is intentionally **non-functional for core business flows**:

- no full CRUD implementation
- no real PostgreSQL repository logic
- no advanced dependency injection/container setup
- no architecture validation scripts
- no AI governance / AI workflow / feature generation docs
- no feature generation scripts/tooling
- no premium logging/tracing toolkit
- no advanced automated tests

Demo adapters return explicit "not implemented" style responses where applicable.

## Upgrade to Premium

The premium version includes:

- full CRUD API implementation
- PostgreSQL repository implementation
- architecture validation scripts
- AI architecture governance docs
- feature generators and automation scripts
- advanced developer tooling and tests

👉 Get the premium template on Gumroad (replace this line with your product URL).

## Quick Start

```bash
cp .env.example .env
go run ./cmd/api
```

Then open:

- `GET /health`
- `GET /api/v1/products` (demo endpoint; returns not implemented in free version)

## Project Structure

```text
cmd/api
application/
  config/
  services/
domain/
  entities/
  ports/
infrastructure/
  repositories/
presentation/
  container/
  errors/
  handlers/
  middleware/
  responses/
  routes/
  server/
docs/
```

## License

MIT
