package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type middlwWareHandler struct{
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler{
	//构建中间件
	m := middlwWareHandler{}
	m.r = r
	return m
}

func (m middlwWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	// 中间件写到这里来
	// check session
	validateUserSession(r)
	m.r.ServeHTTP(w,r)
}



func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()

	router.POST("/user",CreateUser)

	router.POST("/user/:user_name",Login)

	return router
}


func main()  {
	//启动服务
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000",mh)

}








