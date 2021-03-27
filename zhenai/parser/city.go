package parser

import (
	"golang-simple-crawl/engine"
	"regexp"
)

const regCity = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParserCity(contents []byte) engine.RequestResult {
	reg := regexp.MustCompile(regCity)
	resp := reg.FindAllSubmatch(contents, -1)
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
	return result
}
