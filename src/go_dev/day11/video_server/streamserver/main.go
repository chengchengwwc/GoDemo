package main


import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)


type middleWareHandler struct{
	r *httprouter.Router
	l *ConnLimiter
}

//中间件开发
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	if !m.l.GetConn(){
		sendErrorResponse(w,http.StatusTooManyRequests,"To many reqeust")
		return
	}
	m.r.ServeHTTP(w,r)
	defer m.l.ReleaseConn()
}




func NewMiddleWareHandler(r *httprouter.Router,cc int) http.Handler{
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m

}



func RegisterHandler() *httprouter.Router{
	router := httprouter.New()
	router.GET("/videos/:vid-id",streamHandler)
	router.POST("/upload/:vid-id",uploadHandler)
	return router

}

func main(){
	r := RegisterHandler()
	mh := NewMiddleWareHandler(r,10)
	http.ListenAndServe(":9000",mh)

}




