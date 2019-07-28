package main

import (
	"go_dev/day9/rpc"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main(){
	rpc.Register(rpcdemo.DemoService{})
	listener,err := net.Listen("tcp",":1234")
	if(err != nil){
		panic(err)
	}
	for{
		conn,err := listener.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}

		go jsonrpc.ServeConn(conn)


	}


	}

