package main

import (

	"Projectlearn/LeanCrawler/engine"

	"Projectlearn/LeanCrawler/zhenai/parse"
	"Projectlearn/LeanCrawler/scheduler"
)

func main(){

  e := engine.ConcurrentEngine{
  	Scheduler:&scheduler.SimpleScheduler{},
  	WorkerCount:10,
  }

  e.Run(engine.Request{
	  "http://www.zhenai.com/zhenghun/aba",
	  parse.ParseCityUser,
  })

}


