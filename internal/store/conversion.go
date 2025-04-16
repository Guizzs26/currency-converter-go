package store

import "context"

type ConversionRepository interface {
	SaveConversion(context.Context)
	GetExchangeRate(context.Context)
	ListSupportedCurrencies(context.Context)
}

type Conversion struct {
}
