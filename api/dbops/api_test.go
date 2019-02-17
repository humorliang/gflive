package dbops

import "testing"

/*
init(db login,truncate tables)->run tests->clear data(truncate tables)
go test

go test -v 详细信息
*/

func clearTables() {
	//truncate 删除表中所有记录
	/*
	区别：truncate和delete的区别
         1、事务：truncate是不可以rollback的，但是delete是可以rollback的；
              原因：truncate删除整表数据(ddl语句,隐式提交)，delete是一行一行的删除，可以rollback
         2、效果：truncate删除后将重新水平线和索引(id从零开始) ,delete不会删除索引
         3、 truncate 不能触发任何Delete触发器。
         4、delete 删除可以返回行数
	*/
	dbCon.Exec("truncate users")
	dbCon.Exec("truncate video_info")
	dbCon.Exec("truncate comments")
	dbCon.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("ReGet", testReGetUser)
}

//创建用户
func testAddUser(t *testing.T) {
	err := AddUserCredential("user", "123")
	if err != nil {
		t.Errorf("error of addUser:%v", err)
	}
}

//获取用户
func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("user")
	if pwd != "123" || err != nil {
		t.Errorf("error of getUser")
	}
}

//删除用户
func testDeleteUser(t *testing.T) {
	err := DeleteUser("user", "123")
	if err != nil {
		t.Errorf("error of delete user:%v", err)
	}

}

//重新查询一次
func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("user")
	if err != nil {
		t.Errorf("error of reGetUser:%v", err)
	}
	if pwd != "" {
		t.Errorf("Delete user test fails")
	}
}
