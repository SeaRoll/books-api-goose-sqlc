package test

import (
	"books-api/api"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func ParseBodyString[T any](body *bytes.Buffer) T {
	result := new(T)
	err := json.Unmarshal(body.Bytes(), result)
	if err != nil {
		log.Fatalf("Could not parse body: %s", err)
	}
	return *result
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

type PathParams map[string]string

func PerformRequest(req *http.Request, fn func(c echo.Context) error, pathParams *PathParams) (*httptest.ResponseRecorder, error) {
	e := api.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if pathParams != nil {
		for k, v := range *pathParams {
			c.SetParamNames(k)
			c.SetParamValues(v)
		}
	}

	if err := fn(c); err != nil {
		return nil, err
	}
	return rec, nil
}
