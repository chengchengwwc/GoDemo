package main

import (
	"io"
	"time"
	"os"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"net/http"
)


func streamHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid
	video,err := os.Open(vl)
	if err != nil{
		sendErrorResponse(w,http.StatusInternalServerError,"bad")
		return
	}
	w.Header().Set("Content-type","video/mp4")
	http.ServeContent(w,r,"",time.Now(),video)

	defer video.Close()
}

func uploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	r.Body = http.MaxBytesReader(w,r.Body,MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil{
		sendErrorResponse(w,http.StatusBadRequest,"bad")
		return
	}
	file,_,err := r.FormFile("file")
	if err != nil{
		sendErrorResponse(w,http.StatusInternalServerError,"bad")
		return
	}
	
	data,err := ioutil.ReadAll(file)
	if err != nil{
		sendErrorResponse(w,http.StatusInternalServerError,"bad")
		return
	}
	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR + fn,data,0666)
	if err != nil{
		sendErrorResponse(w,http.StatusInternalServerError,"bad")
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w,"Upload Success")

}




