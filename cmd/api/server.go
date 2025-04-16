package main

import (
	"log"
	"net/http"
	"time"
)

func (app *application) bootstrap(r http.Handler) error {
	srv := &http.Server{
		Addr:              app.config.Addr,
		Handler:           r,
		ReadTimeout:       time.Second * 10, // Max time the server waits to read the entire request (header + body)
		ReadHeaderTimeout: time.Second * 5,  // Max time the server waits to read only the request headers
		WriteTimeout:      time.Second * 10, // Max time the server has to write the entire response to the client
		IdleTimeout:       time.Second * 60, // Max time the server waits to keep a connection inactive (keep-alive)
	}

	log.Printf("server started at: http://localhost%s", app.config.Addr)
	return srv.ListenAndServe()
}
