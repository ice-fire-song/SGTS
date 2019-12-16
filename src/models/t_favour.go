package models

import (
	//"database/sql"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	//"time"
)

func AddFavour(gid, uid int) (bool, error){
	if gid < 1 || uid < 1{
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	stmt, err := db.Exec("insert into t_favour(gid, uid,fd_id) VALUES ($1,$2,$3)",
		gid,uid,3)
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
		return false, errors.New("Affected rows is 0 ")
	}
	return  true, nil
}
// 在删除收藏夹的同时要删除其下所有收藏
func RemoveFavour(fd_id int) error{
	if fd_id < 1 {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return err
	}
	//删除数据
	stmt, err := db.Prepare("delete from t_favour where fd_id=$1 ")
    if err != nil {
    	logs.Error(err)
    	return  err
	}

	res, err := stmt.Exec(fd_id)
	if err != nil {
		logs.Error(err)
		return err
	}

	affect, err := res.RowsAffected()
	if affect == 0 {
		return errors.New("Affected rows is 0 ")
	}
	return nil
}
func SeeFavourStatus(gid int64, uid int64)(bool, error) {
	if gid < 0 || uid < 0 {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	var fid int
	row := db.QueryRow("select fid from t_favour where uid=$1 and gid=$2", uid,gid)
    err := row.Scan(&fid)
    if err != nil {
		return false, nil
	}
	if fid > 1 {
		return true, nil
	}
	return false, nil
}


func GetFavourOfGood(uid int)(glList *[]Favour, err error) {
	if uid < 1 {
		err = fmt.Errorf("illegal gid")
		logs.Error(err)
		return
	}
	rows, err := db.Query("select fid, gid from t_favour where uid=$1", uid)
	if err != nil {
		logs.Error(err)
		return
	}
	defer rows.Close()
	glList = new([]Favour)
	var favour Favour
	for rows.Next() {
		var fid int64
		var gid int64
		err = rows.Scan(fid, gid)
		if err != nil {
			logs.Error(err)
			return
		}
		favour.Fid = fid
		favour.Gid = gid
		*glList = append(*glList, favour)
	}
	return
}
func CancelFavour(fid int)(bool, error) {
	if fid < 1 {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	//删除数据
	stmt, err := db.Prepare("delete from t_favour where fid=$1")
	if err != nil {
		logs.Error(err)
		return false, err
	}

	res, err := stmt.Exec(fid)
	if err != nil {
		logs.Error(err)
		return false, err
	}


	affect, err := res.RowsAffected()
	if affect == 0 {
		return false, errors.New("Affected rows is 0 ")
	}
	return true, nil
}

