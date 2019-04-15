package main

// 并发 并行

import (
	"fmt"
	"runtime"
)


func test(){

}


func main(){
	go test()
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	fmt.Println(num)

}


