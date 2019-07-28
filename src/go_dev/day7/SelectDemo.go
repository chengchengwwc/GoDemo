package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"golang.org/x/tools/go/ssa/interp/testdata/src/time"
	"math/rand"
	"sync"

)



func generator() chan int{
	out := make(chan int)
	go func() {
		i := 0
		for{
			time.Sleep(time.Duration(rand.Intn(1500))*1)
			out <- i
			i ++
		}
	}()
	return out
}


func worker1(id int,c chan int){
	for n:= range c{
		time.Sleep(time.Second)
		fmt.Println(id,n)
	}
}

func createWorker(id int) chan <- int{
	c := make(chan int)
	go worker1(id,c)
	return  c
}

func start_select(){
	var c1,c2 =generator(),generator()
	var woker = createWorker(0)
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan <- int
		var activeValue int
		if len(values) > 0{
			activeWorker = woker
			activeValue = values[0]
		}
		select{
		case n :=<-c1:
			values = append(values,n)
		case n := <-c2:
			values = append(values,n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len =", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}














type atomicInt struct {
	value int
	lock sync.Mutex //
}


func (a *atomicInt) increment(){
	defer a.lock.Unlock()

	a.lock.Lock()
	a.value++


}

func (a *atomicInt) get() int{
	defer a.lock.Unlock()

	a.lock.Lock()
	return a.value
}



func atomicDemo(){

}

























