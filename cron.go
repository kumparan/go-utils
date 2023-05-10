package utils

import (
	"time"

	"github.com/robfig/cron/v3"
)

var cronNextAtTimeFormat = "2006-01-02T15:04:05"

// GetCronNextAt supports
//   - Standard crontab specs, e.g. "* * * * ?"
//   - Descriptors, e.g. "@midnight", "@every 1h30m"
//   - if cron parsing error then return current time
func GetCronNextAt(cronTab string) string {
	now := time.Now()
	var schedule, err = cron.ParseStandard(cronTab)
	if err != nil {
		return now.Format(cronNextAtTimeFormat)
	}

	return schedule.Next(now).Format(cronNextAtTimeFormat)
}
