package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Todo struct{
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Title string `json:"title" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"required,min=3,max=250"`
	Completed bool `json:"completed"`
}

var validate = validator.New(validator.WithRequiredStructEnabled()) 

func NewTodo(title string, description string)(Todo, error){

	t:=Todo{
		ID: uuid.New(),
		Title: title,
		Description: description,
		Completed: false,
	}

	if err:= t.Validate(); err !=nil{
		return Todo{}, fmt.Errorf("t.Validate: %w", err)
	}

	return t, nil
}

func (t Todo) Validate() error{
	err:= validate.Struct(t)

	if err != nil{
		return fmt.Errorf("validate.Struct: %w", err)
	}

	return nil
}