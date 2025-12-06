package v1

import (
	"context"

	http_server "github.com/alexeevayaan/portfolio-react-golang-todo/api/gen/http/todo_v1/server"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/dto"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/render"
)

func (h *Handler) CreateTodo(ctx context.Context, request http_server.CreateTodoRequestObject) (http_server.CreateTodoResponseObject, error) {
	input := dto.CreateTodoInput{
		Title:       request.Body.Title,
		Description: request.Body.Description,
	}

	output, err := h.usecase.CreateTodo(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")

		return http_server.CreateTodo400JSONResponse{
			Error: err.Error(),
		}, nil
	}

	return http_server.CreateTodo200JSONResponse{
		ID: output.ID,
	}, nil
}
