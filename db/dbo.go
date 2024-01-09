package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

var DB *pgxpool.Pool

func Connect() {
	var err error
	DB, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	dbo := stdlib.OpenDBFromPool(DB)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	migrate(dbo)
}

func Disconnect() {
	DB.Close()
	DB = nil
}

func WithTX(ctx context.Context, fn func(context.Context, *Queries, pgx.Tx) error) error {
	queries := New()
	tx, err := DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	if err := fn(ctx, queries, tx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
