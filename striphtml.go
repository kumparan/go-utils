package utils

import "github.com/microcosm-cc/bluemonday"

var (
	_bmStrict *bluemonday.Policy
)

// StripHTML strips all HTML from a string, duh
func StripHTML(s string) string {
	if _bmStrict == nil {
		_bmStrict = bluemonday.StrictPolicy()
	}
	return _bmStrict.Sanitize(s)
}
