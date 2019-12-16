package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)


// 判断用户的身份
func IsAdmin(username string) (res bool, err error) {
	if len(username) == 0 {
		err = fmt.Errorf("Call IsAdmin with a empty username")
		logs.Error(err)
		return
	}
	var user_role int
	err = db.QueryRow("select user_role from t_user where username=$1", username).Scan(&user_role)
	if err != nil {
		logs.Error(err)
		return
	}
	if user_role == 0 {
		return false, nil
	} else if user_role == 1 {
		return true, nil
	}
	err = fmt.Errorf("出现未设定的角色")
	logs.Error(err)
	return
}

// 登录使用，判断用户/密码是否正确
func LoginVerification(username, password string) (isUserExist bool, err error) {
	if len(username) == 0 {
		err := fmt.Errorf("Call IsUserExist with a empty username")
		logs.Error(err)
		return false, err
	}
	isUserExist, err = IsUserExist(username)
	if err != nil {
		logs.Error(err)
		return
	}
	if !isUserExist {
		return false, nil
	}

	var uname string
	err = db.QueryRow("select username from t_user where username=$1 and password=$2", username, password).Scan(&uname)
	if err != nil {
		logs.Error(err)
		return
	}
    if username == uname {
    	return true, nil
	}else {
		return false, nil
	}
}
// 判断该用户名是否已存在
func IsUserExist(username string)  (isUserExist bool, err error) {
	if len(username) == 0 {
		err := fmt.Errorf("Call IsUserExist with a empty username")
		logs.Error(err)
		return false, err
	}
	row, err := db.Query("select password from t_user where username=$1", username)
	if err != nil {
		logs.Error(err)
		return
	}
	if row == nil {
		err = errors.New("row is nil")
		logs.Error(err)
		return
	}
	var pwd sql.NullString
	for row.Next() {
		err = row.Scan(&pwd)
		if err != nil {
			logs.Error(err)
			return
		}
	}
	if pwd.Valid {
		isUserExist = true
		return
	} else {
		isUserExist = false
		return
	}
	return
}
//// 登录验证
//// -1 用户不存在，0 用户是普通会员，1 用户是管理员
//func LoginVerification(username, password string) (int, error) {
//	logs.Info("ISSSSSSSSSSSSSSSSSSS")
//	isUserExist, err := IsUserExist(username)
//	logs.Info("ISSSSSSSSSSSSSSSSSSS")
//	if err != nil {
//		logs.Error(err)
//		return 0, err
//	}
//	if !isUserExist {
//		return -1, nil
//	}
//	isAdmin, err := IsAdmin(username)
//	logs.Info("ISSSSSSSSSSSSSSSSSSS")
//	if err != nil {
//		logs.Error(err)
//		return 0, err
//	}
//	if isAdmin {
//		return 1, nil
//	} else {
//		return 0, nil
//	}
//}

// 注册
func RegisterUser(username, password ,mailbox string) error {
	if len(username) == 0 || len(password) == 0 || len(mailbox) == 0{
		err := fmt.Errorf("username or password is null")
		logs.Error(err)
		return err
	}
	// user_role表示普通用户，status为1表示账号正常状态
	stmt, err := db.Exec("insert into t_user(username, password,user_role,mailbox,status) values($1,$2,$3,$4,$5)",
		username, password, 0, mailbox, 1)
	if err != nil {
		logs.Error(err)
		return err
	}
	affect, err := stmt.RowsAffected()
	if err != nil {
		logs.Error(err)
		return err
	}
	if affect == 0 {
		return errors.New("Affected rows is 0 ")
	}
	return nil
}

func GetUserInfo(username string) (user User, err error) {
	if len(username) == 0 {
		err = fmt.Errorf("username can not is null")
		logs.Error(err)
		return
	}
	var uid  int64
	var user_role int64
	var head_sculpture_path string
	var label sql.NullString
	var mailbox string
	var create_time time.Time
	var status int64
	queryStr := "select uid,user_role,head_sculpture_path, label, mailbox, create_time,status from t_user where username=$1 "
	err = db.QueryRow(queryStr, username).Scan(&uid,&user_role,&head_sculpture_path, &label, &mailbox, &create_time,&status)
	if err != nil {
		logs.Error(err)
		return
	}
	user.Uid = uid
	user.Username = username
	user.UserRole = user_role
	user.HeadSculpturePath = head_sculpture_path
	user.Label = label.String
	user.Mailbox = mailbox
	user.CreateTime = create_time
	user.Status = status
	return
}
func ModifyUserInfo(uid int, head_sculpture_path, label, mailbox string) (bool, error) {
	if uid < 0 || len(head_sculpture_path) == 0 || len(label) == 0 {
		err := fmt.Errorf("illegal params")
		logs.Error(err)
		return false, err
	}
	stmt, err := db.Prepare("update t_user set head_sculpture_path=$1, label=$2, mailbox=$3 where uid=$4")
	if err != nil {
		logs.Error(err)
		return false, err
	}
	res, err := stmt.Exec(head_sculpture_path, label, mailbox, uid)
	if err != nil {
		logs.Error(err)
		return false, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		logs.Error(err)
		return false, err
	}
	if affect == 0 {
		return false,errors.New("Affected rows is 0 ")
	}
	return true, nil
}
func ModifyPWD(uid int, oldPassword, newPassword string)(bool, error) {
	if uid < 0 {
		err := fmt.Errorf("illegal params")
		logs.Error(err)
		return false, err
	}
	if len(oldPassword) == 0 || len(newPassword) == 0 {
		logs.Info("newPassword/oldPassword is null")
		return false, nil
	}
	stmt, err := db.Prepare("update t_user set password=$1 where uid=$2")
	if err != nil {
		logs.Error(err)
		return false, err
	}
	res, err := stmt.Exec(newPassword, uid)
	if err != nil {
		logs.Error(err)
		return false, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		logs.Error(err)
		return false, err
	}
	if affect == 0 {
		return false,errors.New("Affected rows is 0 ")
	}
	return true, nil
}