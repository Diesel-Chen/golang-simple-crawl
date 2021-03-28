package parser

import (
	"golang-simple-crawl/engine"
	"regexp"
)

var listReg = regexp.MustCompile(`<li><a href="(/htm/[^"]+)".+<h3>([^<]+)</h3></a></li>`)

func ParserList(cnt []byte) engine.RequestResult {
	resp := listReg.FindAllSubmatch(cnt, -1)
	result := engine.RequestResult{}
	for _, v := range resp {
		name := string(v[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: "http://335pai.com" + string(v[1]),
			ParserFunc: func(cnt []byte) engine.RequestResult {
				return ParserPerson(cnt, name)
			},
		})
	}
	return result
}
