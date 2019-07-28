package parser

import (
	"go_dev/day9/engine"
	"go_dev/day9/model"
	"regexp"
	"strconv"
)

var  ageRe       = regexp.MustCompile(`<td><span class="label">年龄:</span>([\d]+)</td>`)
var  marriageRe  = regexp.MustCompile(`<td><span class="label">婚况:</span>([\d]+)</td>`)




func extractString(contents []byte,re *regexp.Regexp) string{
	math := re.FindSubmatch(contents)
	if len(math) >= 2{
		return string(math[1])
	}else{
		return  ""
	}

}




func ParseProfile(contents []byte,name string) engine.ParseResult{
	profile := model.Profile{}
	age, err := strconv.Atoi(extractString(contents,ageRe))
	if err == nil{
		profile.Age =age
	}
	profile.Marriage = extractString(contents,marriageRe)

    profile.Name = name
	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}


	return  result
	}
