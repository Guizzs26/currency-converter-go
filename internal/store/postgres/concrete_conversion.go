package postgres

import (
	"context"
	"database/sql"

	"github.com/Guizzs26/currency-converter-go/internal/model"
)

type PostgresConversionStore struct {
	DB *sql.DB
}

func NewPostgresConversionStore(db *sql.DB) *PostgresConversionStore {
	return &PostgresConversionStore{
		DB: db,
	}
}

func (cs *PostgresConversionStore) SaveConversion(ctx context.Context, conv *model.Conversion) error {
	return nil
}
