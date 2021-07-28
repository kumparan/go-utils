package utils

import (
	"log"
	"time"

	"github.com/golang-module/carbon"
	"github.com/goodsign/monday"
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

// TimeDurationFromNowForHuman get the time difference relative to now in human-friendly readable format
func TimeDurationFromNowForHuman(t time.Time) string {
	lang := carbon.NewLanguage()
	lang.SetResources(map[string]string{
		"seasons":        "Semi|Panas|Gugur|Dingin",
		"months":         "Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember",
		"months_short":   "Jan|Feb|Mar|Apr|Mei|Jun|Jul|Agu|Sep|Okt|Nov|Des",
		"weeks":          "Minggu|Senin|Selasa|Rabu|Kamis|Jumat|Sabtu",
		"weeks_short":    "Min|Sen|Sel|Rab|Kam|Jum|Sab",
		"constellations": "Aries|Taurus|Gemini|Cancer|Leo|Virgo|Libra|Scorpio|Sagittarius|Capricorn|Aquarius|Pisces",
		"year":           "1 tahun|%d tahun",
		"month":          "1 bulan|%d bulan",
		"week":           "1 minggu|%d minggu",
		"day":            "1 hari|%d hari",
		"hour":           "1 jam|%d jam",
		"minute":         "1 menit|%d menit",
		"second":         "1 detik|%d detik",
		"now":            "baru saja",
		"ago":            "%s yang lalu",
		"from_now":       "%s lagi",
		"before":         "%s sebelum",
		"after":          "%s setelah",
	})
	c := carbon.NewCarbon()
	c = c.SetLanguage(lang)
	c.Time = t

	return c.DiffForHumans()
}
