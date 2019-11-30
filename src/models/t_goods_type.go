package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

func GetGoodsType()(gtList *[]GoodsType, err error){
	rows, err := db.Query("select gt_id, type_name, create_time from t_goods_type")
	if err != nil {
		logs.Error(err)
		return
	}
	defer rows.Close()
	gtList = new([]GoodsType)
	var goodType GoodsType
	var createTime  time.Time
	var type_name  string
	var gt_id int64

	for rows.Next() {
		err = rows.Scan(&gt_id, &type_name, &createTime)
		if err != nil {
			logs.Error(err)
			return
		}
		goodType.Gtid = gt_id
		goodType.TypeName = type_name
		goodType.CreateTime = createTime
		*gtList = append(*gtList, goodType)
	}
	return
}


func GetGoodsByType(gt_id int64, category_id int64)(goodsList *[]Goods, err error)  {
	if gt_id < 1 {
		err = fmt.Errorf("illegal gid")
		logs.Error(err)
		return
	}
	logs.Info("ok")
	var rows *sql.Rows
	if gt_id == 0 {
		queryStr := "select t_user.username, gid, uid, gname, gprice, gdetail,click_number, status, release_time from t_goods, t_user "  +
			"where  category_id=$1 and status=$2 and t_goods.uid=t_user.uid;"
		rows, err = db.Query(queryStr, category_id, 1)
		if err != nil {
			logs.Error(err)
			return
		}
	}else {
		queryStr := "select t_user.username, gid, uid, gname, gprice, gdetail,click_number, status, release_time from t_goods, t_user "  +
			"where gt_id=$1 and  category_id=$2 and status=$3 and t_goods.uid=t_user.uid;"
		rows, err = db.Query(queryStr, gt_id, category_id, 1)
		if err != nil {
			logs.Error(err)
			return
		}
	}


	defer rows.Close()
	goodsList = new([]Goods)
	var good Goods
	var username string
	var gid int64
	var uid int64
	var gname string
	var gprice float64
	var gdetail string
	var status int64
	var click_number int64
	var release_time time.Time

	for rows.Next() {

		err = rows.Scan(&username, &gid, &uid,&gname, &gprice, &gdetail, &click_number, &status, &release_time)
		if err != nil {
			logs.Error(err)
			return
		}
		good.Username = username
		good.Gid = gid
		good.Uid = uid
		good.Gname = gname
		good.Gprice = gprice
		good.Gdetail = gdetail
		good.ClickNumber = click_number
		good.Status = status
		good.ReleaseTime = release_time

		good.CategoryId = category_id
		good.Gtid = gt_id



		*goodsList = append(*goodsList, good)
	}
	logs.Info(goodsList)
	return
}

