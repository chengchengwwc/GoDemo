package dbops


import (
	"strconv"
	"go_dev/day11/video_server/api/defs"
	"sync"
	"log"
)

func InsertSession(sid string,ttl int64,uname string) error{
	ttlstr := strconv.FormatInt(ttl,10)
	stmtIns,err := dbConn.Prepare("INSERT INTO session (session_id,TTL,login_name) VALUES(?,?,?)")
	if err != nil{
		return err
	}
	_,err = stmtIns.Exec(sid,ttlstr,uname)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil

}


func RetrieveSession(sid string) (*defs.SimepleSession,error){
	ss := &defs.SimepleSession{}
	var ttl string
	var uname string
	stmtOut,err := dbConn.Prepare("SELECT TTL,login_name FROM session WHERE session_id=?")
	if err != nil{
		return nil,err
	}

	err = stmtOut.QueryRow(sid).Scan(&ttl,&uname)
	if err != nil{
		return nil,err
	}
	defer stmtOut.Close()
	
	if res,err := strconv.ParseInt(ttl,10,64);err == nil{
		ss = &defs.SimepleSession{TTL:res,Username:uname}
	}else{
		return nil,err
	}
	return ss,nil

}


func RetrieveAllSeeions()(*sync.Map,error){
	m := &sync.Map{}
	stmtOut,err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil{
		log.Printf("%s",err)
		return nil,err
	}
	rows,err := stmtOut.Query()
	if err != nil{
		log.Printf("%s",err)
		return nil,err
	}
	for rows.Next(){
		var id string
		var ttlstr string
		var login_name string
		err = rows.Scan(&id,&ttlstr,&login_name)
		if err != nil{
			log.Printf("retrive sessions error: %s",err)
			break
		}
		ttl,err1 := strconv.ParseInt(ttlstr,10,64)
		if err1 == nil{
			ss := &defs.SimepleSession{Username:login_name,TTL:ttl}
			m.Store(id,ss)
			log.Printf("Session id:%s,ttl:%d",id,ss.TTL)
		}
	}
	defer stmtOut.Close()
	return m,nil

}

func DeleteSession(sid string) error{
	stmtOut,err := dbConn.Prepare("DELETE FROM sessions WHERE session_id=?")
	if err != nil{
		log.Printf("%s",err)
		return err
	}
	_,err = stmtOut.Query(sid)
	if err != nil{
		return err
	}
	return nil
}
