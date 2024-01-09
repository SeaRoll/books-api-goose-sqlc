package api

import (
	"books-api/handler"

	"github.com/labstack/echo/v4"
)

func initBooksRoutes(e *echo.Echo) {
	// INIT ROUTES HERE
	g := e.Group("/api/v1/books")
	g.GET("", handler.GetBooks)
	g.POST("", handler.CreateBook)
	g.PUT("", handler.UpdateBook)
	g.DELETE("/:bookId", handler.DeleteBook)
}
