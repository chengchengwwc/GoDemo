package main

import (
	"fmt"
	"os"
	"flag"
)


func main(){
	//参数个数
	var confPath string
	var loglevel int
	fmt.Println(len(os.Args))
	for _,v := range os.Args{
		fmt.Println(v)
	}
	flag.StringVar(&confPath,"C"," ","please input conf path")
	flag.IntVar(&loglevel,"d",0,"please input log level")
	flag.Parse()
	fmt.Println("Path:",confPath)
	fmt.Println(loglevel)
}