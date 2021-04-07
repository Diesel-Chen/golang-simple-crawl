package main

import (
	"golang-simple-crawl/crawl/engine"
	"golang-simple-crawl/distribute-crawl/persist/client"

	"golang-simple-crawl/crawl/scheduler"
	"golang-simple-crawl/crawl/zhenai/parser"
)

func main() {

	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    client.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParserCityList,
	})

}
