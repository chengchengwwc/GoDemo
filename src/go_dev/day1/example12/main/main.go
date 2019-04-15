package main

import (
	"time"
	"fmt"
)


type student struct{
	name string
}

func writechan(ch chan int){
	for i:=0;i<100;i++{
		ch <- i
	}
	return 
}

func readchan(ch chan int){
	for {
		var b int
		b = <- ch
		fmt.Println(b)
	}
}



func chandemo(){
	inchan := make(chan int,10)
	go writechan(inchan)
	go readchan(inchan)
}

func calc(ch chan int,result chan int,exitchan chan bool){
	
	for v := range ch{
		flag := true
		for i:=2;i<v;i++{
			if (v%i == 0){
				flag = false
				break
			}
		}
		if flag{
			result <- v
		}
	}
	exitchan <- true
}


func goroutedemo(){
	exitchan := make(chan bool,8)
	inchan := make(chan int,1000)
	resultchan := make(chan int,1000)
	go func(){
		for i:=0;i<10000;i++{
			inchan <- i
		}
		close(inchan)
	}()
	
	for i:=0;i<8;i++{
		go calc(inchan,resultchan,exitchan)
	}

	// go func(){
	// 	for v := range resultchan{
	// 		fmt.Println(v)
	// 	}
	// }()


	go func(){
		for i:=0;i<8;i++{
			//启动携程进行检测，不足的时候会阻塞
			<- exitchan
		}
		close(resultchan)
	}()


	for v := range resultchan{
		fmt.Println(v)
	}

	time.Sleep(10*time.Second)

}


func chan_close(){
	//判断chan 关闭
	var ch chan int
	ch = make(chan int,10)
	for i:= 0;i<10;i++{
		ch <- i
	}
	close(ch)
	for{
		var b int
		b,ok := <- ch
		if ok == false{
			break
		}
		fmt.Println(b)
	}
}

func chan_range(){
	var ch chan int
	ch = make(chan int,10)
	for i:=0;i<1000;i++{
		ch <- i
	}
	close(ch)
	for v:= range ch{
		fmt.Println(v)
	}
}



func send(ch chan int,result chan bool){
	for i:=0;i<10;i++{
		ch <- i
	}
	close(ch)
	result <- true
}

func recv(ch chan int,result chan bool){
	for {
		v,ok := <- ch
		if !ok{
			break
		}
		fmt.Println(v)
	}
	result <- true
}



func chan_sync(){
	ch := make(chan int,10)
	exitchan := make(chan bool,2)
	go send(ch,exitchan)
	go recv(ch,exitchan)

	<- exitchan
	<- exitchan


}


func channal_onlyread(){
	 var ch <- chan int
	 fmt.Println(ch)
}

func channal_onlywrite(){
	var ch chan <- int
	fmt.Println(ch)
}


func chan_select(){
	defer func(){
		if err := recover();err!=nil{
			fmt.Println(err)
		}
	}()
	






	var ch chan int
	ch = make(chan int,10)
	go func(){
		for i:=0;i<10;i++{
			ch <- i
		}

	}()
	for {
		//计时器
		select{
		case v := <- ch:
			fmt.Println(v)
		case <- time.After(time.Second):
			fmt.Println("Time out")
		default:
			fmt.Println("getd data timeout")
			time.Sleep(time.Second)
		}
	}
}

func timer(){

}
































func main(){
	var intchan chan int
	var stuchan chan student

	intchan = make(chan int,10)
	stuchan = make(chan student,10)
	stu := student{name:"DDDD"}
	stuchan <- stu
	intchan <- 10

	var stuinterface chan interface{}
	stuinterface = make(chan interface{},10)

	stuinterface <- stu
	var stu01 interface{}
	var stu02 student
	stu01 = <- stuinterface
	stu02,ok := stu01.(student)
	if !ok{
		fmt.Println(stu02)
		return 
	}
	fmt.Println(stu02)











}