package parser

import (
	"golang-simple-crawl/crawl/engine"
	"regexp"
	"strings"
)

var regCity = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

var moreCity = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^/]+/\d)">`)

func ParserCity(contents []byte) engine.RequestResult {
	resp := regCity.FindAllSubmatch(contents, -1)
	result := engine.RequestResult{}
	for _, val := range resp {
		sp := strings.Split(string(val[1]), "/")
		id := sp[len(sp)-1]
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(val[1]),
			ParserFunc: ProfileParser(id, string(val[2]), string(val[1])),
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

func ProfileParser(id, name, url string) func([]byte) engine.RequestResult {
	return func(bytes []byte) engine.RequestResult {
		return ParserProfile(id, url, name, bytes)
	}
}
