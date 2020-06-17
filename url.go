package utils

import (
	"net/url"
	"path"
)

// JoinURL joins URL with the path elements
func JoinURL(baseURL string, pathElements ...string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	var elements []string
	u.Path = path.Join(append(append(elements, u.Path), pathElements...)...)
	return u.String(), nil
}
