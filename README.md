# Books API CRUD

## Description

This is a simple CRUD API for books.

```
{
  "id": 1,
  "title": "The Lord of the Rings",
  "author": "J. R. R. Tolkien",
  "description": "The Lord of the Rings is an epic high fantasy novel written by English author and scholar J. R. R. Tolkien.",
}
```

## Used dependencies

- `Timescale DB` for database (can be normal postgres too)
- `pgx` for postgresql driver
- `sqlc` for generating queries and models
- `goose` for database migrations
- `echo` for routing
- `playground/validator` for validating request DTOs

## Installation

#### Install SQLC for generating queries and tables:

```sh
# install sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
# add go path
export PATH="$PATH:$(go env GOPATH)/bin"
# generate sqlc objects
sqlc generate
```

#### Environment variables (.env in root)

```sh
# database
DATABASE_URL=postgres://postgres:postgres@localhost:5432/books?sslmode=disable
PORT=8080
```

## Running

```sh
docker build -f deployment/Dockerfile -t books-api .
docker compose -f deployment/docker-compose.yml up
docker compose -f deployment/docker-compose.yml down
```

## Project structure

```sh
├── db
│   ├── migrations              # database migrations
│   ├── db.go                   # main database file (GENERATED)
│   ├── dbo.go                  # Connecting to DB
│   ├── queries.sql.go          # queries file (GENERATED)
│   ├── models.sql              # models file (GENERATED)
│   ├── queries.sql             # queries file
│   └── migrator.sql            # goose migrator file
├── deployment
│   ├── Dockerfile              # dockerfile for building api
│   ├── docker-compose.yml      # compose file for deploying db + api
├── api
│   ├── api.go                  # initialize api and routes
│   └── books.go                # books routes
├── handler
│   └── books.go                # book handlers
├── test
│   ├── suit_test.go            # setup and teardown for tests
│   ├── util_test.go            # util functions for tests
│   └── books_test.go           # book tests
├── main.go                     # main file
├── go.mod                      # go modules file
├── go.sum                      # go modules file
```

## Remaking the project

1. Delete all migrations from `db/migrations`, `api/books.go`, `handlers/books.go`
2. Add all migrations to `db/migrations` as `00001_init.sql`
3. Run `sqlc generate`
