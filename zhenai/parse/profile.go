package parse

import (
	"Projectlearn/LeanCrawler/engine"
	"regexp"
	"Projectlearn/LeanCrawler/model"
	"strconv"
)


var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var marry  = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)



func ParseProfile(contents []byte) engine.PaseResult{
	profile := model.Profile{}
	age,err := strconv.Atoi(extractString(contents,ageRe))
	if err == nil{
		profile.Age = age
	}
	height,err := strconv.Atoi(extractString(contents,heightRe))
	if err == nil{
		profile.Height = height
	}

	profile.Marriage =extractString(contents,marry)

	reslt := engine.PaseResult{Items:[]interface{}{profile}}

	return reslt

}

func extractString(contents []byte,re *regexp.Regexp) string{
	match := re.FindSubmatch(contents)
	if len(match) >= 2{
		return  string(match[1])
	}else{
		return  ""
	}
}
