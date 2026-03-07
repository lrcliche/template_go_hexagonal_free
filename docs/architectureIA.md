# Architecture IA Guide (Go Hexagonal Template)

## 1) Purpose of this document

This document defines the architectural rules that **developers and AI coding agents** must follow when changing this template.

Goals:
- Keep a strict **Hexagonal / Clean Architecture** boundary.
- Preserve testability and maintainability.
- Keep runtime wiring centralized.
- Keep HTTP behavior and error contracts consistent.

This guide must be treated as a governance document for changes in this repository and as a reusable baseline for similar Go backend projects.

---

## 2) Architecture overview

This project uses a layered hexagonal structure:

- **Domain**: business entities and business-facing contracts (ports).
- **Application**: use-case orchestration through services.
- **Infrastructure**: technical adapters (PostgreSQL, repository implementations).
- **Presentation**: HTTP transport, middleware, routing, response mapping, DI container.
- **Bootstrap (`cmd/api/main.go`)**: app startup and shutdown lifecycle.

The architecture check script (`scripts/check_architecture.sh`) enforces core boundaries.

---

## 3) Project structure

```text
application/
  config/
  services/

domain/
  entities/
  ports/

infrastructure/
  database/
  repositories/

presentation/
  container/
  handlers/
  routes/
  responses/
  errors/
  middleware/
  logging/

cmd/api/main.go

docs/
scripts/
tests/
migrations/
```

---

## 4) Responsibilities of each layer

### Domain
- Owns core business concepts (`entities`) and business contracts (`ports`).
- Must remain framework-agnostic.
- Exposes domain errors used by upper layers.

### Application
- Implements use-case orchestration in `application/services`.
- Coordinates domain operations and repository ports.
- Must not know infrastructure implementation details.

### Infrastructure
- Implements domain ports (for example, PostgreSQL repository adapters).
- Owns DB connection details and persistence queries.
- Must not move business rules out of domain/application into infrastructure.

### Presentation
- Owns HTTP concerns only: handlers, middleware, routes, response envelope, error mapping, request validation.
- Handlers map request/response and call services.
- Container wires runtime dependencies.

### Bootstrap (`cmd/api/main.go`)
- Configures logger.
- Loads config.
- Builds container.
- Creates HTTP server.
- Handles graceful shutdown.

---

## 5) Dependency direction rules

Dependency flow must point inward:

- `domain` → no dependency on `application`, `infrastructure`, or `presentation`.
- `application` → can depend on `domain` (`entities`, `ports`). Must not import `infrastructure`.
- `infrastructure` → can depend on `domain` ports/entities to implement adapters.
- `presentation` → can depend on `application` services and presentation-local packages.

Additional rules:
- Only `presentation/container` should perform runtime wiring of services and repositories.
- Production code outside the container must not instantiate services directly.
- Tests can instantiate services directly with fakes/mocks.

---

## 6) HTTP conventions

### Response envelope (required)

Success responses must follow:

```json
{
  "errors": [],
  "data": {}
}
```

Error responses must follow:

```json
{
  "errors": [
    {
      "code": "unique-code",
      "message": "Internal Error"
    }
  ]
}
```

Conventions:
- Use `presentation/responses` helpers for consistency.
- Keep transport validation in handlers (request parsing/binding).
- Keep business logic in services, not in handlers.

---

## 7) Logging conventions

- Use Go **standard library `log`** only.
- Prefer structured message style with stable prefixes (component + operation).
- Application and presentation may log operational errors; avoid leaking sensitive information.

---

## 8) Server bootstrap rules

`cmd/api/main.go` must stay minimal and focused on startup lifecycle:

1. Configure logger output/flags.
2. Load configuration.
3. Build DI container.
4. Build router.
5. Start HTTP server with timeouts.
6. Handle SIGINT/SIGTERM and graceful shutdown.

Do not place business logic, SQL, or endpoint logic in `main`.

---

## 9) AI coding rules

When AI agents (Codex, Cursor, Copilot, ChatGPT, etc.) modify this template, they must:

1. Not break hexagonal boundaries.
2. Never import infrastructure/presentation packages from domain.
3. Keep handlers thin; no business logic in handlers.
4. Keep business operations in application services.
5. Ensure repositories implement domain ports.
6. Centralize dependency wiring in `presentation/container`.
7. Avoid creating generic catch-all packages such as `shared`, `common`, `helpers`, `utils`, `core`, `base`, `internal`.
8. Use standard library logging.
9. Respect the HTTP response envelope contract.
10. Run architecture and test checks before finishing changes.

