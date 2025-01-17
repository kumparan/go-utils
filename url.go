package utils

import (
	"net/http"
	"net/url"
	"path"

	log "github.com/sirupsen/logrus"
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
	res, err := http.Get(url) //nolint:gosec
	if err != nil {
		log.WithField("url", url).Error(err)
		return false
	}
	defer func() {
		_ = res.Body.Close()
	}()

	log.Infof("statusCode: %d", res.StatusCode)
	return res.StatusCode < 500
}
