package utils

import (
	"errors"
	"strings"
)

var relativeProtocols = []string{"", "ftp", "http", "gopher", "nntp", "imap",
	"wais", "file", "https", "shttp", "mms",
	"prospero", "rtsp", "rtspu", "sftp",
	"svn", "svn+ssh", "ws", "wss"}

var netLocationProtocols = []string{"", "ftp", "http", "gopher", "nntp", "telnet",
	"imap", "wais", "file", "mms", "https", "shttp",
	"snews", "prospero", "rtsp", "rtspu", "rsync",
	"svn", "svn+ssh", "sftp", "nfs", "git", "git+ssh",
	"ws", "wss"}

var schemeChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-."

// JoinURL is to join base URL with another URL/path/directory based on python lib urljoin
// e.g. JoinURL(https://kumparan.com/trending , feed/category) => https://kumparan.com/feed/category
func JoinURL(base, url string) (string, error) {
	if len(base) == 0 {
		return url, nil
	}
	if len(url) == 0 {
		return base, nil
	}

	// Split baseURL and url into 3 components
	baseScheme, baseNetLocation, basePath, err := splitURL(base)
	if err != nil {
		return "", err
	}
	scheme, netLocation, path, err := splitURL(url)
	if err != nil {
		return "", err
	}

	if scheme == "" {
		scheme = baseScheme
	}

	// if url have a different scheme than baseURL return url
	if scheme != baseScheme || !ContainsString(relativeProtocols, scheme) {
		return url, nil
	}

	if ContainsString(netLocationProtocols, scheme) {
		if len(netLocation) > 0 {
			return joinURL(scheme, netLocation, path), nil
		}
		netLocation = baseNetLocation
	}

	if len(path) == 0 {
		path = basePath
		return joinURL(scheme, netLocation, path), nil
	}

	// to get path from baseURL and remove the last one
	baseParts := strings.Split(basePath, "/")
	if len(baseParts[len(baseParts)-1]) > 0 {
		baseParts = baseParts[:len(baseParts)-1]
	}

	// append path to path from baseURL
	var segments []string
	if path[:1] == "/" {
		segments = strings.Split(path, "/")
	} else {
		splitPath := strings.Split(path, "/")
		segments = baseParts
		segments = append(segments, splitPath...)
	}

	// to pop path if ".." occurs and ignore if "."  occurs
	resolvedPath := make([]string, 0)
	for _, v := range segments {
		if v == ".." {
			resolvedPath = resolvedPath[:len(resolvedPath)-1]
		} else if v == "." {
			continue
		} else {
			resolvedPath = append(resolvedPath, v)
		}
	}

	if segments[len(segments)-1] == "." || segments[len(segments)-1] == ".." {
		resolvedPath = append(resolvedPath, "")
	}

	path = strings.Join(resolvedPath, "/")
	return joinURL(scheme, netLocation, path), nil
}

// splitURL splits url into 3 components (scheme, netLocation, path)
// <scheme>://<netLocation>/<path>
func splitURL(url string) (scheme, netLocation, path string, err error) {
	// check and get scheme
	if strings.Contains(url, ":") {
		var posColon int
		for i, v := range url {
			// 58 rune for ":"
			if v == 58 {
				posColon = i
				break
			}
		}
		for _, v := range url[:posColon] {
			if !strings.Contains(schemeChars, string(v)) {
				break
			}
		}
		scheme = strings.ToLower(url[:posColon])
		url = url[posColon+1:]
	}

	// splitting between netLocation and path
	if url[:2] == "//" {
		netLocation, url = SplitNetLocation(url, 2)
		if (strings.Contains(netLocation, "[") && !strings.Contains(netLocation, "]")) ||
			(!strings.Contains(netLocation, "[") && strings.Contains(netLocation, "]")) {
			return "", "", "", errors.New("error when url split")
		}
	}

	// path is the rest of it
	if url == "/" {
		url = ""
	}

	return scheme, netLocation, url, nil
}

// SplitNetLocation is to get net location
func SplitNetLocation(url string, start int) (domain, rest string) {
	delim := len(url)
	c := "/?#"
	for _, v := range c {
		wdelim := strings.Index(url[2:], string(v))
		if wdelim >= 0 {
			if delim >= wdelim {
				delim = wdelim + 2
			}
		}
	}
	return url[start:delim], url[delim:]
}

// joinURL is to join the url from 3 components
func joinURL(scheme, netLocation, path string) (url string) {
	if len(scheme) > 0 {
		url = scheme + "://" + netLocation
	} else {
		url = netLocation
	}

	if len(path) > 0 {
		if string(path[0]) == "/" {
			url += path
			return
		}
	}

	if url != "" {
		url = url + "/" + path
		return
	}

	url = path
	return
}
