package utils

import (
	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

// FormatToIndonesianMoney format money into Indonesian
// example: Rp10.000.000
func FormatToIndonesianMoney(dec decimal.Decimal) string {
	return "Rp" + accounting.FormatNumberDecimal(dec, 0, ".", ",")
}

// FormatMoney format money by currency code (ISO 4217)
func FormatMoney(value decimal.Decimal, currencyCode string) string {
	cur, _ := currency.ParseISO(currencyCode) // ignore error, unknown currencyCode will be formatted as {{currencyCode}}{{value}}, ex ZXX200
	scale, _ := currency.Cash.Rounding(cur)   // fractional digits

	unit, _ := value.Float64()
	dec := number.Decimal(unit, number.Scale(scale))

	p := message.NewPrinter(language.English)
	if currencyCode == "IDR" {
		p = message.NewPrinter(language.Indonesian)
	}

	return p.Sprintf("%v%v", currency.Symbol(cur), dec)
}
