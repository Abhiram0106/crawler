package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {

	response, getError := http.Get(rawURL)

	if getError != nil {
		return "", getError
	}

	contentType := response.Header.Get("Content-Type")
	if strings.Split(contentType, ";")[0] != "text/html" {
		return "", errors.New("response header not text/html")
	}

	defer response.Body.Close()
	htmlBytes, readError := io.ReadAll(response.Body)

	if readError != nil {
		return "", readError
	}

	return string(htmlBytes), nil
}
