package main

import "fmt"
import "net"


func runServer(addr string)(err error){
	l,err := net.Listen("tcp",addr)
	if err != nil{
		fmt.Println("listen failed:",err)
		return
	}
	for {
		conn,err := l.Accept()
		if err != nil{
			fmt.Println("Accept failed: ",err)
			continue
		}
		
	}
}


func process(conn net.Conn){
	defer conn.Close()
	client := &Client{
		conn:conn,
	}
	err := client.Pr
}

