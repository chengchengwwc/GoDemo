package main


//服务端处理流程：监听端口 接受客户端的链接 创建goroutine 处理链接

import (
	"fmt"
	"net"
)


func main(){
	fmt.Println("start server")
	listen,err := net.Listen("tcp","0.0.0.0:50000")
	if err != nil{
		fmt.Println("lisen failed")
		return 
	}
	for {
		conn,err := listen.Accept()
		if err != nil{
			fmt.Println("accept failed err")
			continue
		}
		go process(conn)

	}

}


func process(conn net.Conn){
	defer conn.Close()
	for {
		buf := make([]byte,512)
		_,err := conn.Read(buf)
		if err != nil{
			fmt.Println("read error")
			return
		}
		fmt.Println(string(buf))



	}
}










