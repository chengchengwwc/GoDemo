package main

import (
	"go_dev/day9/rpc"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main(){
	conn ,err := net.Dial("tcp",":1234")
	if err != nil{
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div",rpcdemo.Args{10,3},&result)
	fmt.Println(err)
	fmt.Printf(result)
}
