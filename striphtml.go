package utils

import (
	"html"
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
		bmStrict = bluemonday.StrictPolicy()
	})
	return html.UnescapeString(bmStrict.Sanitize(s))
}
