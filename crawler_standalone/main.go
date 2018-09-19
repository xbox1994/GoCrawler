package main

import (
	"GoCrawler/crawler_standalone/engine"
	"GoCrawler/crawler_standalone/parser"
)

func main() {
	engine.SimpleEngine{
		WorkerCount: 100,
	}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
