package dbops

import (
	"testing"
)

func clearTables(){
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}


func TestMain(m *testing.M){
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T){
	t.Run("Add",testAddUser)
	t.Run("Get",testGetUser)
	t.Run("Del",testDeleteUser)

}

func testAddUser(t *testing.T){
	err := AddUserCredential("weicheng","123")
	if err != nil{
		t.Errorf("bad add")
	}

}

func testGetUser(t *testing.T){
	_,err := GetUserCredential("weicheng")
	if err != nil{
		t.Errorf("bad Get")
	}
}

func testDeleteUser(t *testing.T){
	err := DeleteUser("weicheng","DDD")
	if err != nil{
		t.Errorf("bad delete")
	}
}

