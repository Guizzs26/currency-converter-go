package handler

import (
	"time"

	"github.com/google/uuid"
)

type ConversionRequest struct {
	From         string  `json:"from"`
	To           string  `json:"to"`
	Amount       float64 `json:"amount"`
	ExchangeRate float64 `json:"exchange_rate"`
}

type ConversionResponse struct {
	ID              uuid.UUID `json:"id"`
	From            string    `json:"from"`
	To              string    `json:"to"`
	Amount          float64   `json:"amount"`
	ExchangeRate    float64   `json:"exchange_rate"`
	ConvertedAmount float64   `json:"converted_amount"`
	CreatedAt       time.Time `json:"created_at"`
}
