package test

import (
	"books-api/db"
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	pool     *dockertest.Pool
	postgres *dockertest.Resource
}

func (suite *HandlerTestSuite) SetupSuite() {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}
	suite.pool = pool

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
	suite.postgres = resource
	time.Sleep(5 * time.Second)
	os.Setenv("DATABASE_URL", fmt.Sprintf("postgres://postgres:mysecretpassword@localhost:%s/foodie?sslmode=disable", resource.GetPort("5432/tcp")))
	db.Connect()
}

func (suite *HandlerTestSuite) TearDownSuite() {
	db.Disconnect()
	if err := suite.pool.Purge(suite.postgres); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func (suite *HandlerTestSuite) SetupTest() {
	err := db.WithTX(context.Background(), func(ctx context.Context, queries *db.Queries) error {
		err := queries.DeleteAllBooks(ctx)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		log.Fatalf("Could not delete all books: %s", err)
	}
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
