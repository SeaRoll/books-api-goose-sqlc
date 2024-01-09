package api

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// INIT ROUTES HERE
	initBooksRoutes(e)

	return e
}

func Run(e *echo.Echo) {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	e.Logger.Fatal(e.Start(":" + port))
}
