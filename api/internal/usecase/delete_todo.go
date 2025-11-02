package usecase

import (
	"context"
	"fmt"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/domain"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/dto"
	"github.com/google/uuid"
)

func (u UseCase) DeleteTodo(ctx context.Context, input dto.DeleteTodoInput) error{

	id, err := uuid.Parse(input.Id)
	if err != nil{
		return domain.ErrUUIDInvalid
	}
	
	err = u.postgres.DeleteTodo(ctx, id)
	if err !=nil{
		return fmt.Errorf("postgres.DeleteTodo: %w", err)
	}

	return nil
} 