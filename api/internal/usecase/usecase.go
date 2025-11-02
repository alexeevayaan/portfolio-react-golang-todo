package usecase

import (
	"context"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/domain"
	"github.com/google/uuid"
)

type Redis interface{
	IsExists(ctx context.Context, idempotencyKey string) bool
}

type Postgres interface {
	CreateTodo(ctx context.Context, todo domain.Todo) error
	GetTodo(ctx context.Context, id uuid.UUID) (domain.Todo, error)
}

type UseCase struct{
	postgres Postgres
	redis Redis
}

func New(postgres Postgres, redis Redis) *UseCase{
	return &UseCase{
		postgres: postgres,
		redis: redis,
	}
}