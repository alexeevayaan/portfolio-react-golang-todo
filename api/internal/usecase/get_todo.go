package usecase

import (
	"context"
	"fmt"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/domain"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/dto"
	"github.com/google/uuid"
)


func (u UseCase) GetTodo(ctx context.Context,  input dto.GetTodoInput) (dto.GetTodoOutput, error){
	var output dto.GetTodoOutput

	id,err:= uuid.Parse(input.Id)
	if err !=nil{
		return output, domain.ErrUUIDInvalid
	}

	todo, err := u.postgres.GetTodo(ctx,id)
	if err !=nil{
		return output,fmt.Errorf("postgres.GetTodo: %w", err)
	}

	output.Todo = todo

	return output, nil

}