package utils

import (
	"fmt"
	"time"
)

const (
	_secondsPerDay = 86400
)

// GetDifferenceDaysForHumans return difference days for humans in indonesian language
func GetDifferenceDaysForHumans(startsAt time.Time, endsAt time.Time) string {
	numDay := (startsAt.Unix() - endsAt.Unix()) / _secondsPerDay

	switch {
	case numDay == 0:
		return "hari ini"
	case numDay == -1:
		return "besok"
	case numDay == 1:
		return "kemarin"
	case numDay < -1:
		return fmt.Sprintf("%d hari lagi", -(numDay))
	default:
		return fmt.Sprintf("%d hari yang lalu", numDay)
	}
}
