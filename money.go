package utils

import (
	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
)

// FormatIntoIndonesianMoney format money into Indonesian
// example: Rp10.000.000,00
func FormatIntoIndonesianMoney(dec decimal.Decimal) string {
	return "Rp" + accounting.FormatNumberDecimal(dec, 0, ".", ",")
}
