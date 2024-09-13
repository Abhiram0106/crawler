package main

import (
	"log"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages *map[string]int) {

	if !strings.Contains(rawCurrentURL, rawBaseURL) {
		return
	}

	normalizedCurrentURL, normalizeRawCurrentURLError := normalizeURL(rawCurrentURL)

	if normalizeRawCurrentURLError != nil {
		log.Fatalln("Failed to normalize raw current url " + rawCurrentURL)
	}

	if _, exists := (*pages)[normalizedCurrentURL]; exists {
		(*pages)[normalizedCurrentURL]++
		return
	} else {
		(*pages)[normalizedCurrentURL] = 1
	}

	log.Println("Crawling " + rawCurrentURL)
	htmlBody, getHTMLError := getHTML(rawCurrentURL)

	if getHTMLError != nil {
		log.Println("getHTMLError = " + getHTMLError.Error())
		return
	}

	log.Println("GettingURLsFromHTML...")
	rawAbsoluteURLs, getURLsError := getURLsFromHTML(htmlBody, rawBaseURL)

	if getURLsError != nil {
		log.Println("getURLsError = " + getURLsError.Error())
		return
	}

	for _, rawURL := range rawAbsoluteURLs {
		crawlPage(rawBaseURL, rawURL, pages)
	}
}
