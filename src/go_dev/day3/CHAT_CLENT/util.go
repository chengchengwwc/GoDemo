package main

import "encoding/binary"
import "encoding/json"
import "errors"
import "fmt"
import "go_dev/day3/proto"
import "net"



func readPackage(conn net.Conn)(msg proto.Message,err error){
	var buf [8192]byte
	n,err := conn.Read(buf[0:4])
	if n != 4{
		err = errors.New("read header failed")
		return
	}
	var packLen uint32
	packLen = binary.BigEndian.Uint32(buf[0:4])
	n,err = conn.Read(buf[0:packLen])
	if n != int(packLen){
		err = errors.New("read body failed")
		return
	}
	err = json.Unmarshal(buf[0:packLen],&msg)
	if err != nil{
		fmt.Println("unmarshal failed,err:",err)
	}
	return
}