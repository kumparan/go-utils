package utils

import "strings"

// GetIDFromSlug in base62
func GetIDFromSlug(slug string) string {
	splittedSlug := strings.Split(slug, "-")
	return splittedSlug[len(splittedSlug)-1]
}
