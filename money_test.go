package utils

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FormatToIndonesianMoney(t *testing.T)  {
	t.Run("Success", func(t *testing.T) {
		price := decimal.NewFromFloat(10000000.88)
		assert.Equal(t, "Rp10.000.001", FormatToIndonesianMoney(price))
	})
}