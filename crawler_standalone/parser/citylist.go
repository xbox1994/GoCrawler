package parser

import (
	"GoTest/crawler_concurrent/engine"
	"regexp"
)

const cityList = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityList)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, match := range matches {
		result.Items = append(result.Items, "City"+string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(match[1]),
			ParserFunc: ParseCity,
		})
		break
	}
	return result
}
