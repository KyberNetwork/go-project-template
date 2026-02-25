# Go Project Development Rules (Template-based)

This document defines recommended rules to develop Go services based on this template.

## 1) Project Structure Rules

- Keep executable entrypoints under `cmd/<service-name>/main.go`.
- Keep domain/business code under `internal/`.
- Reusable app bootstrapping utilities belong in `internal/app/`.
- DB helpers and migration utilities belong in `internal/dbutil/`.
- HTTP server setup belongs in `internal/server/`.
- Store-level interfaces and implementations belong in `internal/storage/`.
- SQL schema migrations must be in `migrations/` and versioned incrementally (`00001_*.up.sql`).

## 2) CLI and Configuration Rules

- Use `urfave/cli/v2` for service flags and runtime configuration.
- Read secrets/config from environment variables via CLI flags (`EnvVars`) rather than hardcoding values.
- Keep PostgreSQL options centralized through shared flags in `internal/app/postgresql.go`.
- Use `.env` loading only for local development convenience.

## 3) Logging and Observability Rules

- Use `zap` as the default logger; use sugared logger only where ergonomics are needed.
- Integrate Sentry via `sentry-go` + `zapsentry` for error reporting in production.
- Keep log flushing via a dedicated flush function and call it before process exit.
- Keep optional external log shipping (`cclog`) isolated and configurable.
- Register pprof endpoints under `/debug` for runtime profiling when applicable.

## 4) Database Rules

- Use PostgreSQL as the primary datastore (`lib/pq` driver + `sqlx`).
- Build DSN consistently with shared helper (`internal/dbutil.FormatDSN`).
- Apply migrations at startup in controlled environments using `golang-migrate`.
- Never mix schema migration SQL with app logic; keep all schema changes in migration files.
- Always wrap DB errors with context (`fmt.Errorf("...: %w", err)`).

## 5) Testing Rules

- Use `internal/testutil` helpers for ephemeral development/test databases.
- Run migrations in test setup so tests validate the real schema.
- Keep test DB lifecycle explicit: create, migrate, test, teardown.

## 6) Code Quality Rules

- Enforce linting with `golangci-lint` using repository config (`.golangci.yml`).
- Keep formatting deterministic by enabling `gofmt`, `gofumpt`, `goimports`, and `gci`.
- Prefer explicit error handling and wrapped errors over silent failures.
- Keep packages cohesive and avoid circular dependencies.
- Use constructor functions (`New...`) for components and dependency wiring.

## 7) Dependency Management Rules

- Keep direct dependencies minimal and intentional in `go.mod`.
- Avoid adding new frameworks unless current stack cannot satisfy requirements.
- Prefer stable, well-maintained libraries already aligned with this template.
- Regularly run `go mod tidy` and review indirect dependency growth.

## Preferred Dependencies (based on this template)

### Core Runtime

- `github.com/urfave/cli/v2`: standard CLI/bootstrap pattern.
- `go.uber.org/zap`: high-performance structured logging.

### HTTP/API

- `github.com/gin-gonic/gin`: HTTP router and middleware.
- `github.com/gin-contrib/pprof`: runtime profiling integration.

### Database

- `github.com/jmoiron/sqlx`: lightweight SQL access with struct mapping.
- `github.com/lib/pq`: PostgreSQL driver.
- `github.com/golang-migrate/migrate/v4`: migration management.

### Error/Monitoring

- `github.com/pkg/errors`: legacy stack/context compatibility where used.
- `github.com/getsentry/sentry-go`: error tracking.
- `github.com/TheZeroSlave/zapsentry`: zap-to-sentry bridge.

### Utilities

- `github.com/joho/godotenv`: local environment loading.
- `github.com/shopspring/decimal`: precise decimal arithmetic for financial values.

## Suggested Development Workflow

1. Copy template and rename module in `go.mod`.
2. Define service command in `cmd/<service-name>/main.go`.
3. Add config flags via `internal/app` package.
4. Implement storage and domain logic under `internal/`.
5. Add/update SQL migrations under `migrations/`.
6. Run lint and format checks before commit.
7. Add integration tests using `internal/testutil`.
8. Build Docker image and validate local compose setup.
