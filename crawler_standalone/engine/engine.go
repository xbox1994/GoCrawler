package engine

import (
	"GoTest/crawler_concurrent/fetcher"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
)

func Run(seeds ...Request) {
	var requestQueue []Request

	// 根据种子初始化队列
	for _, request := range seeds {
		requestQueue = append(requestQueue, request)
	}

	for len(requestQueue) > 0 {
		// 拿出队列中第一个元素
		request := requestQueue[0]
		requestQueue = requestQueue[1:]

		// 下载html内容
		println(request.Url)
		body, e := fetcher.Fetch(request.Url)
		if e != nil {
			log.Error("Fetcher error with url: %s. %v", request.Url, e)
			continue
		}

		// 根据request里提供的解析方法解析当前html内容成为result放入队列
		parseResult := request.ParserFunc(body)
		requestQueue = append(requestQueue, parseResult.Requests...)

		// 现在是打印，以后是保存
		for _, item := range parseResult.Items {
			fmt.Printf("Got item %v\n", item)
		}
	}
}
