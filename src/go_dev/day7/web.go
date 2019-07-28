package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"golang.org/x/tools/go/ssa/interp/testdata/src/strings"
	"io/ioutil"
	"net/http"
	"os"
)


const prefix = "/list/"

type userError string


func (e userError) Message() string{
	return string(e)
}


func (e userError) Error() string{
	return e.Message()
}


func HandleFileList(writer http.ResponseWriter,request *http.Request) error{
	fmt.Println()
	if strings.Index(request.URL.Path,prefix) != 0{
		return userError(fmt.Sprintf("path %s must start with %s",request.URL.Path,prefix))
	}
	path := request.URL.Path[len(prefix):]
	file,err := os.Open(path)
	if err != nil{
		return err
	}
	defer file.Close()
	all,err := ioutil.ReadAll(file)
	if err != nil{
		return err
	}
	writer.Write(all)
	return nil


}


func tryRecover(){
	defer func() {
		r := recover()
		if r == nil{
			fmt.Println("Nothing")
		}
		if err,ok := r.(error); ok{
			fmt.Println(err)
		}else{
			panic(fmt.Sprintf("Bad"))
		}

	}()
}






type appHander func(writer http.ResponseWriter,
	request *http.Request) error


func errWrapper(handler appHander) func(w http.ResponseWriter,r *http.Request){
	return func(writer http.ResponseWriter,request *http.Request){

		err := handler(writer,request)
		if err != nil{
			fmt.Println(err)
			if userErr,ok := err.(userError);ok{
				http.Error(writer,
							userErr.Message(),
							http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch{
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)

		}
	}
}



func StartHttp(){
	http.HandleFunc("/",errWrapper(HandleFileList))

	err := http.ListenAndServe(";8888",nil)
	if err != nil{
		panic(err)
	}

}






