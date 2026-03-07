# Add Feature IA Guide (Go Hexagonal Template)

## 1) Purpose

This document defines a strict, reusable contract for adding new HTTP features using AI coding tools (Codex, Cursor, Copilot, ChatGPT, and local agents) in this repository.

Goals:
- Protect hexagonal architecture boundaries.
- Standardize feature naming and file generation.
- Reduce accidental modifications to unrelated areas.
- Keep generated code maintainable and production-ready.

This guide must be used together with `docs/architectureIA.md`.

---

## 2) Feature naming rules

A feature name is mandatory input when generating a new module. It must follow these rules:

- Use lowercase only.
- Snake_case is allowed.
- Do not use spaces.
- Prefer singular nouns.
- Represent a business concept.

### Valid names

- `product`
- `category`
- `order`
- `payment`
- `product_review`
- `inventory_item`

### Invalid names

- `Products`
- `productService`
- `product-service`
- `ProductHandler`
- `utils`

---

## 3) Feature generation contract

When creating a new feature, AI must generate only the expected files for that feature.

For a feature named `category`, generate:

```text
domain/
  entities/category.go
  ports/category_repository.go

application/
  services/category_service.go

presentation/
  handlers/category_handler.go
  routes/category_routes.go
```

Generation rules:
- Generated files must include explicit `TODO` placeholders where business behavior will be implemented later.
- Do not generate extra folders or generic helper packages.
- Keep naming consistent with the selected feature name.

---

## 4) Container registration

After file generation, dependencies must be wired in:

- `presentation/container/container.go`

Expected pattern:
- `ResolveCategoryService()`
- `ResolveCategoryHandler()`

The container is the single place responsible for constructing services and handlers.

---

## 5) Route registration

Routes for the new feature must be registered in:

- `presentation/routes/router.go`

Expected pattern:
- `categoryRoutes.Register(router, container)`

---

## 6) Handler responsibilities

Handlers must remain thin transport adapters.

Allowed responsibilities:
- Read path parameters.
- Read query parameters.
- Bind request body.
- Validate request structure.
- Call application service.
- Return standardized HTTP response.

Forbidden:
- Business rules in handlers.
- Data persistence logic in handlers.

---

## 7) Service responsibilities

Application services are the use-case layer and must:
- Orchestrate business operations.
- Coordinate domain entities.
- Call repository ports.

Services must not depend on:
- HTTP.
- Gin.
- Database drivers.

---

## 8) AI safety rules

When generating a new feature, AI must:

1. Respect hexagonal architecture boundaries.
2. Avoid modifying unrelated files.
3. Avoid large refactors during feature scaffolding.
4. Avoid introducing generic packages such as:
   - `shared`
   - `utils`
   - `helpers`
   - `common`
   - `core`
   - `internal`
   - `base`
5. Avoid moving business logic to handlers.

---

## 9) Feature implementation checklist

After generation, the developer must complete:

- [ ] Implement repository in `infrastructure/repositories`.
- [ ] Implement database queries.
- [ ] Implement service business logic.
- [ ] Validate input in handler.
- [ ] Register routes in `presentation/routes/router.go`.
- [ ] Add unit tests.

---

## 10) Developer workflow

Typical workflow when adding a new feature:

1. Run the generator:

   ```bash
   ./scripts/add_feature.sh category
   ```

2. Implement repository:

   - `infrastructure/repositories/category_repository.go`

3. Implement service logic:

   - `application/services/category_service.go`

4. Register routes:

   - `presentation/routes/router.go`

---

## 11) AI-assisted workflow

Developers can use AI to complete implementation details after generation.

Example prompt:

> Implement the Category feature following the rules in architectureIA.md and AddFeatureIA.md.

AI-generated code must follow all architecture constraints defined in this project.
