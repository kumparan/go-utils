package utils

import (
	"sync"

	"github.com/microcosm-cc/bluemonday"
)

var (
	bmStrict *bluemonday.Policy
	once     sync.Once
)

// StripHTML strips all HTML from a string, duh
func StripHTML(s string) string {
	once.Do(func() {
		if bmStrict == nil {
			bmStrict = bluemonday.StrictPolicy()
		}
	})
	return bmStrict.Sanitize(s)
}
