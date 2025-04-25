package main

import (
	"log"

	"github.com/Guizzs26/currency-converter-go/internal/config"
	"github.com/Guizzs26/currency-converter-go/internal/db"
	"github.com/Guizzs26/currency-converter-go/internal/env"
	"github.com/Guizzs26/currency-converter-go/internal/handler"
	"github.com/Guizzs26/currency-converter-go/internal/service"
	"github.com/Guizzs26/currency-converter-go/internal/store"
)

type application struct {
	config config.Config
	store  store.Storage
}

func main() {
	cfg := env.InitConfig()

	db, err := db.NewPostgresConnection(cfg.DB.ConnStr, cfg.DB.MaxOpenConns, cfg.DB.MaxIdleConns, cfg.DB.MaxIdleTime)
	if err != nil {
		log.Fatal(err)
	}
	storage := store.NewPostgresStorage(db)

	app := &application{
		config: cfg,
		store:  *storage,
	}

	conversionService := service.NewConversionService(app.store.Conversion)
	conversionHandler := handler.NewConversionHandler(conversionService)

	r := app.configureRouter(conversionHandler)

	log.Fatal(app.bootstrap(r))
}
