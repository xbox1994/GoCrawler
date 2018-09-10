package main

import (
	"GoTest/crawler_standalone/engine"
	"GoTest/crawler_standalone/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
