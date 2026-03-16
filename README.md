# web

Go REST API built with [Gin](https://github.com/gin-gonic/gin) and PostgreSQL.

## Tech Stack

- **Go** 1.21+
- **Gin** — HTTP router
- **PostgreSQL** — database
- **lib/pq** — Postgres driver

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/users` | List all users |
| GET | `/posts` | List all posts |
| GET | `/comments` | List all comments |
| GET | `/posts/:postId/comments` | Comments for a specific post |

## Setup

### 1. Clone & configure

```bash
cp .env.example .env
# Edit .env with your database credentials
```

### 2. Run

```bash
go run main.go
```

Server starts on `:8080` by default.

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | `localhost` | PostgreSQL host |
| `DB_PORT` | `5432` | PostgreSQL port |
| `DB_USER` | `postgres` | Database user |
| `DB_PASSWORD` | _(empty)_ | Database password |
| `DB_NAME` | `postgres` | Database name |
| `DB_SSLMODE` | `disable` | SSL mode (`disable` / `require`) |

## Project Structure

```
.
├── main.go          # Entry point, Gin setup, routes
├── services/        # Request handlers (GetUsers, GetPosts, ...)
├── repo/            # Data access layer
├── util/            # Utilities
├── data/            # Seed data / migrations
└── go.mod
```
