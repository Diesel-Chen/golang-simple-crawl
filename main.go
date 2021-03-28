package main

import (
	"golang-simple-crawl/engine"
	"golang-simple-crawl/persist"
	"golang-simple-crawl/scheduler"
	"golang-simple-crawl/zhenai/parser"
)

func main() {

	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParserCityList,
	})

}
