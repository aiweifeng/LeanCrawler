package engine

import (
	"log"
	"Projectlearn/LeanCrawler/fetch"

)

func Run(seed... Request){
	var requests []Request
	for _,r:=range seed{
		requests =append(requests,r)
	}

	for len(requests) > 0{
		r :=requests[0]
		requests =requests[1:]

		parseResult,err:= worker(r)
		if err!=nil{
			continue
		}

		requests=append(requests,parseResult.Requests...)

		//对应结果Items 循环出来
		for _,item := range parseResult.Items{
			log.Printf("go get item %v",item)
		}



	}
}

func worker(r Request) (PaseResult,error)  {
	log.Printf("fetching %s",r.Url)

	//调用fetch 去获取网页内容
	body,err := fetch.Fetch(r.Url)
	if err !=nil{
		log.Printf("fetcher:error"+"fetcher.url %s,%v",r.Url,err)
		return PaseResult{},err
	}
	//把requests加入 队列里面
	return r.ParserFunc(body),err
}

