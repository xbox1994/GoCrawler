package main

import (
	"GoTest/crawler_concurrent/engine"
	"GoTest/crawler_concurrent/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
