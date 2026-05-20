# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project overview

An HTTP API service built with Gin (Go web framework) and GORM (ORM for MySQL). The README describes it as "AC搜索接口服务" (AC search interface service).

## Build & development

```bash
# Build
go build -o app main.go

# Run (requires .env with DB_HOST, DB_PORT, DB_DATABASE, DB_USERNAME, DB_PASSWORD)
go run main.go

# Update dependencies after changing imports
go mod tidy
```

No test files, linters, or Makefile exist in this project.

## Architecture

The project follows a layered convention:

| Directory | Purpose |
|---|---|
| `main.go` | Entry point: initializes databases, loads routes, starts server |
| `routers/` | Route registration. `Load()` returns `*gin.Engine`; each route is added there |
| `actions/` | HTTP handlers. Convention: filename = action name (snake_case), function name = action + HTTP method suffix |
| `services/` | Business logic packages, each in its own sub-folder |
| `settings/` | Global config loaded from `.env` via `godotenv`. Exports `Debug` bool. `RequireEnvs()` validates required env vars |
| `databases/` | DB init and connection management |
| `structs/models/` | Application domain models |
| `structs/requests/` | Shared request structs |

### Database layer (`databases/`)

```
databases/
├── init.go          # Blank imports for enabled database drivers (comment/uncomment to toggle)
├── mysql/mysql.go   # MySQL connection via GORM (reads DB_* env vars, auto-migrates)
├── entities/        # Database table definitions (GORM models)
├── logics/          # Entity constants + base CRUD logic, one sub-folder per entity
└── scopes/          # Reusable GORM query scopes
```

Database connections are initialized via `init()` functions (Go package init), triggered by blank imports in `databases/init.go`. Currently only MySQL is implemented; comments show planned support for MongoDB, ClickHouse, and Elasticsearch.

### Key conventions

- Settings/config are package-level variables (`settings.Debug`, `mysql.Db`, etc.)
- HTTP handlers use `gin.Context` directly (no service layer abstraction yet)
- Router is initialized in `init()` as a package variable in `routers/`
- File naming: snake_case; function naming: PascalCase
