package main

import "context"
import "fmt"
import "io/ioutil"
import "net/http"
import "time"


type Result struct{
	r *http.Response
	err error
}

func process(){
	ctx,cancel := context.WithTimeout(context.Background(),2*time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport:tr}
	c := make(chan Result,1)
	req,err := http.NewRequest("GET","http://www.baidu.com",nil)
	if err != nil{
		fmt.Println(err)
		return
	}
	go func(){
		resp,err := client.Do(req)
		pack := Result{r:resp,err:err}
		c <- pack
	}()
	select{
	case <- ctx.Done():
		tr.CancelRequest(req)
		res := <- c
		fmt.Println(res.err)
	case res:= <- c:
		defer res.r.Body.Close()
		out,_ := ioutil.ReadAll(res.r.Body)
		fmt.Println(out)
	}
	return

}

func main(){
	process()
}
