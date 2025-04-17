package store

import (
	"context"
	"database/sql"

	"github.com/Guizzs26/currency-converter-go/internal/model"
	"github.com/Guizzs26/currency-converter-go/internal/store/postgres"
)

type ConversionRepository interface {
	SaveConversion(ctx context.Context, conv *model.Conversion) error
	// GetExchangeRate(ctx context.Context, from, to string) (float64, error)
	// ListSupportedCurrencies(ctx context.Context) ([]string, error)
}

type Storage struct {
	Conversion ConversionRepository
}

func NewPostgresStorage(db *sql.DB) *Storage {
	return &Storage{
		Conversion: postgres.NewPostgresConversionStore(db),
	}
}
