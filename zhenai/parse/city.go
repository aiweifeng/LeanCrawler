package parse

import (
	"Projectlearn/LeanCrawler/engine"
	"regexp"
)

const city  = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

const next  = `href="(http://www.zhenai.com/zhenghun/[^"]+)"`



func ParseCityUser(content []byte) engine.PaseResult{

	res := regexp.MustCompile(city)
	matcher := res.FindAllSubmatch(content,-1);
	result := engine.PaseResult{}

	for _,m :=range matcher{
		name := string(m[2])
		result.Items=append(result.Items,"user"+name)
		result.Requests=append(result.Requests,
			engine.Request{Url:string(m[1]),
			ParserFunc:ParseProfile,

		})
	}

	res = regexp.MustCompile(next)
	matcher = res.FindAllSubmatch(content,-1);
	for _,m :=range matcher{
		result.Requests=append(result.Requests,
			engine.Request{Url:string(m[1]),
				ParserFunc:ParseCityUser,

			})
	}


   return result
}

