package engine

import (
	"log"
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
