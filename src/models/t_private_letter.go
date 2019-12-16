package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

type Sender struct { // 私信用户列表相关信息
	UserId   int64  `json:"user_id"`

	FriendId  int64  `json:"friend_id"`
	Username              string  `json:"username"`
	HeadSculpturePath   string  `json:"head_sculpture_path"`
	NoReadCount int64    `json:"no_read_count"`
}
// 获取与某个用户私信的所有用户
func GetSenderList(uid int)(SenderList *[]Sender, err error)  {
	if uid < 1 {
		err = fmt.Errorf("illegal gid")
		logs.Error(err)
		return
	}
	// 获取用户列表
	queryStr1 := "select t_user.uid, username, head_sculpture_path  from t_user where t_user.uid in  (select friend_id from t_private_letter,t_user where user_id=$1"+
	"and t_user.uid=t_private_letter.friend_id and t_private_letter.status=$2 group by friend_id);"
	// 1, 未读 2,已读 3 删除
	// queryStr2 := "select friend_id, username,head_sculpture_path,count(plid) as count from t_private_letter,t_user where user_id=$1 " +
	// 	"t_user.uid=t_private_letter.friend_id and status=$2 group by friend_id"
	rows, err := db.Query(queryStr1, uid, 1)
	if err != nil {
		logs.Error(err)
		return
	}
	defer rows.Close()
	SenderList = new([]Sender)
	var sender Sender
	var user_id   int64 = int64(uid)

	var friend_id          int64
	var username          string
	var head_sculpture_path   string
	var no_read_count int64
	for rows.Next() {

		err = rows.Scan(&friend_id, &username, &head_sculpture_path, &no_read_count)
		if err != nil {
			logs.Error(err)
			return
		}
		sender.UserId = user_id
		sender.FriendId = friend_id
		sender.Username = username
		sender.HeadSculpturePath = head_sculpture_path
		sender.NoReadCount = no_read_count
		*SenderList = append(*SenderList, sender)
	}
	return
}
// 获取某用户所有的私信
func GetRecords(user_id int64, friend_id int64) (plList *[]PrivateLetter, err error){
	if user_id < 0 || friend_id < 0 {
		err = fmt.Errorf("illegal param")
		logs.Error(err)
		return
	}
	var pl PrivateLetter
	query := "select(plid,user_id,friend_id,sender_id,receiver_id,massage_type,message_content,send_time,status)" +
		"from t_private_letter where user_id=$1 and friend_id=$2 order by send_time asc"
	rows, err := db.Query(query, user_id, friend_id)
	defer rows.Close()
	plList = new([]PrivateLetter)
	for rows.Next() {
		var plid int64
		var user_id int64
		var friend_id int64
		var sender_id int64
		var receiver_id int64
		var massage_type int64
		var massage_content sql.NullString
		var send_time time.Time
		var status int64
		err = rows.Scan(&plid,&user_id,&friend_id,&sender_id,&receiver_id,&massage_type,&massage_content,&send_time,&status)
		if err != nil {
			logs.Error(err)
			return
		}
		pl.PLid = plid
		pl.UserId = user_id
		pl.FriendId = friend_id
		pl.SenderId = sender_id
		pl.ReceiverId = receiver_id
		pl.MassageType = massage_type
		pl.MassageContent = massage_content.String
		pl.SendTime = send_time
		pl.Status = status
		*plList = append(*plList, pl)
	}
	return
}


func SendPrivateLetter(senderId int, receiverId int, messageContent string) (bool, error){
	if senderId < 1 || receiverId < 1 || len(messageContent) == 0 {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	// message_type 1 普通消息 2 系统消息
	qureyStr1 := "insert into t_private_letter(user_id,friend_id,sender_id,receiver_id,massage_type,message_content)  VALUES ($1,$2,$3,$4,$5,$6);"
	qureyStr2 := "insert into t_private_letter(user_id,friend_id,sender_id,receiver_id,massage_type,message_content)  VALUES ($1,$2,$3,$4,$5,$6);"
	stmt, err := db.Exec(qureyStr1, senderId,receiverId,senderId,receiverId,1,messageContent)
	if err != nil {
		logs.Error(err)
		return false, err
	}
	affect, err := stmt.RowsAffected()
	if err != nil {
		logs.Error(err)
		return false, err
	}
	if affect == 0 {
		return false,errors.New("Affected rows is 0 ")
	}

	stmt, err = db.Exec(qureyStr2,receiverId,senderId,senderId,receiverId,1,messageContent)
	if err != nil {
		logs.Error(err)
		return false, err
	}
	affect, err = stmt.RowsAffected()
	if err != nil {
		logs.Error(err)
		return false, err
	}
	if affect == 0 {
		return false,errors.New("Affected rows is 0 ")
	}
	return true, nil

}
// 删除指定私信
func DeletePrivateLetter(plid int)(bool, error) {
	if plid < 1 {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	//删除数据
	stmt, err := db.Prepare("delete from t_private_letter where plid=$1")
    if err != nil {
    	logs.Error(err)
    	return false, err
	}

	res, err := stmt.Exec(plid)
	if err != nil {
		logs.Error(err)
		return false, err
	}


	affect, err := res.RowsAffected()
	if affect == 0 {
		return false,errors.New("Affected rows is 0 ")
	}
	return true, nil
}

// 将两个私信用户的所有私信记录全转为已读
func ToReaded(user_id, friend_id int) (bool, error) {
	if user_id < 0 || friend_id < 0 {
		err := fmt.Errorf("illegal params")
		logs.Error(err)
		return false, err
	}
	stmt, err := db.Prepare("update t_private_letter set status=$1 where user_id=$2 and friend_id=$3")
	if err != nil {
		logs.Error(err)
		return false, err
	}
	res, err := stmt.Exec(2, user_id, friend_id)
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