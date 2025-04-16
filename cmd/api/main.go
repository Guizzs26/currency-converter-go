package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func main() {
	app := &application{
		config: config{
			addr: ":3333",
		},
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Ok"))
		})
	})

	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: r,
	}

	log.Printf("server started at: http://localhost%s", app.config.addr)
	srv.ListenAndServe()
}
