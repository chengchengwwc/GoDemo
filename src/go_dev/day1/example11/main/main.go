package main

import (
	"fmt"
	"time"
	"sync"
)




var (
  m = make(map[int]uint64)
  lock sync.Mutex

)



type task struct {
	n int
}

func calc(t *task){
	var sum uint64
	for i :=1;i<t.n;i++{
		sum += uint64(i)

	}
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()

}


func main(){
	for i:=0;i<1000;i++{
		t := &task{n:1}
		go calc(t)
	}
	time.Sleep(10*time.Second)
	lock.Lock()
	for k,v := range m{
		fmt.Println(k)
		fmt.Println(v)
	}
	lock.Unlock()






}
