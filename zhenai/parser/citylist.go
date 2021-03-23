package parser

import (
	"golang-simple-crawl/engine"
	"regexp"
)

const regCityList = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.RequestResult {
	reg := regexp.MustCompile(regCityList)
	resp := reg.FindAllSubmatch(contents, -1)
	result := engine.RequestResult{}
	for _, val := range resp {
		result.Items = append(result.Items, "cityName "+string(val[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(val[1]),
			ParserFunc: ParserCity,
		})
	}
	return result
}
