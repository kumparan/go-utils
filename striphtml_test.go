package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StripHTML(t *testing.T) {
	t.Run("no html", func(t *testing.T) {
		assert.Equal(t, "podo wae mas'bro&&",
			StripHTML("podo wae mas'bro&&"))
	})

	t.Run("with html", func(t *testing.T) {
		assert.Equal(t, "dang sarupa bah",
			StripHTML("<a href=\"http://lappetkaw.com\">dang sarupa bah</a>"))
	})

	t.Run("with incomplete html", func(t *testing.T) {
		assert.Equal(t, ">plz click hier",
			StripHTML("><a href='https://www.downloadmoreram.com'>plz click hier"))
	})
}
