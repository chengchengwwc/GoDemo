package dbops

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func ReadVideoDelteionRecord(count int) ([]string,error){
	stmOut,err := dbConn.Prepare("SELECT video_id from video_del LIMIT ?")
	var ids []string
	if err != nil{
		return ids,err
	}
	rows,err := stmOut.Query(count)
	if err != nil{
		log.Printf("bad")
		return ids, err
	}
	for rows.Next(){
		var id string
		if err := rows.Scan(&id);err != nil{
			return ids,err
		}
		ids = append(ids,id)
	}
	defer stmOut.Close()
	return ids,nil
}

func DelVideoDeltetionRecord(vid string) error{
	stmtDel ,err := dbConn.Prepare("DELETE FROM video_del_recode where ?")
	if err != nil {
		return err
	}
	_,err = stmtDel.Exec(vid)
	if err != nil{
		return err
	}

	defer stmtDel.Close()
	return nil

}











