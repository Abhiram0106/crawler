package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {

	parsedUrl, err := url.Parse(inputURL)

	if err != nil {
		return "", nil
	}

	return parsedUrl.Host + strings.TrimSuffix(parsedUrl.Path, "/"), nil
}
