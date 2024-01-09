package handler_test

import (
	"books-api/handler"
	"books-api/util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
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
	req := httptest.NewRequest(http.MethodPost, "/", util.EncodeJSON(createBookDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res, err := util.PerformRequest(req, handler.CreateBook)
	if err != nil {
		t.Fatal(err)
	}
	if res.Code != 200 {
		t.Fatalf("Expected 200, got %d", res.Code)
	}
}
