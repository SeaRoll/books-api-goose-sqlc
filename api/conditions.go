package api

import (
	"books-api/handler"

	"github.com/labstack/echo/v4"
)

func initConditionsRoutes(e *echo.Echo) {
	// INIT ROUTES HERE
	g := e.Group("/api/v1/conditions")
	g.GET("", handler.GetConditions)
	g.POST("", handler.InsertCondition)
}
