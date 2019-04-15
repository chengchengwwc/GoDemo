package main


import "fmt"
import "github.com/hpcloud/tail"
import "time"

func main(){
	filename := "./my.log"
	tails,err := tail.TailFile(filename,tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset:0,Whence:2},
		MustExist:false,
		Poll: true,
	})
	if err != nil{
		fmt.Println("tail file err:",err)
		return
	}
	var msg *tail.Line
	var ok bool
	for true{
		msg,ok = <- tails.Lines
		if !ok {
			fmt.Println("tail file close reopen filename %s\n",tails.filename)
			time.Sleep(100*time.Millisecond)
			continue
		}
		fmt.Println(msg)
	}
}