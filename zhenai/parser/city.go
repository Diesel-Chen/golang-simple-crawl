package parser

import (
	"golang-simple-crawl/engine"
	"regexp"
)

var regCity = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

var moreCity = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^/]+/\d)">`)

func ParserCity(contents []byte) engine.RequestResult {
	resp := regCity.FindAllSubmatch(contents, -1)
	result := engine.RequestResult{}
	for _, val := range resp {
		name := string(val[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(val[1]),
			ParserFunc: func(bytes []byte) engine.RequestResult {
				return ParserProfile(name, bytes)
			},
		})
	}
	resp = moreCity.FindAllSubmatch(contents, -1)
	for _, v := range resp {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(v[1]),
			ParserFunc: ParserCity,
		})
	}

	return result
}
