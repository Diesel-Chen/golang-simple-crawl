package engine

import (
	"golang-simple-crawl/fetcher"
	"log"
)

func Run(seed ...Request){
	requests:=make([]Request,0)
	for _,v:=range seed{
		requests = append(requests,v)
	}
	for len(requests)>0{
		request:=requests[0]
		requests = requests[1:]
		body,err:=fetcher.Fetch(request.Url)
		if err!=nil{
			log.Printf("Fetch  url: %s  error %s:",request.Url,err)
			continue
		}
		result:=request.ParserFunc(body)
		requests = append(requests,result.Requests...)
		for _,item:=range result.Items{
			log.Printf("got item : %s\n",item)
		}
	}
}
