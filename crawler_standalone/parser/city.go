package parser

import (
	"GoTest/crawler_standalone/engine"
	"regexp"
)

const city = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(city)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, match := range matches {
		name := string(match[2])
		result.Items = append(result.Items, "User"+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(match[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})
	}
	return result
}
