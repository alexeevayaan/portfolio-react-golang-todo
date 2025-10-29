package v1

import (
	"net/http"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/render"
	"github.com/rs/zerolog/log"
)

func (h *Handler) CreateTodo (w http.ResponseWriter, r *http.Request){
	// ctx:= r.Context()
	log.Info().Msg("CreateTodo")
	render.JSON(w, map[string]string{
		"abas":"asdasd",
	},http.StatusCreated)
}