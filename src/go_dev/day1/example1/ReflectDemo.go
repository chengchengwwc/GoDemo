package example1


import (
	"fmt"
	"reflect"
)

func test(b interface{}){
	var a := reflect.TypeOf(b)
	fmt.Println(a)
	var c := reflect.ValueOf(b)
	d := c.Kind()
	fmt.Println(d)
	iv :=  c.Interace()
	fmt.Println(iv)
	
}


