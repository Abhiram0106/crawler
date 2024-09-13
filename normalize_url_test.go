package main

import (
	"log"
	"testing"
)

func TestNormalizeUrl(t *testing.T) {
	expectedURL := "example.com"
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "no change",
			inputURL: expectedURL,
			expected: expectedURL,
		},
		{
			name:     "remove scheme s",
			inputURL: "https://example.com/",
			expected: expectedURL,
		},
		{
			name:     "remove scheme",
			inputURL: "http://example.com/path",
			expected: expectedURL + "/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "example.com/path/",
			expected: expectedURL + "/path",
		},
		{
			name:     "remove scheme s and trailing slash",
			inputURL: "https://example.com/path/",
			expected: expectedURL + "/path",
		},
		{
			name:     "remove scheme and trailing slash",
			inputURL: "http://example.com/path/",
			expected: expectedURL + "/path",
		},
		{
			name:     "remove scheme s and but not trailing http",
			inputURL: "https://example.com/path/http",
			expected: expectedURL + "/path/http",
		},
		{
			name:     "remove scheme and trailing slash but not trailing http",
			inputURL: "http://example.com/path/http/",
			expected: expectedURL + "/path/http",
		},
		{
			name:     "remove scheme and trailing slash but not trailing https",
			inputURL: "http://example.com/path/https/",
			expected: expectedURL + "/path/https",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualPath, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			log.Println(actualPath)
			if actualPath != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actualPath)
			}
		})
	}
}
