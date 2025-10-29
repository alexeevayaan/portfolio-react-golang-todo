package http

import (
	h "net/http"

	v1 "github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/controller/http/v1"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/usecase/todo"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/render"
	"github.com/go-chi/chi/v5"
)


func TodoRouter(r *chi.Mux, uc * todo.UseCaseTodo){
	ver1:=v1.New(uc) 

	r.Route("/", func(r chi.Router) {
		r.Get("/a", func(w h.ResponseWriter, r *h.Request) {
			render.JSON(w,map[string]string{
				"abas":"abas",
			},201)
		})
		r.Post("/todo", ver1.CreateTodo)
	})
}