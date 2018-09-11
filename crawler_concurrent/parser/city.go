package parser

import (
	"GoTest/crawler_concurrent/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	userMatches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, match := range userMatches {
		name := string(match[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(match[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})
	}

	cityMatches := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, match := range cityMatches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(match[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
