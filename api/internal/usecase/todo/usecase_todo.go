package todo

import (
	"context"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/domain"
)

type Redis interface{
	IsExists(ctx context.Context, idempotencyKey string) bool
}

type Postgres interface {
	CreateTodo(ctx context.Context, todo domain.Todo) error
}

type UseCaseTodo struct{
	postgres Postgres
	redis Redis
}

func New(postgres Postgres, redis Redis) *UseCaseTodo{
	return &UseCaseTodo{
		postgres: postgres,
		redis: redis,
	}
}