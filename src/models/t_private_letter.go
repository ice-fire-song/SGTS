package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

// 获取某用户所有的私信
func GetAllPrivateLetter(uid int) (plList *[]PrivateLetter, err error){
	if uid < 0 {
		err = fmt.Errorf("illegal uid")
		logs.Error(err)
		return
	}
	var pl PrivateLetter
	query := "select(plid,user_id,friend_id,sender_id,receiver_id,massage_type,massage_content,send_time,status)" +
		"from t_private_letter where user_id=$1 order by send_time asc"
	rows, err := db.Query(query, uid)
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
func SendPrivateLetter(senderId int, receiverId int, massageContent string) (bool, error){
	if senderId < 1 || receiverId < 1 || len(massageContent) == 0 {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	stmt, err := db.Exec("insert into t_private_letter(user_id,friend_id,sender_id,receiver_id,massage_type,massage_content,status) VALUES ($1,$2,$3,$4,$5,$6,$7)",
		senderId,receiverId,senderId,receiverId,0,massageContent,0)
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
	return true, nil

}
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