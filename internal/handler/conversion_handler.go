package handler

import (
	"net/http"

	"github.com/Guizzs26/currency-converter-go/internal/helpers"
	"github.com/Guizzs26/currency-converter-go/internal/model"
	"github.com/Guizzs26/currency-converter-go/internal/service"
)

type ConversionHandler struct {
	conversionService service.ConversionService
}

func NewConversionHandler(s service.ConversionService) *ConversionHandler {
	return &ConversionHandler{
		conversionService: s,
	}
}

func (ch *ConversionHandler) Convert(w http.ResponseWriter, r *http.Request) {
	var req ConversionRequest

	if err := helpers.ReadJSON(w, r, &req); err != nil {
		helpers.BadRequestError(w, r, err)
		return
	}

	if err := helpers.Validate.Struct(req); err != nil {
		helpers.ValidationErrorJSON(w, r, err)
		return
	}

	conv := &model.Conversion{
		From:         req.From,
		To:           req.To,
		Amount:       req.Amount,
		ExchangeRate: req.ExchangeRate,
	}

	ctx := r.Context()

	result, err := ch.conversionService.Convert(ctx, conv)
	if err != nil {
		helpers.InternalServerError(w, r, err)
		return
	}

	resp := ConversionResponse{
		ID:              result.ID,
		From:            result.From,
		To:              result.To,
		Amount:          result.Amount,
		ExchangeRate:    result.ExchangeRate,
		ConvertedAmount: result.ConvertedAmount,
		CreatedAt:       result.CreatedAt,
	}

	if err := helpers.WriteJSON(w, resp, 201); err != nil {
		helpers.InternalServerError(w, r, err)
	}
}
