package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_FormatTimeRFC3339(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		s := FormatTimeRFC3339(nil)
		assert.EqualValues(t, "", s)
	})

	t.Run("Now", func(t *testing.T) {
		now, err := time.Parse(time.RFC3339Nano, "2016-06-06T03:55:00Z")
		assert.NoError(t, err)
		s := FormatTimeRFC3339(&now)
		assert.EqualValues(t, "2016-06-06T03:55:00Z", s)
	})
}

func Test_FormatToWesternIndonesianTime(t *testing.T)  {
	t.Run("Success", func(t *testing.T) {
		now, err := time.Parse(time.RFC3339Nano, "2016-12-07T03:55:00Z")
		assert.NoError(t, err)
		s := FormatToWesternIndonesianTime(&now)
		assert.EqualValues(t, "07 Des 2016 10:55 WIB", s)
	})
}
