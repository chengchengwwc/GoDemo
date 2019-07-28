package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"sync"
)


type worker struct {
	in chan int
	wg *sync.WaitGroup
}






func work(id int,c chan int,wg *sync.WaitGroup){
	for n:= range c{
		fmt.Println(id,n)
		wg.Done()
	}
}


func CreateworkDemo(id int,wg *sync.WaitGroup)worker{
	w := worker{
		in:make(chan int),
		wg:wg,

	}
	go work(id,w.in,w.wg)
	return w

}



func chanDemo(){
	var workers [10]worker
	var wg  sync.WaitGroup
	wg.Add(10)

	for i :=0 ;i<10;i++{

		workers[i] = CreateworkDemo(i,&wg)
	}


	for i:=0;i<10;i++{
		workers[i].in <- i

	}
	//
	wg.Wait()




	}


func channelClose(){
	c := make(chan int,3)
	close(c)
}

















