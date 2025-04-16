package store

import "database/sql"

type Storage struct {
	Conversion ConversionRepository
}

func NewPostgresStorage(db *sql.DB) *Storage {
	return &Storage{
		Conversion: NewPostgreConversionRepository(db),
	}
}
