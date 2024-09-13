package main

import (
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {

	reader := strings.NewReader(htmlBody)
	tokenizer := html.NewTokenizer(reader)

	urls := []string{}

	// var helper func(tokenizer *html.Tokenizer)
	//
	// helper = func(tokenizer *html.Tokenizer) {
	// 	token := tokenizer.Next()
	// 	switch token {
	// 	case html.ErrorToken:
	// 		return
	// 	case html.StartTagToken, html.EndTagToken:
	// 		tagName, hasAttr := tokenizer.TagName()
	// 		if len(tagName) == 1 && tagName[0] == 'a' && hasAttr {
	// 			_, urlBytes, _ := tokenizer.TagAttr()
	// 			if urlBytes[0] != '/' && urlBytes[0] != 'h' {
	// 				break
	// 			}
	//
	// 			urlString := string(urlBytes)
	// 			println("urlBytes= " + urlString)
	// 			if urlBytes[0] == '/' {
	// 				urlString = rawBaseURL + urlString
	// 			}
	// 			urls = append(urls, urlString)
	// 		}
	// 	}
	// 	helper(tokenizer)
	// }
	// helper(tokenizer)

	// return urls, nil

	for {
		token := tokenizer.Next()
		switch token {
		case html.ErrorToken:
			return urls, nil
		case html.StartTagToken, html.EndTagToken:
			tagName, hasAttr := tokenizer.TagName()
			if len(tagName) == 1 && tagName[0] == 'a' && hasAttr {
				for hasAttr {
					attrKey, attrVal, moreAttr := tokenizer.TagAttr()
					if string(attrKey) == "href" {
						urlString := string(attrVal)
						if attrVal[0] == '/' {
							if rawBaseURL[len(rawBaseURL)-1] == '/' {
								rawBaseURL = strings.TrimSuffix(rawBaseURL, "/")
							}
							urlString = rawBaseURL + urlString
						}
						urls = append(urls, urlString)
						break
					}
					hasAttr = moreAttr
				}

			}
		}
	}
}
