package handler

import (
	"books-api/db"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func mapBooks(books []db.Book) []BookDTO {
	result := make([]BookDTO, len(books))
	for i, book := range books {
		result[i] = BookDTO{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Description: book.Description,
		}
	}
	return result
}

func GetBooks(c echo.Context) error {
	books := []db.Book{}
	err := db.WithTX(c.Request().Context(), func(ctx context.Context, queries *db.Queries, tx pgx.Tx) error {
		var err error
		books, err = queries.ListBooksPaged(ctx, tx, db.ListBooksPagedParams{Limit: 10, Offset: 0})
		return err
	})
	if err != nil {
		return returnServerError(c, err)
	}
	return c.JSON(200, BooksDTO{Books: mapBooks(books)})
}

func CreateBook(c echo.Context) error {
	insertParams := InsertBookDTO{}
	err := c.Bind(&insertParams)
	if err != nil {
		return returnUserError(c, err)
	}
	err = validateStruct(insertParams)
	if err != nil {
		return returnUserError(c, err)
	}

	err = db.WithTX(c.Request().Context(), func(ctx context.Context, queries *db.Queries, tx pgx.Tx) error {
		_, err := queries.InsertBook(ctx, tx, db.InsertBookParams{
			Title:       insertParams.Title,
			Author:      insertParams.Author,
			Description: insertParams.Description,
		})
		return err
	})
	if err != nil {
		return returnServerError(c, err)
	}
	return returnSuccess(201, c)
}

func UpdateBook(c echo.Context) error {
	updateParams := UpdateBookDTO{}
	err := c.Bind(&updateParams)
	if err != nil {
		return returnUserError(c, err)
	}
	err = validateStruct(updateParams)
	if err != nil {
		return returnUserError(c, err)
	}

	err = db.WithTX(c.Request().Context(), func(ctx context.Context, queries *db.Queries, tx pgx.Tx) error {
		return queries.UpdateBook(ctx, tx, db.UpdateBookParams{
			ID:          updateParams.ID,
			Title:       updateParams.Title,
			Author:      updateParams.Author,
			Description: updateParams.Description,
		})
	})
	if err != nil {
		return returnServerError(c, err)
	}

	return returnSuccess(200, c)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("bookId")
	if id == "" {
		return returnUserError(c, errors.New("bookId is required"))
	}

	idInt, err := parseToInt32(id)
	if err != nil {
		return returnUserError(c, err)
	}

	err = db.WithTX(c.Request().Context(), func(ctx context.Context, queries *db.Queries, tx pgx.Tx) error {
		return queries.DeleteBook(ctx, tx, idInt)
	})
	if err != nil {
		return returnServerError(c, err)
	}

	return returnSuccess(200, c)
}
