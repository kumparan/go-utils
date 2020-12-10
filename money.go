package utils

import (
	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
)

// FormatToIndonesianMoney format money into Indonesian
// example: Rp10.000.000,00
func FormatToIndonesianMoney(dec decimal.Decimal) string {
	return "Rp" + accounting.FormatNumberDecimal(dec, 0, ".", ",")
}
