package handler

import "time"

type SuccessDTO struct {
	Success bool `json:"success"`
}

type ErrorDTO struct {
	Message string `json:"message"`
}

type InsertBookDTO struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Author      string `json:"author" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"required,min=1"`
}

type UpdateBookDTO struct {
	ID          int32  `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Author      string `json:"author" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"required,min=1"`
}

type BookDTO struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type BooksDTO struct {
	Books []BookDTO `json:"books"`
}

type InsertConditionDTO struct {
	OccurredAt  time.Time `json:"occurredAt" validate:"required"`
	Location    string    `json:"location" validate:"required,min=1,max=255"`
	Device      string    `json:"device" validate:"required,min=1,max=255"`
	Temperature float64   `json:"temperature" validate:"required"`
	Humidity    float64   `json:"humidity" validate:"required"`
}

type BucketConditionDTO struct {
	Day     time.Time `json:"day"`
	AvgTemp float64   `json:"avgTemp"`
}
