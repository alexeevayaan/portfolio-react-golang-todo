package v1

import (
	"context"

	http_server "github.com/alexeevayaan/portfolio-react-golang-todo/api/gen/http/todo_v1/server"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/dto"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/render"
)

func (h *Handler) DeleteTodoByID(ctx context.Context, request http_server.DeleteTodoByIDRequestObject) (http_server.DeleteTodoByIDResponseObject, error) {
	input := dto.DeleteTodoInput{
		Id: request.ID.String(),
	}

	err := h.usecase.DeleteTodo(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")
		return http_server.DeleteTodoByID400JSONResponse{Error: err.Error()}, nil
	}

	return http_server.DeleteTodoByID204Response{}, nil
}
