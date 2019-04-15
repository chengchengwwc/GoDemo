package main

import (
	"log"
	"io"
	"time"
	"net"
	"fmt"
	"io/ioutil"
	"net/http"
)

var url = []string{
	"abcd",
	"deff",
	"DDDD",
}


func tmp(){
	for _,v := range url{
		c := http.Client{
			Transport:&http.Transport{
				Dial:func(network,addr string)(net.Conn,error){
					timeout := time.Second*2
					return net.DialTimeout(network,addr,timeout)
				},
			},
		}
		resp,err := c.Head(v)
		if err != nil{
			fmt.Println(err)
			continue
		}
		fmt.Println(resp.Status)


	}
}

func FormServer(w http.ResponseWriter,request *http.Request){
	w.Header().Set("Content.Type","text/html")
	switch request.Method{
	case "GET":
		fmt.Println("D")
	case "POST":
		request.ParseForm()
		io.WriteString(w,request.Form["in"][0])
		io.WriteString(w,"\n")
		io.WriteString(w,request.FormValue("in"))
	}
}

//处理panic
func logPanics(handle http.HandleFunc) http.HandleFunc{
	return func(writer http.ResponseWriter,request *http.Request){
		defer func(){
			if x := recover();x!= nil{
				log.Printf("%v",x)
			}
		}()
		handle(writer,request)
		
	}
}







func main(){
	res,err := http.Get("https://www.baidu.com")
	if err != nil{
		fmt.Println(err)
		return
	}
	// byte
	data,err := ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println(err)
		return 
	}
	//byte to string
	fmt.Println(string(data))
}





