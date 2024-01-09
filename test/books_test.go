package test

import (
	"books-api/handler"
	"net/http"
	"strconv"

	"github.com/stretchr/testify/assert"
)

func (suite *HandlerTestSuite) TestCreateAndGetBook() {
	expectedBooks := []handler.BookDTO{
		{
			Title:       "Test Book",
			Description: "Test Description",
			Author:      "Test Author",
		},
	}
	res, err := PerformRequest(BuildJsonRequest(http.MethodPost, "/books", handler.InsertBookDTO{
		Title:       "Test Book",
		Description: "Test Description",
		Author:      "Test Author",
	}), handler.CreateBook, nil)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 201, res.Code)
	assert.Equal(suite.T(), handler.SuccessDTO{Success: true}, ParseBodyString[handler.SuccessDTO](res.Body))

	res, err = PerformRequest(BuildJsonRequest(http.MethodGet, "/books", nil), handler.GetBooks, nil)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, res.Code)
	books := ParseBodyString[handler.BooksDTO](res.Body)
	assert.Equal(suite.T(), len(expectedBooks), len(books.Books))
	for i := 0; i < len(expectedBooks); i++ {
		assert.True(suite.T(), books.Books[i].ID > 0)
		assert.Equal(suite.T(), expectedBooks[i].Title, books.Books[i].Title)
		assert.Equal(suite.T(), expectedBooks[i].Description, books.Books[i].Description)
		assert.Equal(suite.T(), expectedBooks[i].Author, books.Books[i].Author)
	}
}

func (suite *HandlerTestSuite) TestCreateBookWithInvalidData() {
	res, err := PerformRequest(BuildJsonRequest(http.MethodPost, "/books", handler.InsertBookDTO{
		Title:       "",
		Description: "Test Description",
		Author:      "Test Author",
	}), handler.CreateBook, nil)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 400, res.Code)
}

func (suite *HandlerTestSuite) TestUpdateBook() {
	expectedBooks := []handler.BookDTO{
		{
			Title:       "Updated Title",
			Description: "Updated Description",
			Author:      "Updated Author",
		},
	}

	res, err := PerformRequest(BuildJsonRequest(http.MethodPost, "/books", handler.InsertBookDTO{
		Title:       "Test Book",
		Description: "Test Description",
		Author:      "Test Author",
	}), handler.CreateBook, nil)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 201, res.Code)

	res, err = PerformRequest(BuildJsonRequest(http.MethodGet, "/books", nil), handler.GetBooks, nil)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, res.Code)
	books := ParseBodyString[handler.BooksDTO](res.Body)

	res, err = PerformRequest(BuildJsonRequest(http.MethodPut, "/books", handler.UpdateBookDTO{
		ID:          books.Books[0].ID,
		Title:       "Updated Title",
		Description: "Updated Description",
		Author:      "Updated Author",
	}), handler.UpdateBook, nil)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, res.Code)
	assert.Equal(suite.T(), handler.SuccessDTO{Success: true}, ParseBodyString[handler.SuccessDTO](res.Body))

	res, err = PerformRequest(BuildJsonRequest(http.MethodGet, "/books", nil), handler.GetBooks, nil)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, res.Code)
	books = ParseBodyString[handler.BooksDTO](res.Body)
	assert.Equal(suite.T(), len(expectedBooks), len(books.Books))
	for i := 0; i < len(expectedBooks); i++ {
		assert.True(suite.T(), books.Books[i].ID > 0)
		assert.Equal(suite.T(), expectedBooks[i].Title, books.Books[i].Title)
		assert.Equal(suite.T(), expectedBooks[i].Description, books.Books[i].Description)
		assert.Equal(suite.T(), expectedBooks[i].Author, books.Books[i].Author)
	}
}

func (suite *HandlerTestSuite) TestDeleteBook() {
	_, err := PerformRequest(BuildJsonRequest(http.MethodPost, "/books", handler.InsertBookDTO{
		Title:       "Test Book",
		Description: "Test Description",
		Author:      "Test Author",
	}), handler.CreateBook, nil)
	assert.Nil(suite.T(), err)

	res, err := PerformRequest(BuildJsonRequest(http.MethodGet, "/books", nil), handler.GetBooks, nil)
	assert.Nil(suite.T(), err)
	books := ParseBodyString[handler.BooksDTO](res.Body)
	assert.Equal(suite.T(), 1, len(books.Books))

	res, err = PerformRequest(BuildJsonRequest(http.MethodDelete, "/books", nil), handler.DeleteBook, &PathParams{
		"bookId": strconv.Itoa(int(books.Books[0].ID)),
	})
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, res.Code)
	assert.Equal(suite.T(), handler.SuccessDTO{Success: true}, ParseBodyString[handler.SuccessDTO](res.Body))

	res, err = PerformRequest(BuildJsonRequest(http.MethodGet, "/books", nil), handler.GetBooks, nil)
	assert.Nil(suite.T(), err)
	books = ParseBodyString[handler.BooksDTO](res.Body)
	assert.Equal(suite.T(), 0, len(books.Books))
}
