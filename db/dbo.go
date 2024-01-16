package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

var dbo *sql.DB

func Connect() {
	var err error
	dbo, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	migrate(dbo)
}

func Disconnect() {
	err := dbo.Close()
	if err != nil {
		log.Fatalf("failed to close db connection: %v", err)
	}
	dbo = nil
}

func GetConnection() *Queries {
	return New(dbo)
}

func WithTX(ctx context.Context, fn func(context.Context, *Queries) error) error {
	queries := New(dbo)
	tx, err := dbo.Begin()
	if err != nil {
		return err
	}
	if err := fn(ctx, queries); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func migrate(db *sql.DB) {
	// setup database connection
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("failed to set dialect: %v", err)
	}
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("failed to migrate: %v", err)
		panic(err)
	}
}
