package main

import (
	"go_dev/day9/engine"
	"go_dev/day9/scheduler"
)

import "go_dev/day9/zhenai/parser"

func main(){
	 m := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
	    WorkerCount:10}
	 m.Run()

	 engine.SimpleEngine{}.Run(engine.Request{
		Url:"https://www.zhenai.com/zhenghun/",
		ParserFunc:parser.ParseCityList,
	})

}


