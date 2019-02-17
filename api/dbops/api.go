package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"github.com/humorliang/gflive/api/defs"
	"github.com/humorliang/gflive/api/utils"
	"time"
)

//创建用户
func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbCon.Prepare("INSERT INTO users(login_name,pwd) VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

//获取用户
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbCon.Prepare("SELECT pwd FROM users WHERE login_name=?")
	if err != nil {
		log.Printf("get user error:%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

//删除用户
func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbCon.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("delete user error:%s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

//添加视频
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbCon.Prepare(`INSERT INTO video_info
				(id,author_id,name,display_ctime) VALUES(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmtIns.Close()
	return res, nil
}
