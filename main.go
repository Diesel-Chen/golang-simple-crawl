package main

import (
	"golang-simple-crawl/engine"
	"golang-simple-crawl/zhenai/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun/guangzhou",
		ParserFunc: parser.ParserCityList,
	})

}
