package handler

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
