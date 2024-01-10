# API Documentation

## Endpoints

### Get Books

**Endpoint: /api/v1/books**

**Method: GET**

**Description:** Retrieve list of books

**Example Request**

```sh
GET /api/v1/books
```

**Example Response**

```json
{
  "books": [
    {
      "id:": 1,
      "title": "The Lord of the Rings",
      "author": "J.R.R. Tolkien",
      "description": "The Lord of the Rings is an epic high-fantasy novel written by English author and scholar J. R. R. Tolkien."
    }
  ]
}
```

### Create Book

**Endpoint: /api/v1/books**

**Method: POST**

**Description:** Create a new book

**Example Request**

```sh
POST /api/v1/books
```

**Example Request Body**

```json
{
  "title": "The Lord of the Rings",
  "author": "J.R.R. Tolkien",
  "description": "The Lord of the Rings is an epic high-fantasy novel written by English author and scholar J. R. R. Tolkien."
}
```

**Example Response**

```json
{
  "success": true
}
```

### Update Book

**Endpoint: /api/v1/books**

**Method: PUT**

**Description:** Update an existing book

**Example Request**

```sh
PUT /api/v1/books
```

**Example Request Body**

```json
{
  "id": 1,
  "title": "The Lord of the Rings",
  "author": "J.R.R. Tolkien",
  "description": "The Lord of the Rings is an epic high-fantasy novel written by English author and scholar J. R. R. Tolkien."
}
```

**Example Response**

```json
{
  "success": true
}
```

### Delete Book

**Endpoint: /api/v1/books/:bookId**

**Method: DELETE**

**Description:** Delete an existing book

**Example Request**

```sh
DELETE /api/v1/books/1
```

**Example Response**

```json
{
  "success": true
}
```

### Insert condition

**Endpoint: /api/v1/conditions**

**Method: POST**

**Description:** Insert a new condition

**Example Request**

```sh
POST /api/v1/conditions
```

**Example Request Body**

```json
{
  "occurredAt": "2020-01-01T00:00:00.000Z",
  "location": "Bedroom",
  "device": "Thermometer",
  "temperature": 20,
  "humidity": 50
}
```

**Example Response**

```json
{
  "success": true
}
```

### Get conditions (avg per day)

**Endpoint: /api/v1/conditions**

**Method: GET**

**Description:** Retrieve list of conditions (avg per day)

**Example Request**

```sh
GET /api/v1/conditions
```

**Example Response**

```json
[
  {
    "day": "2020-01-01T00:00:00.000Z",
    "avgTemp": 0.52348
  }
]
```
