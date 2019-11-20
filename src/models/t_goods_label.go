package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func AddLabel(gid int, label string) (bool, error){
	if gid < 1 || len(label) == 0 {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	stmt, err := db.Exec("insert into t_goods_label(gid,label_name) VALUES ($1,$2)",
		gid,label)
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
func GetLabelOfGood(gid int)(glList *[]string, err error) {
	if gid < 1 {
		err = fmt.Errorf("illegal gid")
		logs.Error(err)
		return
	}
	rows, err := db.Query("select label_name from t_goods_label where gid=$1", gid)
	if err != nil {
		logs.Error(err)
		return
	}
	defer rows.Close()
	glList = new([]string)
	for rows.Next() {
		var label string
		err = rows.Scan(label)
		if err != nil {
			logs.Error(err)
			return
		}
		*glList = append(*glList, label)
	}
	return
}
func DeleteLabel(gid int, label string)(bool, error) {
	if gid < 1 || len(label) == 0 {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	//删除数据
	stmt, err := db.Prepare("delete from t_goods_label where gid=$1 and labal_name=$2")
	if err != nil {
		logs.Error(err)
		return false, err
	}

	res, err := stmt.Exec(gid, label)
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
func ModifyLabel(gid int, oldLabel string, newLabel string) (bool, error) {
	//更新数据
	stmt, err := db.Prepare("update t_goods_label set label_name=$1 where gid=$2 and label_name=$3")
	if err != nil {
		logs.Error(err)
		return false, err
	}

	res, err := stmt.Exec(newLabel, gid, oldLabel)
	if err != nil {
		logs.Error(err)
		return false, err
	}

	affect, err := res.RowsAffected()
	if affect != 0 {
		return false, errors.New("Affected rows is 0 ")
	}
	return true, nil
}