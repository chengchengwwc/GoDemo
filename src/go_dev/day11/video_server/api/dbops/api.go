package dbops


import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"go_dev/day11/video_server/api/defs"
	"go_dev/day11/video_server/api/util"
	"time"
)

// func openConn() *sql.DB{
// 	dbConn,err := sql.Open("mysql","root:root@tcp(localhost:3306)/video_server?charset=utf-8")
// 	if err != nil{
// 		panic(err.Error())
// 	}
// 	return dbConn

// }


func AddUserCredential(loginName string,pwd string) error{
	stmtIns,err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)")
	if err != nil{
		return err
	}
	_,err = stmtIns.Exec(loginName,pwd)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil

}

func GetUserCredential(loginName string)(string,error){
	stmtOut,err := dbConn.Prepare("SELECT pwd from users WHERE login_name=?")
	if err != nil{
		log.Printf("%s",err)
		return "",err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil{
		return "",err
	}

	defer stmtOut.Close()
	return pwd,nil
}

func DeleteUser(loginName string,pwd string) error{
	stmtDel,err := dbConn.Prepare("DElETE FROM users WHERE login_name=? and password=?")
	if err != nil{
		log.Print("DeleteUsers error: %s",err)
		return err
	}
	_,err = stmtDel.Exec(loginName,pwd)
	if err != nil{
		return err
	}

	defer stmtDel.Close()
	return nil

}

func AddNewVideo(aid int,name string)(*defs.VideoInfo,error){
	// create id
	vid,err := util.NewUUID()
	if err != nil{
		return nil,err
	}
	t := time.Now()
	//固定格式
	ctime := t.Format("Jan 02 2006, 15:04:05")

	stmIns,err := dbConn.Prepare(`INSERT INTO video_info (id,author_id,name,display_ctime) VALUES(?,?,?,?)`)
	if err != nil{
		return nil,err
	}
	_,err = stmIns.Exec(vid,aid,name,ctime)
	if err != nil{
		return nil,err
	}

	res := &defs.VideoInfo{Id:vid,AuthorId:aid,Name:name,DisplayCtime:ctime}

	defer stmIns.Close()
	return res,nil

}

func GetVideoInfo(vid string)(*defs.VideoInfo,error){
	var aid int 
	var dct string
	var name string
	
	stmOut, err := dbConn.Prepare("SELECT author_id,name,display_ctime from video_info WHERE id =?")
	if err != nil{
		return nil,err
	
	}
	err = stmOut.QueryRow(vid).Scan(&aid,&name,&dct)
	if err == sql.ErrNoRows{
		return nil,nil
	}
	defer stmOut.Close()

	res := &defs.VideoInfo{Id:vid,AuthorId:aid,Name:name,DisplayCtime:dct}
	return res,nil
}

func DeleteVideoInfo(vid string) error{
	stmDel,err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil{
		return err
	}
	_,err = stmDel.Exec(vid)

	defer stmDel.Close()


	if err != nil{
		return err
	}
	return nil
}


func AddNewComments(vid string,aid int,content string) error{
	id,err := util.NewUUID()
	if err != nil{
		return err
	}
	stmIns,err := dbConn.Prepare("INSERT INTO comments (id,video_id,author_id content) VALUES(?,?,?,?)")
	if err != nil{
		return err
	}
	_,err = stmIns.Exec(id,vid,aid,content)
	if err != nil{
		return err
	}
	defer stmIns.Close()
	return nil
}


func ListComments(vid string,from,to int)([]*defs.Comment,error){
	stmOut,err := dbConn.Prepare(`
	   SELECT comments.id,users.LOgin_name,comments.content FROM comments
	   INNER JOIN users ON comments.author_id = users.id WHERE comments_id =? and
	   comments.time > FROM_UNIXTIME(?) AND comments_time <= FROM_UNIXTIME(?)
	`)
	var res []*defs.Comment

	rows,err := stmOut.Query(vid,from,to)
	if err != nil{
		return res,err
	}
	for rows.Next(){
		var id,name,content string
		err := rows.Scan(&id,&name,&content)
		if err != nil{
			return res,err
		}
		c := &defs.Comment{Id:id,VideoId:vid,Author:name,Content:content}
		res = append(res,c)
		
	}
	defer stmOut.Close()
	return res,nil

}



























