package http

import (
	http_server "github.com/alexeevayaan/portfolio-react-golang-todo/api/gen/http/todo_v1/server"
	v1 "github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/controller/http/v1"
	todo "github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/usecase"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/logger"
	"github.com/go-chi/chi/v5"
)

func Router(r *chi.Mux, uc *todo.UseCase) {
	v1 := v1.New(uc)

	r.Route("/v1", func(r chi.Router) {
		r.Use(logger.Middleware)

		mux := http_server.NewStrictHandler(v1, []http_server.StrictMiddlewareFunc{})
		http_server.HandlerFromMux(mux, r)
		// r.Post("/todo", v1.CreateTodo)
		// r.Get("/todo/{id}", v1.GetTodo)
		// r.Delete("/todo/{id}", v1.DeleteTodo)
	})
}
