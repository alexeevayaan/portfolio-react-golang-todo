package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() *chi.Mux{
	r:=chi.NewRouter()
	
	r.Use(middleware.Recoverer)

	r.Get("/ayaan", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<div><a>authorize</a></div>`))
		
	})

	return r
}