package parse

import (
	"Projectlearn/LeanCrawler/engine"
	"regexp"

	"fmt"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9A-Za-z]+)"[^>]*>([^<]+)</a>`

func ParseCity(content []byte) engine.PaseResult{
   res :=  regexp.MustCompile(cityListRe)
   matcher := res.FindAllSubmatch(content,-1)

   //limit := 2
   reslut := engine.PaseResult{}

   for _,m := range matcher{
   	  fmt.Printf("ciyt:%v,url:%v",string(m[2]),string(m[1]))
   	  reslut.Requests =append(reslut.Requests,engine.Request{
   	  	 Url:string(m[1]),
   	  	 ParserFunc:ParseCityUser,
	  })
   	  reslut.Items=append(reslut.Items,"city:"+string(m[2]))

   	 // limit--
   	  //if limit == 0{
   	  	//break
	  //}
   	}

   return reslut;

}