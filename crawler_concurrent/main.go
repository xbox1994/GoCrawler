package main

import (
	"GoTest/crawler_concurrent/engine"
	"GoTest/crawler_concurrent/parser"
	"GoTest/crawler_concurrent/persist"
	"GoTest/crawler_concurrent/scheduler"
)

func main() {
	itemChan, e := persist.ItemSaver("dating_profile")
	if e != nil {
		panic(e)
	}

	engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.QueuedScheduler{},
	//	WorkerCount: 10,
	//	ItemChan:    persist.ItemSaver(),
	//}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
