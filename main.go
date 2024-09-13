package main

import (
	"fmt"
	"log"
)

func main() {
	urlsMap := make(map[string]int)
	crawlPage("https://wagslane.dev/", "https://wagslane.dev/", &urlsMap)

	log.Println("Found the following URLs")
	for url, count := range urlsMap {
		fmt.Printf("[%d] %s\n", count, url)
	}
}