---

## 10) Guidelines for adding new features

For each new module/aggregate:

1. **Domain first**
   - Add or evolve entity and domain validations.
   - Add/update domain port interfaces and domain errors as needed.

2. **Application use cases**
   - Add/update service methods and input models.
   - Keep orchestration and business flow here.

3. **Infrastructure adapter**
   - Implement/extend repository adapter for the new port behavior.
   - Keep SQL/persistence details in infrastructure.

4. **Presentation**
   - Add/extend handler, route, request/response mapping.
   - Reuse `presentation/errors` and `presentation/responses` conventions.

5. **Container wiring**
   - Register new dependencies in `presentation/container` only.

6. **Validation**
   - Run tests.
   - Run `scripts/check_architecture.sh`.




## AI AGENT INSTRUCTIONS

These rules are mandatory for autonomous AI coding agents (Codex, Cursor, Copilot, ChatGPT) operating in this repository.

### 1) Hexagonal architecture enforcement

- Preserve strict layer boundaries: `domain`, `application`, `infrastructure`, `presentation`.
- Domain remains framework-agnostic and must not import HTTP, DB, logging, or transport concerns.
- Application layer coordinates use cases and ports only; it must not depend on concrete infrastructure implementations.
- Infrastructure layer implements adapters and repositories for application/domain contracts.
- Presentation layer owns HTTP concerns (request/response parsing, status codes, headers, DTO mapping).
- All dependencies must point inward toward business rules.

### 2) Generic utility folders are forbidden

- Do not create or use generic catch-all folders such as `utils`, `common`, `helpers`, `shared`, or `misc` at module or root level.
- Cross-cutting code must live in an explicit, purpose-driven package (for example: `logging`, `httpx`, `clock`, `uuid`, `errors`).
- New code must be placed by bounded context and architectural role, not by technical convenience.

### 3) HTTP response standards

- Handlers must return consistent JSON responses.
- Success responses:
  - `200 OK` for reads/updates.
  - `201 Created` for resource creation.
  - `204 No Content` only when no body is returned.
- Client errors:
  - `400 Bad Request` for malformed input.
  - `404 Not Found` for missing resources.
  - `409 Conflict` for business conflicts.
  - `422 Unprocessable Entity` for domain validation failures.
- Server errors:
  - `500 Internal Server Error` for unexpected failures.
- Error body format must be stable and include machine-readable code and human-readable message (example: `{ "error": { "code": "...", "message": "..." } }`).
- Never leak internal details (SQL, stack traces, infrastructure internals) in client responses.

### 4) Logging conventions

- Use structured logs with key-value fields.
- Minimum fields for every log entry: `timestamp`, `level`, `message`, `trace_id` (when available), and operation context.
- Log levels:
  - `DEBUG`: local diagnostics only.
  - `INFO`: normal lifecycle/business milestones.
  - `WARN`: recoverable anomalies.
  - `ERROR`: failed operations.
- Never log secrets or sensitive data (tokens, passwords, personal data beyond approved fields).
- HTTP handlers should log request lifecycle at `INFO` and failures at `WARN`/`ERROR` with correlation metadata.

### 5) Container dependency wiring

- Runtime object creation must be centralized in `presentation/container`.
- Use cases must receive dependencies through constructor injection.
- Handlers/routes must resolve dependencies from the container; they must not instantiate repositories/services directly.
- Infrastructure adapters must be wired through interfaces/ports, preserving testability and inversion of control.

### 6) Server bootstrap rules

- Server startup entrypoint must be minimal and deterministic.
- Bootstrap flow must follow this order:
  1. Load configuration.
  2. Initialize logger.
  3. Build container and wire dependencies.
  4. Register routes/middleware.
  5. Start HTTP server.
- Graceful shutdown handling is required (signal capture, context timeout, resource cleanup).
- Bootstrap code must not contain business logic.

### 7) Commit conventions for autonomous AI tools

- Each autonomous change must produce a focused commit with a clear scope.
- Commit message format:
  - `<type>(<scope>): <summary>`
  - Types: `feat`, `fix`, `refactor`, `test`, `docs`, `chore`.
- The commit body must briefly state:
  - Architectural impact (if any).
  - Files/layers affected.
  - Validation performed (tests/checks run).
- Do not mix unrelated architectural changes in a single commit.
