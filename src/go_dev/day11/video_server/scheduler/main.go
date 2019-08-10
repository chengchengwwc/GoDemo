package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"go_dev/day11/video_server/scheduler/taskrunner"
)


func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	router.get("/video-delete-record/:vid-id",vidDelRecHandler)
	return router

}





func main(){
	go taskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001",r)


}


