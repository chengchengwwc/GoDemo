package main

import "fmt"
import "go_dev/day3/proto"
import "net"


var userId int
var paswd string
var msgChan chan proto.UserRecvMessageReq

func init(){
	msgChan = make(chan proto.UserRecvMessageReq,10000)
}

func main(){
	fmt.Scanf("%d %s\n",&userId,&paswd)
	conn,err := net.Dial("tcp","localhost:8000")
	if err != nil{
		fmt.Println("Error dialing ",err.Error())
		return
	}
	err = login(conn,userId,paswd)
	if err != nil{
		fmt.Println("login failed err: ",err)
		return
	}
	go processServerMessage(conn)
	for {
		logic(conn)
	}
}



