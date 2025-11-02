package v1

import (
	"errors"
	"net/http"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/domain"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/dto"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/render"
	"github.com/go-chi/chi/v5"
)


func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	input := dto.GetTodoInput{
		Id: chi.URLParam(r, "id"),
	}

	output, err := h.usecase.GetTodo(ctx, input)
	if err != nil{
		ErrStatus := http.StatusBadRequest
		if errors.Is(err, domain.ErrNotFound){
			ErrStatus = http.StatusNotFound
		}
		render.Error(ctx, w, err, ErrStatus, "request failed")
	}

	render.JSON(w, output, http.StatusOK)
} 