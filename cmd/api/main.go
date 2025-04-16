package main

import (
	"log"

	"github.com/Guizzs26/currency-converter-go/internal/env"
	"github.com/Guizzs26/currency-converter-go/internal/env/config"
)

type application struct {
	config config.Config
}

func main() {
	cfg := env.InitConfig()

	app := &application{
		config: cfg,
	}

	r := app.configureRouter()

	log.Fatal(app.bootstrap(r))
}
