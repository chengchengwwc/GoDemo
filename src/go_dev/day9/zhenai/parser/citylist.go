package parser

import (
	"go_dev/day9/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenhun/[0-9a-z])+"[^>]*>([^<])+</a>`


func ParseCityList(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}
	limt := 10


	for _,m := range matches {
		name := string(m[2])
		result.Items = append(result.Items,string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url       :string(m[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes,name)
				
			},
		})
		limt --
		if limt == 0{
			break
		}

	}
	return result
}