package main

import "bufio"
import "fmt"
import "os"


func main(){
	
	file,err := os.Open("C:/test.log")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	str,err := reader.ReadString('\n')
	if err != nil{
		fmt.Println(err)
		return 
	}
	fmt.Println(str)
}
