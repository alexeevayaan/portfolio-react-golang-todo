package http

import (
	v1 "github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/controller/http/v1"
	todo "github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/usecase"
	"github.com/go-chi/chi/v5"
)


func Router(r *chi.Mux, uc * todo.UseCase){
	v1:=v1.New(uc) 

	r.Route("/", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/todo", v1.CreateTodo)
			r.Get("/todo/{id}", v1.GetTodo)
			r.Delete("/todo{id}", v1.DeleteTodo)			

		})
	})
}