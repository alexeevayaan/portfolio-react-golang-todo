package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() *chi.Mux{
	r:=chi.NewRouter()
	
	r.Use(middleware.Recoverer)

	return r
}