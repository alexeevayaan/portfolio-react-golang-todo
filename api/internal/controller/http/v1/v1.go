package v1

import todo "github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/usecase"

type Handler struct{
	usecase *todo.UseCase
}

func New(uc *todo.UseCase) *Handler{
	return &Handler{
		usecase: uc,
	}
}