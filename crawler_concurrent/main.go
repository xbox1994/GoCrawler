package main

import (
	"GoTest/crawler_concurrent/engine"
	"GoTest/crawler_concurrent/parser"
	"GoTest/crawler_concurrent/scheduler"
)

func main() {
	engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
