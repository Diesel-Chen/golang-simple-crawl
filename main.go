package main

import (
	"golang-simple-crawl/engine"
	"golang-simple-crawl/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

}
