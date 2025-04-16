package main

import (
	"log"
)

func main() {
	cfg := config{
		addr: ":3333",
	}
	app := &application{
		config: cfg,
	}

	r := app.configureRouter()

	log.Fatal(app.bootstrap(r))
}
