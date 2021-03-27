package main

import (
	"golang-simple-crawl/engine"
	"golang-simple-crawl/scheduler"
	"golang-simple-crawl/zhenai/parser"
)

func main() {

	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 30,
	}
	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParserCityList,
	})

}
