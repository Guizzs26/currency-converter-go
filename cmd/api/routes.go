package main

import (
	"net/http"

	"github.com/Guizzs26/currency-converter-go/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) configureRouter(handler *handler.ConversionHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Ok"))
		})

		r.Route("/currency", func(r chi.Router) {
			r.Post("/", handler.Convert)
		})
	})

	return r
}
