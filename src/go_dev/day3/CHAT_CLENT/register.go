package main

import "encoding/binary"
import "encoding/json"
import "fmt"
import "go_dev/day3/proto"
import "net"

func register(conn net.Conn,userId int,paswd string)(err error){
	var msg proto.Message
	msg.Cmd = proto.UserRegister
	var registerCmd proto.RegisterCmd
	registerCmd.User.UserId = userId
	registerCmd.User.Nick = fmt.Sprintf("stu%d",userId)
	registerCmd.User.Sex = "man"
	registerCmd.User.Passwd = paswd
	registerCmd.User.Header = "https://baidu.com/header/1.jpg"
	data,err := json.Marshal(registerCmd)
	if err != nil{
		return 
	}
	msg.Data = string(data)
	data,err = json.Marshal(msg)
	if err != nil{
		return
	}
	var buf [4]byte
	packLen := uint32(len(data))
	
	binary.BigEndian.PutUint32(buf[0:4],packLen)
	n,err := conn.Write(buf[:])
	if err != nil || n != 4{
		fmt.Println("write data failed")
		return
	}
	_,err = conn.Write([]byte(data))
	if err != nil{
		return
	}
	msg,err = readPackage(conn)
	if err != nil{
		fmt.Println("read pacakge failed err:",err)
	}
	fmt.Println(msg)
	return
}