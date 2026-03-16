# Development Guide

Welcome to the **Go CoffeeShop** development guide! This document will help you set up your environment and understand the development workflow for this project.

## 🛠 Prerequisites

Before you begin, ensure you have the following installed:

- **Go:** 1.21 or later
- **Docker & Docker Compose:** For running infrastructure (Postgres, RabbitMQ) and services.
- **Protocol Buffers (protoc):** For generating gRPC code.
- **Buf:** A modern tool for working with Protocol Buffers.
- **Wire:** For compile-time dependency injection.
- **Sqlc:** For generating type-safe Go code from SQL.
- **Golangci-lint:** For linting the codebase.

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/thangchung/go-coffeeshop.git
cd go-coffeeshop
```

### 2. Set up Infrastructure

You can start the core infrastructure (PostgreSQL and RabbitMQ) using Docker Compose:

```bash
make docker-compose-core-start
```

This will start:
- **PostgreSQL:** Port 5432
- **RabbitMQ:** Port 5672 (Management UI: http://localhost:15672, guest/guest)

### 3. Run the Services

You can run all services using:

```bash
make docker-compose-start
```

Or run individual services for faster development:

```bash
# In separate terminals:
make run-product
make run-counter
make run-barista
make run-kitchen
make run-proxy
make run-web
```

The Web UI will be available at: http://localhost:8888

## 🏗 Development Workflow

### Adding/Modifying Proto Definitions

The project uses gRPC for service-to-service communication. Proto files are located in the `proto/` directory.

1.  Modify or add `.proto` files in `proto/`.
2.  Generate Go code using `buf`:
    ```bash
    buf generate
    ```

### Working with the Database (SQLC & Migrations)

We use `sqlc` for type-safe database access and `golang-migrate` for migrations.

1.  **Migrations:** Add or modify migrations in `db/migrations/`.
2.  **SQL Queries:** Add or modify `.sql` files within each service's `infras/postgresql/query/` directory (if applicable, check `sqlc.yaml`).
3.  **Generate Code:**
    ```bash
    make sqlc
    ```

### Dependency Injection (Wire)

The project uses `google/wire` for dependency injection. Each service has a `wire.go` file in its `internal/<service>/app/` directory.

1.  Add new dependencies to the `New` functions or providers.
2.  Update the `wire.go` file to include new providers.
3.  Generate the `wire_gen.go` file:
    ```bash
    make wire
    ```

### Adding a New Service

To add a new service (e.g., `auth`):

1.  Create the directory structure: `internal/auth/{app,domain,usecases,infras}`.
2.  Create the entry point: `cmd/auth/main.go`.
3.  Define the configuration: `cmd/auth/config/`.
4.  Add a Dockerfile: `docker/Dockerfile-auth`.
5.  Update `docker-compose.yaml`.
6.  Update `Makefile` to include the new service in `run-auth` and `wire` targets.

## 🧪 Testing

Run all tests using the Makefile:

```bash
make test
```

For specific service tests:

```bash
go test ./internal/counter/...
```

## 🧹 Linting

Ensure your code adheres to the project's standards:

```bash
make linter-golangci
```

## 📜 Coding Standards & Conventions

- **Clean Architecture:** Keep business logic in `domain/` and `usecases/`. Infrastructure details (DB, gRPC, MQ) belong in `infras/`.
- **Domain-Driven Design (DDD):** Use Aggregate Roots, Entities, and Value Objects.
- **Error Handling:** Use custom errors defined in the `domain/` layer.
- **Logging:** Use `golang.org/x/exp/slog`.
- **Formatting:** Use `go fmt` or your IDE's built-in formatter.

## ❓ Troubleshooting

- **Wire errors:** If you see "redeclared in this block", ensure your `wire.go` has the `//go:build wireinject` tag.
- **Database migrations:** If a migration fails, you might need to manually fix the `schema_migrations` table in Postgres.
- **Port conflicts:** Ensure ports 5432, 5672, 15672, 5000-5002, and 8888 are free.

---
*Happy Coding!* ☕️
