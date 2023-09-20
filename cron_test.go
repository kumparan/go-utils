package utils

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestCron_GetCronNextAt(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2010-01-02T15:00:00Z")
	patch := gomonkey.ApplyFunc(time.Now, func() time.Time { return now })
	defer patch.Reset()

	assert.Equal(t, now.Format(cronNextAtTimeFormat), GetCronNextAt("* ngaco * cron nya ")) // wrong cron tab,return current time as next at
	assert.Equal(t, now.Add(1*time.Hour).Format(cronNextAtTimeFormat), GetCronNextAt("@hourly"))
	assert.Equal(t, now.Add(1*time.Minute).Format(cronNextAtTimeFormat), GetCronNextAt("*/1 * * * *"))
	assert.Equal(t, now.Add(1*time.Hour).Format(cronNextAtTimeFormat), GetCronNextAt("0 */1 * * *"))  // every hour
	assert.Equal(t, now.Add(2*time.Hour).Format(cronNextAtTimeFormat), GetCronNextAt("0 17 */1 * *")) // every 17:00
	assert.Equal(t, now.Add(30*time.Second).Format(cronNextAtTimeFormat), GetCronNextAt("@every 30s"))
}
