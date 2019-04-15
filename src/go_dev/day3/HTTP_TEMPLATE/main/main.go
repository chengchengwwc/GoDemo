package main

import (
	"os"
	"fmt"
	"net/http"
	"html/template"
)

type Result struct{
	output string
}



var myTemplate *template.Template

type Person struct {
	Name string
	Titel string
	Age string
}

func (p *Result) Write(b []byte)(n int,err error){
	fmt.Println("called by template")
	p.output += string(b)
	return len(b),nil
}





func userInfo(w http.ResponseWriter, r *http.Request){
	fmt.Println("handle hello")
	fmt.Fprintf(w,"hello")
	p := Person{
		Name:"DDD",
		Titel:"DDD",
		Age:"SSS",
	}
	resultwrite := &Result{}
	myTemplate.Execute(resultwrite,p)


	//模板渲染
	myTemplate.Execute(w,p)
	file,err := os.OpenFile("./tmp.txt",os.O_CREATE|os.O_WRONLY,0755)
	if err != nil{
		fmt.Println(err)
		return
	}
	//写入文件
	myTemplate.Execute(file,p)
}



func initTemplate(filenames string) ( err error){
	myTemplate,err = template.ParseFiles(filenames)
	if err != nil{
		fmt.Println(err)
		return 
	}
	return 



}




func main(){
	initTemplate("DDDD")
	http.HandleFunc("/user/info",userInfo)
	err := http.ListenAndServe("0.0.0.0:8080",nil)
	if err != nil{
		fmt.Println(err)
	}

}


