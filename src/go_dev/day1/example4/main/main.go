package main


import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	str,err := reader.ReadString('\n')
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(str)





}
