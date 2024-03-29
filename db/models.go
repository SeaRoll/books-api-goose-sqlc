// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"encoding/json"
	"time"
)

type Book struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type Condition struct {
	Time        time.Time       `json:"time"`
	Location    string          `json:"location"`
	Device      string          `json:"device"`
	Temperature float64         `json:"temperature"`
	Humidity    float64         `json:"humidity"`
	Value       json.RawMessage `json:"value"`
}
