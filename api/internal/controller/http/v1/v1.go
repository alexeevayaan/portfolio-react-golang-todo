package v1

import "github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/usecase/todo"

type Handler struct{
	usecase *todo.UseCaseTodo
}

func New(uc *todo.UseCaseTodo) *Handler{
	return &Handler{
		usecase: uc,
	}
}