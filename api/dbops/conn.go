package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbCon *sql.DB
	err   error
)

//初始化数据库连接
func init() {
	dbCon, err = sql.Open("mysql", "root:123456@tcp(192.168.1.100:6666)/video_cms?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
