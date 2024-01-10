// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Book struct {
	ID          int32
	Title       string
	Author      string
	Description string
}

type Condition struct {
	Time        pgtype.Timestamptz
	Location    string
	Device      string
	Temperature float64
	Humidity    float64
	Value       []byte
}
