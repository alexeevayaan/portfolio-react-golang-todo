package dto

import "github.com/google/uuid"

type CreateTodoOutput struct{
	ID uuid.UUID `json:"id"`
}

type CreateTodoInput struct{
	Title string `json:"title"`
	Description string `json:"description"`
}