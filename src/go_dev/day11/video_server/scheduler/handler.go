package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"go_dev/day11/video_server/api/dbops"
)

func vidDelRecHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	vid := p.ByName("vid-id")
	if len(vid) == 0{
		sendResponse(w,400,"video is should not be empty")
		return
	}
	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil{
		sendResponse(w,500,"Internal server error")
		return
	}
	sendResponse(w,200,"")
}


