package main

import "encoding/json"
import "fmt"
import "go_dev/day3/proto"
import "net"
import "os"


func processServerMessage(conn net.Conn){
	for{
		msg,err := readPackage(conn)
		if err != nil{
			fmt.Println("read err :",err)
			os.Exit(0)
		}
		var userStatus proto.UserStatusNotify
		err = json.Unmarshal([]byte(msg.Data),&userStatus)
		if err != nil{
			fmt.Println("unmarshal failedm,err:",err)
			return
		}
		switch msg.Cmd{
		case proto.UserStatusNotifyRes:
			updateUserStatus(userStatus)
		case proto.UserRecvMessageCmd:
			recvMessageFromServer(msg)
		}


	}
}


func recvMessageFromServer(msg proto.Message){
	var recvMsg proto.UserRecvMessageReq
	err := json.Unmarshal([]byte(msg.Data),&recvMsg)
	if err != nil{
		fmt.Println("unmarshal failed err:",err)
		return
	}
	fmt.Println("%d:%s\n",recvMsg.UserId,recvMsg.Data)
}
