package utils

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_FormatToIndonesianMoney(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		price := decimal.NewFromFloat(10000000.88)
		result := FormatToIndonesianMoney(price)
		assert.Equal(t, result, "Rp10.000.001")
	})
}

func TestFormatMoney(t *testing.T) {
	assert.Equal(t, "Rp10.000.000", FormatMoney(decimal.NewFromFloat(10000000), "IDR"))
	assert.Equal(t, "Rp10.000.001", FormatMoney(decimal.NewFromFloat(10000000.88), "IDR"))
	assert.Equal(t, "$1.49", FormatMoney(decimal.NewFromFloat(1.49), "USD"))
	assert.Equal(t, "$1.49", FormatMoney(decimal.NewFromFloat(1.4889), "USD"))
	assert.Equal(t, "¥149", FormatMoney(decimal.NewFromFloat(149), "JPY"))
	assert.Equal(t, "¥149", FormatMoney(decimal.NewFromFloat(148.89), "JPY"))
	assert.Equal(t, "€1.49", FormatMoney(decimal.NewFromFloat(1.49), "EUR"))
	assert.Equal(t, "€1.49", FormatMoney(decimal.NewFromFloat(1.4889), "EUR"))
	assert.Equal(t, "£1.49", FormatMoney(decimal.NewFromFloat(1.49), "GBP"))
	assert.Equal(t, "£1.49", FormatMoney(decimal.NewFromFloat(1.4889), "GBP"))
	assert.Equal(t, "A$1.49", FormatMoney(decimal.NewFromFloat(1.49), "AUD"))
	assert.Equal(t, "A$1.49", FormatMoney(decimal.NewFromFloat(1.4889), "AUD"))
	assert.Equal(t, "CA$1.49", FormatMoney(decimal.NewFromFloat(1.49), "CAD"))
	assert.Equal(t, "CA$1.49", FormatMoney(decimal.NewFromFloat(1.4889), "CAD"))
}
