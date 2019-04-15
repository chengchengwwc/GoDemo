package main

import (
	"io"
	"bufio"
	"fmt"
	"os"
)


type ChartCount struct{
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}


func main(){
	file,err := os.Open("C:/test.log")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	var count ChartCount
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println(err)
			break
		}
		//数据分组
		runeArr := []rune(str)
		for _, v := range runeArr{
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount ++
			case v == ' '|| v == '\t':
				count.SpaceCount ++
			case v >= '0' && v <= '9':
				count.NumCount ++
			default:
				count.OtherCount++
		
			}
		} 
	}
	fmt.Println(count)
}