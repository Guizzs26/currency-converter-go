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
	query := `
		INSERT INTO conversions (from, to, amount, exchange_rate, converted_amount)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	err := cs.DB.QueryRowContext(
		ctx,
		query,
		conv.From,
		conv.To,
		conv.Amount,
		conv.ExchangeRate,
		conv.ConvertedAmount,
	).Scan(
		&conv.ID,
		&conv.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
