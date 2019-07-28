package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"golang.org/x/tools/go/ssa/interp/testdata/src/time"
	"runtime"
)

func goTestDemo(){
	fmt.Println("DDD")
	var a [10]int
	for i:=0;i<10;i++{
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched() // 交出控制权
			}
		}(i)
	}
	time.Sleep(1)
	fmt.Println(a)


}