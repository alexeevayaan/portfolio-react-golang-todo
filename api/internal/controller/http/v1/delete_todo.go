package v1

import (
	"net/http"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/dto"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/render"
	"github.com/go-chi/chi/v5"
)


func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	input := dto.DeleteTodoInput{
		Id: chi.URLParam(r, "id"),
	}

	err := h.usecase.DeleteTodo(ctx, input)
	if err != nil{
		render.Error(ctx, w, err, http.StatusBadRequest, "request failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}