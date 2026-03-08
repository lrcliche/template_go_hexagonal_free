# Architecture Overview (Free Community Version)

This repository demonstrates a **Go Hexagonal Architecture skeleton** with a clean package layout:

- `domain`: core entities and ports
- `application`: service/use case contracts and orchestration skeletons
- `infrastructure`: adapter placeholders (non-functional in free version)
- `presentation`: HTTP handlers, middleware, and router wiring

## Intent of this free version

The goal is to show project structure and layering principles.

Runtime behavior is intentionally limited:

- repository methods return `not implemented` for premium-only flows
- only a small demo endpoint is exposed
- no production persistence or full CRUD flow is included

For the complete production-ready implementation, use the premium template.
