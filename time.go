package utils

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/goodsign/monday"
	"github.com/oklog/ulid"
)

// FormatTimeRFC3339 Format time according to RFC3339Nano
func FormatTimeRFC3339(t *time.Time) (s string) {
	if t != nil {
		s = t.Format(time.RFC3339Nano)
	}
	return
}

// FormatToWesternIndonesianTime format to western indonesian time
// expected format: 12 April 2020 14:30 WIB
func FormatToWesternIndonesianTime(layout string, t *time.Time) string {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
	}
	return monday.Format(t.In(location), layout, monday.LocaleIdID)
}

// StringMillisToTime convert millis to time in UTC
func StringMillisToTime(millis string) time.Time {
	return time.Unix(0, StringToInt64(millis)*int64(time.Millisecond)).UTC()
}

// StringMillisToPointerTime convert millis to pointer time in UTC
func StringMillisToPointerTime(millis string) *time.Time {
	if millis == "" {
		return nil
	}

	t := StringMillisToTime(millis)
	return &t
}

// Int64MillisToTime convert millis to time in UTC
func Int64MillisToTime(millis int64) time.Time {
	return time.Unix(0, millis*int64(time.Millisecond)).UTC()
}

// Int64MillisToPointerTime convert millis to pointer time in UTC
func Int64MillisToPointerTime(millis int64) *time.Time {
	t := Int64MillisToTime(millis)
	return &t
}

func GenerateULIDFromTime(t time.Time) string {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return strings.ToLower(ulid.MustNew(ulid.Timestamp(t), entropy).String())
}
