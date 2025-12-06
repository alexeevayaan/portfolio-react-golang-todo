package v1

import (
	"context"
	"errors"

	http_server "github.com/alexeevayaan/portfolio-react-golang-todo/api/gen/http/todo_v1/server"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/domain"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/dto"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/render"
)

func (h *Handler) GetTodoByID(ctx context.Context, request http_server.GetTodoByIDRequestObject) (http_server.GetTodoByIDResponseObject, error) {
	input := dto.GetTodoInput{
		Id: request.ID.String(),
	}

	output, err := h.usecase.GetTodo(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return http_server.GetTodoByID404JSONResponse{Error: err.Error()}, nil
		default:
			return http_server.GetTodoByID400JSONResponse{Error: err.Error()}, nil
		}
	}

	var todo http_server.GetTodoByID200JSONResponse
	todo.ID = output.ID
	todo.Title = output.Title
	todo.Description = output.Description
	todo.Completed = output.Completed
	todo.CreatedAt = output.CreatedAt
	todo.UpdatedAt = output.UpdatedAt

	return todo, nil
}
