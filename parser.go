package utils

import "time"

// ParseDuration parse duration with default
func ParseDurationWithDefault(in string, defaultDuration time.Duration) time.Duration {
	dur, err := time.ParseDuration(in)
	if err != nil {
		return defaultDuration
	}
	return dur
}
