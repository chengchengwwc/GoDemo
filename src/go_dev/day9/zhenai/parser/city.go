package parser

import (
	"go_dev/day9/engine"
	"regexp"
)

const cityRe = `<a href="<http://album.zhenai.com/u/[0-9]"`

func ParseCity(constent []byte) engine.ParseResult{
	re := regexp.MustCompile(cityRe)
	mathces := re.FindAllSubmatch(constent,-1)
	result := engine.ParseResult{}
	for _,m := range mathces{
		result.Items = append(result.Items,"User "+ string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			ParserFunc:nil,
		})
	}
	return result
}
