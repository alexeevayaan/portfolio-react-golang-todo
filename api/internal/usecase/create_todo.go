package usecase

import (
	"context"
	"fmt"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/domain"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/dto"
)

func (u *UseCase) CreateTodo(ctx context.Context, input dto.CreateTodoInput) (dto.CreateTodoOutput, error){

	var output dto.CreateTodoOutput

	todo, err:= domain.NewTodo(input.Title, input.Description)
	if err !=nil{
		return output, fmt.Errorf("domain.Todo: %w", err)
	}

	err = u.postgres.CreateTodo(ctx, todo)
	if err != nil {
		return output, fmt.Errorf("postgres.CreateTodo: %w", err)
	}

	output.ID = todo.ID

	return output, nil
}