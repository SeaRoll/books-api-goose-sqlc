-- name: ListBooksPaged :many
SELECT *
FROM books
OFFSET $1
LIMIT $2;

-- name: InsertBook :one
INSERT INTO books (title, author, description)
VALUES ($1, $2, $3)
RETURNING id;

-- name: UpdateBook :exec
UPDATE books
SET title = $2, author = $3, description = $4
WHERE id = $1;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: DeleteAllBooks :exec
DELETE FROM books;

-- name: GetConditions :many
SELECT time_bucket('1 minute', time)::timestamptz AS bucket,
       avg(temperature) AS avg_temp
FROM conditions
GROUP BY bucket
ORDER BY bucket ASC;

-- name: InsertCondition :exec
INSERT INTO conditions (time, location, device, temperature, humidity)
VALUES ($1, $2, $3, $4, $5);
