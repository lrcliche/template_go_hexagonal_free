# Main Entrypoint Placement Audit

## Scope
This audit evaluates whether the application entrypoint should live at repository root (`main.go`) or under `cmd/api/main.go` for this hexagonal backend template.

## Option A: `main.go` at repository root

### Pros
- Very simple for tiny, single-binary services.
- Slightly shorter path for `go run` commands.

### Cons
- Scales poorly when introducing additional binaries (`worker`, `migrate`, `seed`, `cli`).
- Creates top-level clutter as operational binaries grow.
- Looks less standardized for commercial-quality backend templates.

## Option B: `cmd/api/main.go`

### Pros
- Explicitly models executable boundaries and keeps app packages focused.
- Scales cleanly to additional binaries: `cmd/worker`, `cmd/migrate`, `cmd/seed`, `cmd/cli`.
- Aligns with common Go production template conventions.
- Keeps repository root premium and organized for template consumers.

### Cons
- Slightly more nesting for very small projects.

## Current Project Assessment
- The current `cmd/api/main.go` contains only bootstrap responsibilities: config loading, dependency container wiring, HTTP server lifecycle, and graceful shutdown.
- Dependency wiring is centralized through `presentation/container`, preserving hexagonal boundaries.
- No business logic is located in the entrypoint.

## Final Decision
**Use `cmd/api/main.go` (Option B).**

This repository is positioned as a professional reusable backend template, not a throwaway single-binary toy app. The `cmd/api` structure provides better long-term maintainability, extensibility, and a cleaner commercial presentation while keeping architecture responsibilities explicit.

## Operational Note
No file move is required because the project already uses the recommended structure.
