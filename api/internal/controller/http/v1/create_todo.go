package v1

import (
	"encoding/json"
	"net/http"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/dto"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/render"
)

func (h *Handler) CreateTodo (w http.ResponseWriter, r *http.Request){
	ctx:= r.Context()

	input:= dto.CreateTodoInput{}

	err:= json.NewDecoder(r.Body).Decode(&input)
	if err !=nil{
		render.Error(ctx, w, err, http.StatusBadRequest, "json decode error")

		return
	}

	output, err := h.usecase.CreateTodo(ctx, input)

	if err != nil{
		render.Error(ctx, w, err, http.StatusBadRequest, "request failed")
		return
	}

	render.JSON(w, output, http.StatusOK)
}