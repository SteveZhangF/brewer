package basic

import (
	"github.com/shopspring/decimal"
)

type Currency struct {
	Value    decimal.Decimal `json:"value"`
	Currency string          `json:"currency"`
}
