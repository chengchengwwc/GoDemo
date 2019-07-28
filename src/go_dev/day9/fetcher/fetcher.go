package fetcher

import (
	"bufio"
	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"net/http"
	_ "fmt"
	"time"
	"unicode"
)



func determineEncoding(r *bufio.Reader) encoding.Encoding{
	byte,err := r.Peek(1024)
	if err != nil{
		return unicode.UTF8
	}
	e,_,_ := charset.DetermineEncoding(byte,"")
	return e


}


var rateLimite = time.Tick(100*time.Millisecond)


func Fetch(url string)([]byte,error){
	//定时器
	<- rateLimite
	resp,err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil,errors.New("Bad")
	}

	// 类型转换
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader,e.NewDecoder())

	all,err := ioutil.ReadAll(utf8Reader)
	if err != nil{
		panic(err)
	}
	return all,err

}