package example1

import (
	"fmt"
	"reflect"
)


type Student struct{
	Name string
	Age int
	Score float32
}

func (s *Student) Set(name string,age int, score float32){
	s.Name = name
	s.Age = age
	s.Score = score
}




func TestStruct(a interface{}){
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct{
		fmt.Println("expect struct")
		return 
	}

	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct{
		fmt.Printf("AAAA")
	}



	num := val.NumField()
	fmt.Println(num)
	for i:=0;i<num;i++{
		fmt.Println(val.Field(i))
	}
	nummethod := val.NumMethod()


}




func main(){
	var a Student = Student{
		Name:"Stu01",
		Age:18,
		Score:98.5
	}

}
