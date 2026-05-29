# backend

Go REST API for [curatedli.st](https://curatedli.st) — collaborative curated lists. Built with Gin, MySQL (`go-sqlbuilder`), and [Magic](https://magic.link) passwordless auth. Deployed to Google Cloud Run.

## Requirements

- Go 1.13+
- A MySQL 8 instance (schema in [`../init.sql`](../init.sql))

## Run locally

The simplest path is the top-level compose stack (MySQL + backend + frontend), run from the parent directory:

```bash
docker-compose up --build      # backend on :8080, db on :3306
```

To run the API directly against a local MySQL:

```bash
go run cmd/backend/main.go      # listens on :8080
```

Configuration is read from environment variables (defaults in `internal/config/config.go`):

| Variable        | Default                  | Notes                                   |
|-----------------|--------------------------|-----------------------------------------|
| `DATABASE_USER` | `test`                   |                                         |
| `DATABASE_PASS` | `test`                   |                                         |
| `DATABASE_URL`  | `tcp(localhost:3306)`    | DSN host part; prod uses a `unix(...)` Cloud SQL socket |
| `DATABASE_NAME` | `curatedlist_test`       |                                         |

## Build & test

```bash
go build -v ./cmd/backend/main.go
go test -v ./...                # all packages
go test -v ./internal/config    # a single package
```

## Architecture

`cmd/backend/main.go` is the composition root: it wires each domain's `Repository → Service → API` by hand and starts the Gin server. Each domain under `internal/` (`user`, `list`) is a vertical slice with the same layers:

- `repository.go` — SQL via `go-sqlbuilder`, operating on `Aggregate` structs (DB row shape, `sql.NullXxx` + `db:` tags)
- `service.go` — business logic; converts `Aggregate → DTO`
- `api.go` — Gin handlers: bind request bodies into `commands/`, enforce auth/ownership, emit JSON
- `dto.go` / `aggregate.go` — JSON shape vs. DB shape, with `ToXxx()` converters

Auth (`internal/middleware/auth.go`) validates the `Authorization: Bearer <token>` Magic token and stores the issuer (`iss`) on the request context; handlers resolve `iss → user`. Deletes are soft (a `deleted` flag). Item creation scrapes an `og:image`/`twitter:image` for the item picture (`internal/item/metabolize.go`).

## API routes

Public: `GET /status`, `GET /lists/`, `GET /lists/id/:id`, `GET /users/username/:username[/lists|/favs]`.
Authenticated (Bearer token): `POST/DELETE /lists/...`, list item + fav endpoints, and `POST/PUT /users/...`.

## Deploy

Pushing to `master` triggers `.github/workflows/go.yml`: it runs `go test ./...`, then `gcloud builds submit` + `gcloud run deploy back` to Cloud Run (project from the `GOOGLE_CLOUD_RUN_PROJECT` secret, region `europe-west1`). Production connects to Cloud SQL over a unix socket.

Manual deploy:

```bash
gcloud builds submit --tag gcr.io/curatedlist-project/back
gcloud run deploy back --image gcr.io/curatedlist-project/back \
  --region europe-west1 --platform managed --allow-unauthenticated
```
