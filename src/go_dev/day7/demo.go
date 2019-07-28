package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)














func HttpDemo(){
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/list/"):]
		file ,err := os.Open(path)
		if err != nil{
			http.Error(writer,err.Error(),http.StatusInternalServerError)
			return
		}
		defer file.Close()
		all ,err := ioutil.ReadAll(file)
		if err != nil{
			panic(err)
		}
		writer.Write(all)

	})
	err := http.ListenAndServe(":8888",nil)
	if(err != nil){
		panic(err)
	}
}






func writeFile(filename string){
	file,err := os.Create(filename)
	if err != nil{
		//错误处理
		if pathError,ok := err.(*os.PathError);ok{
			fmt.Println(pathError.Err)
		}
		panic(err)
	}

	defer file.Close()


	writer := bufio.NewWriter(file)
	for i:=0;i<20;i++{
		fmt.Fprintln(writer,10)
	}
	writer.Flush()

	}









func main(){
	writeFile("abc.txt")
	fmt.Println("DDD")
}
