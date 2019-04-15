package main

import (
	"time"
	"os"
	"fmt"
	"encoding/json"
)

type PathError struct{
	path string
	op string
	time_ string
	message string
}
//实现接口自定义错误
func (p *PathError)Error()string{
	return fmt.Sprintf("%s %s %s",p.path,p.op,p.message)

}



func OpenDemo(filename string) error{
	file,err := os.Open(filename)
	if err != nil{
		return &PathError{
			path:filename,
			op:"swr",
			message:err.Error(),
		}
	}
	defer file.Close()
	return nil
}





type User struct{
	Username string `json:username`
	NickName string `json:nickname`
	Age int
	Birthday string
	Sex string
	Email string
	Phone string
}

func testMap(){
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["user_name"]= "DDD"
	m["age"] = 12
	data,err := json.Marshal(m)
	if err != nil{
		fmt.Println(err)
		return 
	}
	fmt.Println(data)
}

func testSlice()(ret string,err error){
	var s []map[string]interface{}
	m := make(map[string]interface{})
	m["user_name"]= "DDD"
	m["age"] = 12

	s = append(s,m)
	data,err := json.Marshal(s)
	if err != nil{
		fmt.Println(err)
		return 
	}
	ret = string(data)
	return

}

func test(){
	var s []map[string]interface{}
	data,err := testSlice()
	if err != nil{
		fmt.Println("BAD")
		return 
	}
	err_ := json.Unmarshal([]byte(data),&s)
	fmt.Println(s)
	fmt.Println(err_)
	
}









func main(){
	user1 := &User{
		Username:"WEIHCNEG",
		NickName:"WANG",
		Age:20,
		Birthday:"2001",
		Sex:"MAN",
		Email:"dfdfd",
		Phone:"182323",
	}
    //json序列化
	data,err := json.Marshal(user1)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("%s\n",string(data))



}




