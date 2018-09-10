package engine

import (
	"GoTest/crawler_concurrent/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (SimpleEngine) Run(seeds ...Request) {
	var requestQueue []Request

	// 根据种子初始化队列
	for _, request := range seeds {
		requestQueue = append(requestQueue, request)
	}

	for len(requestQueue) > 0 {
		// 拿出队列中第一个元素
		request := requestQueue[0]
		requestQueue = requestQueue[1:]

		parseResult, err := worker(request)
		if err != nil {
			continue
		}
		requestQueue = append(requestQueue, parseResult.Requests...)

		// 现在是打印，以后是保存
		for _, item := range parseResult.Items {
			log.Printf("Got item %v\n", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	// 下载html内容
	body, e := fetcher.Fetch(r.Url)
	if e != nil {
		log.Printf("Fetcher error with url: %s. %v", r.Url, e)
		return ParseResult{}, e
	}

	// 根据request里提供的解析方法解析当前html内容成为result放入队列
	return r.ParserFunc(body), nil
}
