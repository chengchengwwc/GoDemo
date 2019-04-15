package main

import(
	"fmt"
	"os"
	"text/template"
)

type Person struct{
	Name string
	Age string
}

func main(){
	t,err := template.ParseFiles("./index.html")
	if err != nil{
		fmt.Println("Parse file err:",err)
		return 
	}
	p := Person{Name:"DDD",Age:"DDDD"}
	if err := t.Execute(os.Stdout,p);err != nil{
		fmt.Println("There was an error:",err.Error())
	}


}