package utils

import "strings"

// GetStoryIDFromStorySlug in base62
func GetStoryIDFromStorySlug(slug string) string {
	splittedSlug := strings.Split(slug, "-")
	return splittedSlug[len(splittedSlug)-1]
}
