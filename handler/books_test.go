package handler_test

import (
	"books-api/handler"
	"books-api/util"
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	util.SetupIntegrationTest(m)
}

func TestGetBooksAndCreateWorks(t *testing.T) {
	util.BeforeEach()
	createBookDTO := handler.InsertBookDTO{
		Title:       "Test Book",
		Description: "Test Description",
		Author:      "Test Author",
	}

	req := util.BuildJsonRequest(http.MethodPost, "/books", createBookDTO)
	res, err := util.PerformRequest(req, handler.CreateBook)
	if err != nil {
		t.Fatal(err)
	}
	if res.Code != 201 {
		t.Fatalf("Expected 201, got %d", res.Code)
	}

	req = util.BuildJsonRequest(http.MethodGet, "/books", nil)
	res, err = util.PerformRequest(req, handler.GetBooks)
	if err != nil {
		t.Fatal(err)
	}
	if res.Code != 200 {
		t.Fatalf("Expected 200, got %d", res.Code)
	}

	var books handler.BooksDTO
	err = util.ParseBodyString(res.Body, &books)
	if err != nil {
		t.Fatal(err)
	}

	if len(books.Books) != 1 {
		t.Fatalf("Expected 1 book, got %d", len(books.Books))
	}

	if books.Books[0].Title != createBookDTO.Title {
		t.Fatalf("Expected title %s, got %s", createBookDTO.Title, books.Books[0].Title)
	}
	if books.Books[0].Description != createBookDTO.Description {
		t.Fatalf("Expected description %s, got %s", createBookDTO.Description, books.Books[0].Description)
	}
	if books.Books[0].Author != createBookDTO.Author {
		t.Fatalf("Expected author %s, got %s", createBookDTO.Author, books.Books[0].Author)
	}
}
