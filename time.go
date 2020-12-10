package utils

import (
	"github.com/goodsign/monday"
	"log"
	"time"
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
