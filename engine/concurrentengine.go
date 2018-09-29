package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	 ReadyNotifilter
	 //传入request方法
	 Submit(Request)

	 Run()
	 WorkerChan() chan Request

}

type ReadyNotifilter interface {
	//传入线程对应的接收request 的channel
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seed... Request){

	  out := make(chan PaseResult)

      e.Scheduler.Run()


	//启动worker
	for i:=0;i<e.WorkerCount;i++{
		createWork(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}

	 //把resquest放入到channel里面
     for _,r := range  seed{
		 if isDuplicate(r.Url){
			 continue
		 }

		 e.Scheduler.Submit(r)
	 }


	 itemCount :=0
	 //循环遍历结果
	 for{
	 	result := <-out

	 	for _,item := range result.Items{
	 		log.Printf("Got item:#%d,%v\n",itemCount,item)
	 		itemCount++
		}

		for _,request := range result.Requests{

			if isDuplicate(request.Url){
				continue
			}

			e.Scheduler.Submit(request)
		}
	 }
	 
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool{
	if visitedUrls[url]{
		return  true
	}
	visitedUrls[url] = true

	return  false
}


/**
启动工作任务
 */
func createWork(in chan Request,out chan PaseResult,s ReadyNotifilter)  {
	go func() {
		for{
			s.WorkerReady(in)
			request := <- in
			result,err :=  worker(request)
			if err!=nil{
				continue
			}
			out <-result
		}
	}()

}
