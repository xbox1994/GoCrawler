package main

import (
	"GoCrawler/crawler_concurrent/engine"
	"GoCrawler/crawler_concurrent/parser"
	"GoCrawler/crawler_concurrent/persist"
	"GoCrawler/crawler_concurrent/scheduler"
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
