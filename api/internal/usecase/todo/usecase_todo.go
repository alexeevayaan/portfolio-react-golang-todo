package todo

import (
	"context"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/domain"
)

type Postgres interface {
	CreateTodo(ctx context.Context, todo domain.Todo) error
}

type UseCaseTodo struct{
	postgres Postgres
}

func New(postgres Postgres) *UseCaseTodo{
	return &UseCaseTodo{
		postgres: postgres,
	}
}