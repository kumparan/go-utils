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

func Test_FormatToWesternIndonesianTime(t *testing.T) {
	t.Run("Success with 3 digit month name", func(t *testing.T) {
		layout := "02 Jan 2006 15:04 WIB"
		now, err := time.Parse(time.RFC3339Nano, "2016-12-07T03:55:00Z")
		assert.NoError(t, err)
		s := FormatToWesternIndonesianTime(layout, &now)
		assert.EqualValues(t, "07 Des 2016 10:55 WIB", s)
	})

	t.Run("Success with full month name", func(t *testing.T) {
		layout := "02 January 2006 15:04 WIB"
		now, err := time.Parse(time.RFC3339Nano, "2016-12-07T03:55:00Z")
		assert.NoError(t, err)
		s := FormatToWesternIndonesianTime(layout, &now)
		assert.EqualValues(t, "07 Desember 2016 10:55 WIB", s)
	})
}

func Test_StringMillisToTime(t *testing.T) {
	millis := "1615963569481"
	expected := "2021-03-17T06:46:09.481Z"

	assert.Equal(t, expected, StringMillisToTime(millis).Format(time.RFC3339Nano))
}

func Test_StringMillisToPointerTime(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		millis := "1615963569481"
		expected := "2021-03-17T06:46:09.481Z"

		assert.Nil(t, nil)
		assert.EqualValues(t, expected, StringMillisToPointerTime(millis).Format(time.RFC3339Nano))
	})

	t.Run("success", func(t *testing.T) {
		millis := ""
		assert.Nil(t, StringMillisToPointerTime(millis))
	})
}

func Test_Int64MillisToTime(t *testing.T) {
	millis := int64(1615963569481)
	expected := "2021-03-17T06:46:09.481Z"

	assert.Equal(t, expected, Int64MillisToTime(millis).Format(time.RFC3339Nano))
}

func Test_Int64MillisToPointerTime(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		millis := int64(1615963569481)
		expected := "2021-03-17T06:46:09.481Z"

		assert.Nil(t, nil)
		assert.EqualValues(t, expected, Int64MillisToPointerTime(millis).Format(time.RFC3339Nano))
	})

	t.Run("0 millis", func(t *testing.T) {
		millis := int64(0)
		expected := "1970-01-01T00:00:00Z"

		assert.EqualValues(t, expected, Int64MillisToPointerTime(millis).Format(time.RFC3339Nano))
	})
}
