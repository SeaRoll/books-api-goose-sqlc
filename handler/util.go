package handler

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func returnSuccess(c echo.Context) error {
	return c.JSON(200, SuccessDTO{Success: true})
}

func returnUserError(c echo.Context, err error) error {
	log.Printf("User error: %v", err)
	return c.JSON(400, ErrorDTO{Message: err.Error()})
}

func returnServerError(c echo.Context, err error) error {
	log.Printf("Internal server error: %v", err)
	return c.JSON(500, ErrorDTO{Message: "Internal server error"})
}

func parseToInt32(s string) (int32, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}
	return int32(i), nil
}

func validateStruct(i interface{}) error {
	v := validator.New()
	return v.Struct(i)
}
