package parser

import (
	"golang-simple-crawl/crawl/engine"
	"regexp"
)

const regCityList = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.RequestResult {
	reg := regexp.MustCompile(regCityList)
	resp := reg.FindAllSubmatch(contents, -1)
	result := engine.RequestResult{}
	limit := 50
	for _, val := range resp {
		limit--
		if limit < 0 {
			break
		}
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(val[1]),
			ParserFunc: ParserCity,
		})
	}
	return result
}
