package util

import (
	"books-api/api"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func ParseBodyString[T any](body *bytes.Buffer, result *T) error {
	return json.Unmarshal(body.Bytes(), result)
}

func BuildJsonRequest(method string, url string, body interface{}) *http.Request {
	if body == nil {
		return httptest.NewRequest(method, url, nil)
	}
	req := httptest.NewRequest(method, url, EncodeJSON(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return req
}

func EncodeJSON(data interface{}) *bytes.Buffer {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)
	return b
}

func PerformRequest(req *http.Request, fn func(c echo.Context) error) (*httptest.ResponseRecorder, error) {
	e := api.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := fn(c); err != nil {
		return nil, err
	}
	return rec, nil
}
