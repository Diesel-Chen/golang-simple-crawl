package engine

import (
	"golang-simple-crawl/fetcher"
	"log"
	"strings"
)

func Work(request Request) RequestResult {
	request.Url = strings.Replace(request.Url, "http://", "https://", 1)
	log.Printf("Fetch Url:%s", request.Url)
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetch  url: %s  error %s:", request.Url, err)
	}
	return request.ParserFunc(body)
}
