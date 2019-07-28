package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func HelloServer(w http.ResponseWriter,req *http.Request){
	io.WriteString(w,"Hello World")
}

func demo_test(){
	http.HandleFunc("/hello",HelloServer)
	http.ListenAndServe(":8000",nil)
}


func get_demo(){
	resp,err := http.Get("http://www.baidu.com")
	if err != nil{
		fmt.Println(err)
		return
	}
	s,err := httputil.DumpResponse(resp,true)
	if err != nil{
		fmt.Println("err")
		return
	}
	fmt.Println(s)

}

func get_demo2(){
	//替换头部信息

	//net/http/ppof可以查看基本性能
	//标准库
	// bufio
	//log
	//encoding/json
	//regexg
	//time
	//strings/math/rand
	// 查看标准库文档
	//godoc -http:8888
	//

	request,err :=http.NewRequest(http.MethodGet,"https://www.baidu.com",nil)
	if err != nil{
		fmt.Println(err)
		return
	}
	request.Header.Add("User-Agent","Mozilila/5.0")
	resp,err := http.DefaultClient.Do(request)
	if err != nil{
		fmt.Println(err)
		return
	}
	s,err := httputil.DumpResponse(resp,true)
	if err != nil{
		fmt.Println("err")
		return
	}
	fmt.Println(s)




}





















