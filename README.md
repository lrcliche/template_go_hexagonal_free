# Go Hexagonal Architecture Template

![Go Hexagonal Architecture](https://via.placeholder.com/1200x400)

Production-ready backend template built with:

* Go
* Gin
* PostgreSQL
* Hexagonal Architecture (Ports & Adapters)

Designed for developers who want **clean, scalable and maintainable backend systems from day one.**

---

# ⭐ If this project helps you, please star the repo

Buy the **PRO version** here:

👉 https://YOUR_GUMROAD_LINK

---

# Why this template?

Most Go backend projects start messy.

Controllers talk directly to databases.
Business logic spreads everywhere.
Scaling becomes painful.

This template provides a **clean architecture foundation** that keeps your backend organized and maintainable.

Key principles used:

* Hexagonal Architecture
* Dependency Injection
* Domain Driven Design principles
* Clean separation of layers
* Testable architecture
* AI-ready development workflow

---

# What is included in the FREE version

This repository contains the **basic architecture structure**.

Included:

✔ Project folder structure
✔ Example domain entity
✔ Example use case
✔ Infrastructure repository example
✔ Basic HTTP server with Gin
✔ Simple CRUD example
✔ Architecture documentation

This version is intended to **help you understand the architecture**.

---

# Project Structure

```
cmd/
internal/

domain/
entities/
ports/

application/
usecases/

infrastructure/
database/
repositories/

presentation/
http/
handlers/

config/
```

The project follows **Hexagonal Architecture (Ports & Adapters)**.

Dependencies always point inward.

---

# Architecture Flow

```
HTTP Request
    ↓
Controller (Presentation)
    ↓
Use Case (Application)
    ↓
Domain Entities
    ↓
Repository Port
    ↓
Database Adapter
```

This ensures that **business rules never depend on frameworks or databases**.

---

# Other Versions

The **Other versions** expand this template into a complete backend starter kit.

Includes:

✔ Feature generator (create modules automatically)
✔ Architecture validation rules
✔ Production-ready folder structure
✔ Authentication example
✔ Environment configuration
✔ More real-world modules
✔ Extended documentation

Get the Other versions here:

👉 https://lramos.gumroad.com/l/hexagonagoia

---

# Who is this template for?

This project is ideal for:

* Go backend developers
* Developers learning clean architecture
* Developers building scalable APIs
* Teams creating microservices
* Developers using AI coding assistants

---

# Getting Started

Clone the repository

```
git clone https://github.com/lrcliche/template_go_hexagonal_free
```

Run the project

```
go run cmd/main.go
```

---

# License

MIT License.

This repository contains the **free educational version**.

The **production-ready version is available here**:

👉  https://lramos.gumroad.com/l/hexagonagoia

---

# Support the project

If you find this project useful:

⭐ Star the repository
🍋 Buy the PRO version
💬 Share feedback

---

# Author

Luis Ramos

Senior Backend Developer

Specialized in:

* Go backend systems
* Hexagonal architecture
* Payment integrations
* Scalable backend platforms