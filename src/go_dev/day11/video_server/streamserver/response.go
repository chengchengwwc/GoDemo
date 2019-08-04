package main

import (
	"io"
	"net/http"
	
)


func sendErrorResponse(w http.ResponseWriter,src int,errMsg string){
	w.WriteHeader(src)
	io.WriteString(w,errMsg)
}