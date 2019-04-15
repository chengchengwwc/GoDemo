package example1

import (
	"fmt"
)


type test interface{
	GetName() string
	Run()

}

type BMW struct{
	Name string
}


func (p *BMW) GetName() string{
	return p.Name
}

func(p *BMW)Run(){
	fmt.Println(p.Name)
}

func main(){
	var car test
	var bmw BMW
	//指针类型实现的接口
	car = &bmw
	
	
}