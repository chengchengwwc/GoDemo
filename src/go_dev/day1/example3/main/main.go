package main

import (
	"bufio"
	"fmt"
)

type student struct{
	Name string
	Age int
	Score float32
}

var inputReader *bufio.ReadWriter
var input string
var err error



func main(){
	var str = "stu01 10 89.5"
	var stu student
	fmt.Sscanf(str,"%s %d %f",&stu.Name,&stu.Age,&stu.Score)

}