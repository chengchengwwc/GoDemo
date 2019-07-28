package engine

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}


type Scheduler interface {
	Submit(request Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}




func (e *ConcurrentEngine) Run(seeds ...Request){
	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}

	out := make(chan ParseResult)

	//e.Scheduler.ConfigureMasterWorkerChan(in)

	for i:=0;i<e.WorkerCount;i++{
		createWorker(out,e.Scheduler)
	}

	for{
		result := <- out
		for _,item := range result.Items{
			fmt.Printf("Got item:%v",item)
		}
		for _,request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
	}

func createWorker(out chan ParseResult,s Scheduler){
	in := make(chan Request)
	go func() {
		for {
			// tell scheduler  i am ready
			s.WorkerReady(in)
			request := <- in
			result,err := worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}