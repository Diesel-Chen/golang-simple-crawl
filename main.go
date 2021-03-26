package main

import (
	"golang-simple-crawl/engine"
	"golang-simple-crawl/zhenai/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
	//c, err := fetcher.Fetch("https://album.zhenai.com/u/113482579")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(c))

}
