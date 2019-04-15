package main

import(
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct{
	UserId   int `db:"user_id"`
	UserName string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"` 
}

type Place struct{
	Country string `db:"country"`
	City    string `db:city`
	TelCode string `db:"email"`
}

var Db *sqlx.DB

func init(){
	//初始化
	database,err := sqlx.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err)
		return
	}
	Db = database
}

func main(){
	//插入操作
	r,err := Db.Exec("insert into person(username,sex,email)values(?,?,?)","dd","dd","ss")
	if err != nil{
		fmt.Println(err)
		return 
	}
	id,err := r.LastInsertId()
	if err != nil{
		fmt.Println("exec failed")
	}
	fmt.Println("insert succ:",id)

	//查询操作
	var person Person
	errm := Db.Select(&person,"select user_id,username,sex,email from person where user_id=?",1)
	if errm != nil{
		fmt.Println(errm)
		return
	}
	//删除操作
	_,errd := Db.Exec("delete from person where user_id=?",1)
	if errd != nil{
		fmt.Println(errd)
	}
	//更新操作
	_,erru := Db.Exec("update person set username=? where user_id=?","DDD",1)
	if erru != nil{
		fmt.Println(erru)
	}

	












}




