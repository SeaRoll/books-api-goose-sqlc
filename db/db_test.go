package db_test

import (
	"books-api/db"
	"books-api/util"
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestMain(m *testing.M) {
	util.SetupIntegrationTest(m)
}

func TestConnect(t *testing.T) {
	fmt.Println("TestConnect")
}

// Test get, create, update, delete
func TestGetBooksAndCreateWorks(t *testing.T) {
	util.BeforeEach()
	books := []db.Book{}
	ctx := context.Background()
	err := db.WithTX(ctx, func(txCtx context.Context, q *db.Queries, tx pgx.Tx) error {
		var err error

		// create book
		_, err = q.InsertBook(txCtx, tx, db.InsertBookParams{
			Title:       "Test Book",
			Description: "Test Description",
			Author:      "Test Author",
		})
		if err != nil {
			return err
		}

		// get books
		books, err = q.ListBooksPaged(txCtx, tx, db.ListBooksPagedParams{
			Limit:  10,
			Offset: 0,
		})
		return err
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(books) != 1 {
		t.Fatalf("Expected 1 book, got %d", len(books))
	}

	book := books[0]
	if book.ID == 0 {
		t.Fatal("Expected book ID to be set")
	}
	if book.Title != "Test Book" {
		t.Fatalf("Expected title to be 'Test Book', got '%s'", book.Title)
	}

	if book.Description != "Test Description" {
		t.Fatalf("Expected description to be 'Test Description', got '%s'", book.Description)
	}
	if book.Author != "Test Author" {
		t.Fatalf("Expected author to be 'Test Author', got '%s'", book.Author)
	}

	fmt.Println("TestGetBooks")
}

func TestUpdateBook(t *testing.T) {
	util.BeforeEach()
	books := []db.Book{}
	ctx := context.Background()
	err := db.WithTX(ctx, func(txCtx context.Context, q *db.Queries, tx pgx.Tx) error {
		var err error

		// create book
		id, err := q.InsertBook(txCtx, tx, db.InsertBookParams{
			Title:       "Test Book",
			Description: "Test Description",
			Author:      "Test Author",
		})
		if err != nil {
			return err
		}

		// update book
		err = q.UpdateBook(txCtx, tx, db.UpdateBookParams{
			ID:          id,
			Title:       "Test Book 2",
			Description: "Test Description 2",
			Author:      "Test Author 2",
		})
		if err != nil {
			return err
		}

		// get books
		books, err = q.ListBooksPaged(txCtx, tx, db.ListBooksPagedParams{
			Limit:  10,
			Offset: 0,
		})
		return err
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(books) != 1 {
		t.Fatalf("Expected 1 book, got %d", len(books))
	}

	book := books[0]
	if book.ID == 0 {
		t.Fatal("Expected book ID to be set")
	}
	if book.Title != "Test Book 2" {
		t.Fatalf("Expected title to be 'Test Book 2', got '%s'", book.Title)
	}
	if book.Description != "Test Description 2" {
		t.Fatalf("Expected description to be 'Test Description 2', got '%s'", book.Description)
	}
	if book.Author != "Test Author 2" {
		t.Fatalf("Expected author to be 'Test Author 2', got '%s'", book.Author)
	}
}

func TestDeleteBook(t *testing.T) {
	util.BeforeEach()
	books := []db.Book{}
	ctx := context.Background()
	err := db.WithTX(ctx, func(txCtx context.Context, q *db.Queries, tx pgx.Tx) error {
		var err error

		// create book
		id, err := q.InsertBook(txCtx, tx, db.InsertBookParams{
			Title:       "Test Book",
			Description: "Test Description",
			Author:      "Test Author",
		})
		if err != nil {
			return err
		}

		// delete book
		err = q.DeleteBook(txCtx, tx, id)
		if err != nil {
			return err
		}

		// get books
		books, err = q.ListBooksPaged(txCtx, tx, db.ListBooksPagedParams{
			Limit:  10,
			Offset: 0,
		})

		return err
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(books) != 0 {
		t.Fatalf("Expected 0 books, got %d", len(books))
	}
}
