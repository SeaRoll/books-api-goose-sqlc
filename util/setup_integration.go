package util

import (
	"books-api/db"
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ory/dockertest/v3"
)

func SetupIntegrationTest(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("timescale/timescaledb", "latest-pg15", []string{"POSTGRES_DB=foodie", "POSTGRES_USER=postgres", "POSTGRES_PASSWORD=mysecretpassword"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	time.Sleep(5 * time.Second)
	os.Setenv("DATABASE_URL", fmt.Sprintf("postgres://postgres:mysecretpassword@localhost:%s/foodie?sslmode=disable", resource.GetPort("5432/tcp")))
	db.Connect()
	defer db.Disconnect()

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func BeforeEach() {
	err := db.WithTX(context.Background(), func(ctx context.Context, queries *db.Queries, tx pgx.Tx) error {
		err := queries.DeleteAllBooks(ctx, tx)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		log.Fatalf("Could not delete all books: %s", err)
	}
}
