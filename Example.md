# Example: Applied patterns in this template

## 1. Hexagonal Architecture

The project separates business rules from technical concerns.

- `domain`: core business entity and repository contract.
- `application`: services that orchestrate business operations.
- `infrastructure`: PostgreSQL implementation.
- `presentation`: Gin handlers, container and routing.

## 2. Repository Pattern

`ProductRepository` is defined as a port inside `domain/ports`.
The PostgreSQL adapter implements that contract in `infrastructure/repositories`.

## 3. Dependency Inversion

The application services do not know about `pgx` or SQL.
They only depend on the `ProductRepository` interface.

## 4. Constructor Injection + Container

Dependencies are wired through `presentation/container` and composed in `cmd/api/main.go`.
This keeps object creation centralized and testable.

## 5. TDD-oriented design

The application services are built around repository abstractions.
This makes them easy to test with a fake repository.

## Example flow

1. The HTTP handler receives a request.
2. The handler maps the request to a service input.
3. The service validates and orchestrates the operation.
4. The repository adapter persists data in PostgreSQL.
5. The handler returns the HTTP response envelope.
