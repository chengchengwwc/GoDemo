package main

import (
	"io"
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func CopyFile(dstName, srcName string)(written int64,err error){
	//文件备份
	src,err := os.Open(srcName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer src.Close()
	dst,err := os.OpenFile(dstName,os.O_WRONLY|os.O_CREATE,0644)
	if err != nil{
		return
	}
	defer dst.Close()
	return io.Copy(dst,src)


}







func main(){
	filename := "MyFile.gz"
	fi,err := os.Open(filename)
	defer fi.Close()
	if err != nil{
		fmt.Println("Bad")
	}
	ft,err := gzip.NewReader(fi)
	if err != nil{
		fmt.Println("BAD")
	}
	r := bufio.NewReader(ft)
	for {
		line,err := r.ReadString('\n')
		if err != nil{
			fmt.Println("BAD")
			os.Exit(0)
		}
		fmt.Println(line)
	}








}