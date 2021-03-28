package main

import (
	"golang-simple-crawl/engine"
	"golang-simple-crawl/persist"
	"golang-simple-crawl/scheduler"
	parser2 "golang-simple-crawl/tupian/parser"
)

func main() {

	e := &engine.ConcurrentEngine{
		WorkerCount: 20ƒ,
		Scheduler:   &scheduler.QueuedScheduler{},
		ItemChan:    persist.PictureSaver(),
	}
	//e := &engine.SimpleEngine{}
	e.Run(engine.Request{
		Url:        "http://335pai.com/toukuizipai/index.html",
		ParserFunc: parser2.ParserList,
	})

}
