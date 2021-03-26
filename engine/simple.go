package engine

import (
	"golang-simple-crawl/fetcher"
	"log"
	"strings"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seed ...Request) {
	requests := make([]Request, 0)
	for _, v := range seed {
		requests = append(requests, v)
	}
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]
		result := Work(request)
		requests = append(requests, result.Requests...)
		for _, item := range result.Items {
			log.Printf("got item : %s\n", item)
		}
	}
}
func Work(request Request) RequestResult {
	request.Url = strings.Replace(request.Url, "http://", "https://", 1)
	log.Printf("Fetch Url:%s", request.Url)
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetch  url: %s  error %s:", request.Url, err)
	}
	return request.ParserFunc(body)
}
