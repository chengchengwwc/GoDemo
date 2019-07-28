package session

import (
	"time"
	"sync"
	"go_dev/day11/video_server/api/defs"
	"go_dev/day11/video_server/api/dbops"
	"go_dev/day11/video_server/api/util"
)

type SimepleSession struct{
	Username string
	TTL      int64
}

//线程安全
var sessionMap *sync.Map


func init(){
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano()/1000000
}


func deleteexpiredSession(sid string){
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}


func LoadSessionsFromDB(){
	//读入全局
	r,err := dbops.RetrieveAllSeeions()
	if err != nil{
		return
	}
	r.Range(func(k,v interface{}) bool{
		ss := v.(*defs.SimepleSession)
		sessionMap.Store(k,ss)
		return true
	})


}

func GenerateNewSessionId(un string) string{
	id,_ := util.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30 * 60 * 1000
	ss := &defs.SimepleSession{Username:un,TTL:ttl}
	sessionMap.Store(id,ss)
	dbops.InsertSession(id,ttl,un)
	return id
}

func IsSessionExpired(sid string)(string,bool){
	ss,ok := sessionMap.Load(sid)
	if ok{
		ct := nowInMilli()
		if ss.(*defs.SimepleSession).TTL < ct{
			return "",true
		}
		return ss.(*defs.SimepleSession).Username,false

	}
	return "",false

}




