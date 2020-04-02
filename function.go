package utils

import (
	"runtime"
	"time"
)

// RetryStopper :nodoc:
type RetryStopper struct {
	error
}

// Retry :nodoc:
func Retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(RetryStopper); ok {
			// Return the original error for later checking
			return s.error
		}

		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return Retry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}

// NewRetryStopper :nodoc:
func NewRetryStopper(err error) RetryStopper {
	return RetryStopper{err}
}

// MyCaller will return the method caller. skip value defines how many steps to be skipped.
// skip=0 will always return the MyCaller
// skip=1 returns the caller of the MyCaller
// and so on...
func MyCaller(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return details.Name()
	}
	return "failed to identify method caller"
}
