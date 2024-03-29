// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package db

import (
	"context"
	"encoding/json"
	"time"
)

const deleteAllBooks = `-- name: DeleteAllBooks :exec
DELETE FROM books
`

func (q *Queries) DeleteAllBooks(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllBooks)
	return err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteBook, id)
	return err
}

const getConditions = `-- name: GetConditions :many
SELECT time_bucket('1 day', time)::timestamptz AS bucket,
       avg(temperature) AS avg_temp
FROM conditions
GROUP BY bucket
ORDER BY bucket ASC
`

type GetConditionsRow struct {
	Bucket  time.Time `json:"bucket"`
	AvgTemp float64   `json:"avg_temp"`
}

func (q *Queries) GetConditions(ctx context.Context) ([]GetConditionsRow, error) {
	rows, err := q.db.QueryContext(ctx, getConditions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetConditionsRow
	for rows.Next() {
		var i GetConditionsRow
		if err := rows.Scan(&i.Bucket, &i.AvgTemp); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getConditionsAverageValueField = `-- name: GetConditionsAverageValueField :many
SELECT
  time_bucket('1 day', time)::timestamptz AS bucket,
  AVG((value->>($1)::text)::numeric) AS avg_value
FROM conditions
GROUP BY bucket
ORDER BY bucket ASC
`

type GetConditionsAverageValueFieldRow struct {
	Bucket   time.Time `json:"bucket"`
	AvgValue float64   `json:"avg_value"`
}

func (q *Queries) GetConditionsAverageValueField(ctx context.Context, dollar_1 string) ([]GetConditionsAverageValueFieldRow, error) {
	rows, err := q.db.QueryContext(ctx, getConditionsAverageValueField, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetConditionsAverageValueFieldRow
	for rows.Next() {
		var i GetConditionsAverageValueFieldRow
		if err := rows.Scan(&i.Bucket, &i.AvgValue); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertBook = `-- name: InsertBook :one
INSERT INTO books (title, author, description)
VALUES ($1, $2, $3)
RETURNING id
`

type InsertBookParams struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (q *Queries) InsertBook(ctx context.Context, arg InsertBookParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertBook, arg.Title, arg.Author, arg.Description)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertCondition = `-- name: InsertCondition :exec
INSERT INTO conditions (time, location, device, temperature, humidity, value)
VALUES ($1, $2, $3, $4, $5, $6)
`

type InsertConditionParams struct {
	Time        time.Time       `json:"time"`
	Location    string          `json:"location"`
	Device      string          `json:"device"`
	Temperature float64         `json:"temperature"`
	Humidity    float64         `json:"humidity"`
	Value       json.RawMessage `json:"value"`
}

func (q *Queries) InsertCondition(ctx context.Context, arg InsertConditionParams) error {
	_, err := q.db.ExecContext(ctx, insertCondition,
		arg.Time,
		arg.Location,
		arg.Device,
		arg.Temperature,
		arg.Humidity,
		arg.Value,
	)
	return err
}

const listBooksPaged = `-- name: ListBooksPaged :many
SELECT id, title, author, description
FROM books
OFFSET $1
LIMIT $2
`

type ListBooksPagedParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListBooksPaged(ctx context.Context, arg ListBooksPagedParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooksPaged, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Author,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBook = `-- name: UpdateBook :exec
UPDATE books
SET title = $2, author = $3, description = $4
WHERE id = $1
`

type UpdateBookParams struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	_, err := q.db.ExecContext(ctx, updateBook,
		arg.ID,
		arg.Title,
		arg.Author,
		arg.Description,
	)
	return err
}
