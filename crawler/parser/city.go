package parser

import (
	"GoTest/crawler/engine"
	"regexp"
)

const city = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(city)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, match := range matches {
		result.Items = append(result.Items, "User" + string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(match[1]),
			ParserFunc: engine.NilParserFunc,
		})
	}
	return result
}
