package main

import (
	"os"
	"time"
	"fmt"
	"math/rand"
	"go_dev/day1/example2/balance"
)

func main(){
	var insts [] *balance.Instance
	for i :=0;i<10;i++{
		host := fmt.Sprintf("192.168.%d.%d",rand.Intn(255),rand.Intn(255))
		port := 22
		one := balance.NewInstance(host,port)
		insts = append(insts,one)
	}

	var conf = "random"
	if len(os.Args) >1 {
		conf = os.Args[1]
	}

	for {
		inst,err := balance.DoBalance(conf,insts)
		if err != nil{
			fmt.Println(err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)

	}

}