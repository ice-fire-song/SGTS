package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

type Good struct {
	Gid   int `json:"gid"`
	Uid   int `json:"uid"`
	GType int `json:"gtype"`
	GName string `json:"gname"`
	GPrice  float64 `json:"gprice"`
	GDetail string `json:"gdetail"`
	CategoryId  int `json:"category_id"`
	ClickNumber int `json:"click_number"`
	Status  int  `json:"status"`
	//MobilePhoneNumber string `json:"mobilephone_number"`
	//GLiaison  string `json:"gliaison"`
	//Openid  string `json:"openid"`
	//Qq       string `json:"qq"`
	ReleaseTime time.Time `json:"release_time"`
}

func AddGood(goodInfo *Good) (bool, error) {
	if goodInfo == nil {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	stmt, err := db.Exec("insert into t_goods(uid,gtype,gname,gprice,gdetail,category_id,status) VALUES ($1,$2,$3,$4,$5,$6,$7)",
		goodInfo.Uid,goodInfo.GType,goodInfo.GName,goodInfo.GPrice,goodInfo.GDetail,goodInfo.CategoryId,goodInfo.Status)
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

func ModifyGood(goodInfo *Good) (bool, error) {
	if goodInfo == nil {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	stmt, err := db.Prepare("update t_goods set gtype=$1,gname=$2,gprice=$3,gdetail=$4,category_id=$5,status=$6")
	if err != nil {
		logs.Error(err)
		return false, err
	}
	res, err := stmt.Exec(goodInfo.GType,goodInfo.GName,goodInfo.GPrice,goodInfo.GDetail,goodInfo.CategoryId,goodInfo.Status)
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

// 根据货品类型获取
// 分为商品、需求、免费货品
func GetAllGoods(goodType int, label int, keyWord string)(goodsList *[]Good, err error) {
	var query string
	var rows *sql.Rows
	if goodType == 0 {
		query = "select(gid,uid,gtype,gname,gprice,gdetail,category_id,status,release_time)" +
			"from t_goods where label=$1 and gname like $2 order by send_time asc"
		rows, err = db.Query(query, label, keyWord)
	}else if label == 0 {
		query = "select(gid,uid,gtype,gname,gprice,gdetail,category_id,status,release_time)" +
			"from t_goods where gtype=$1 and gname like $2 order by send_time asc"
		rows, err = db.Query(query, goodType, keyWord)
	}else {
		query = "select(gid,uid,gtype,gname,gprice,gdetail,category_id,status,release_time)" +
			"from t_goods where gtype=$1 and gname like $2 gid = (select gid from t_goods_label where label=$3)order by send_time asc"
		rows, err = db.Query(query, goodType, keyWord,label)
	}
	if err != nil {
		logs.Error(err)
		return
	}

	defer rows.Close()
	goodsList = new([]Good)
	var good Good
	for rows.Next() {
		var gid int
		var uid int
		var gtype int
		var gname string
		var gprice float64
		var gdetail string
		var category_id int
		var status int
		var release_time time.Time
		err = rows.Scan(&gid,&uid,&gtype,&gname,&gprice,&gdetail,&category_id,&status,&release_time)
		if err != nil {
			logs.Error(err)
			return
		}
	    good.Gid = gid
	    good.Uid = uid
	    good.GType = gtype
	    good.GName = gname
	    good.GPrice = gprice
	    good.GDetail = gdetail
	    good.CategoryId = category_id
	    good.Status = status
	    good.ReleaseTime = release_time
		*goodsList = append(*goodsList, good)
	}
	return
}

func DeleteGood(gid int)(bool, error) {
	if gid < 1 {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	//删除数据
	stmt, err := db.Prepare("delete from t_goods where gid=$1")
	if err != nil {
		logs.Error(err)
		return false, err
	}

	res, err := stmt.Exec(gid)
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