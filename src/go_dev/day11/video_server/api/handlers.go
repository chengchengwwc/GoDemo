package main

import (
	"io"
	"encoding/json"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"go_dev/day11/video_server/api/defs"
	"go_dev/day11/video_server/api/dbops"
	"go_dev/day11/video_server/api/session"
)


func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	//io.WriteString(w,"Create user Handler")
	res,_ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err :=  json.Unmarshal(res,ubody); err != nil{
		sendErorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	if err := dbops.AddUserCredential(ubody.Username,ubody.Pwd); err != nil{
		sendErorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success:true,SessionId:id}
	if resp, err := json.Marshal(su); err != nil{
		sendErorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	} else{
		sendNormalResponse(w,string(resp),201)
	}

}


func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	uname := p.ByName("user_name")
	io.WriteString(w,uname)

}


