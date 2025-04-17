package service

import (
	"context"
	"fmt"

	"github.com/Guizzs26/currency-converter-go/internal/model"
	"github.com/Guizzs26/currency-converter-go/internal/store"
)

type ConversionService interface {
	Convert(ctx context.Context, input *model.Conversion) (*model.Conversion, error)
}

type ConversionServiceImpl struct {
	repo store.ConversionRepository
}

func NewConversionService(repo store.ConversionRepository) *ConversionServiceImpl {
	return &ConversionServiceImpl{
		repo: repo,
	}
}

func (cs *ConversionServiceImpl) Convert(ctx context.Context, input *model.Conversion) (*model.Conversion, error) {
	convertedAmount := input.Amount * input.ExchangeRate

	conversion := &model.Conversion{
		From:            input.From,
		To:              input.To,
		Amount:          input.Amount,
		ExchangeRate:    input.ExchangeRate,
		ConvertedAmount: convertedAmount,
	}

	if err := cs.repo.SaveConversion(ctx, conversion); err != nil {
		return nil, fmt.Errorf("failed to save convertion: %w", err)
	}

	return conversion, nil
}
