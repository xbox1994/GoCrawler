package engine

import (
	"GoCrawler/crawler_distributed/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	// 下载html内容
	body, e := fetcher.Fetch(r.Url)
	if e != nil {
		log.Printf("Fetcher error with url: %s. %v", r.Url, e)
		return ParseResult{}, e
	}

	// 根据request里提供的解析方法解析当前html内容成为result放入队列
	return r.ParserFunc(body, r.Url), nil
}
