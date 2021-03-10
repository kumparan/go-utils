package utils

import (
	"fmt"

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
	cur, err := currency.ParseISO(currencyCode)
	if err != nil {
		return fmt.Sprintf("%s%s", currencyCode, accounting.FormatNumberDecimal(value, 2, ",", "."))
	}

	scale, _ := currency.Cash.Rounding(cur) // fractional digits
	unit, _ := value.Float64()
	dec := number.Decimal(unit, number.Scale(scale))

	p := message.NewPrinter(language.English)
	if currencyCode == "IDR" {
		p = message.NewPrinter(language.Indonesian)
	}

	return p.Sprintf("%v%v", currency.Symbol(cur), dec)
}
