package utils

import (
	"net/http"
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

// IsURLReachable check is the url reachable
func IsURLReachable(url string) bool {
	res, err := http.Head(url) //nolint:gosec
	if err != nil {
		return false
	}
	defer func() {
		_ = res.Body.Close()
	}()

	return res.StatusCode == http.StatusOK
}
