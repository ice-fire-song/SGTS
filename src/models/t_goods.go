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
	//Openid  string `jso
	// n:"openid"`
	//Qq       string `json:"qq"`
	ReleaseTime time.Time `json:"release_time"`
	FirstImgPath string `json:"first_img_path"`
}

func AddGood(goodInfo *Goods) (bool, error) {
	if goodInfo == nil {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	stmt, err := db.Exec("insert into t_goods(uid,gname,gprice,gdetail,category_id,gt_id,first_img_path) VALUES ($1,$2,$3,$4,$5,$6,$7)",
		goodInfo.Uid,goodInfo.Gname,goodInfo.Gprice,goodInfo.Gdetail,goodInfo.CategoryId,goodInfo.Gtid,goodInfo.FirstImgPath)
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
func GetAllGoods(category_id int64, label int64, keyWord string)(goodsList *[]Good, err error) {
	var query string
	var rows *sql.Rows
	if category_id == -1 {
		query = "select(gid,uid,gtype,gname,gprice,gdetail,category_id,status,release_time,first_img_path)" +
			"from t_goods where label=$1 and gname like $2 order by send_time asc"
		rows, err = db.Query(query, label, keyWord)
	}else if label == 0 {
		query = "select(gid,uid,gtype,gname,gprice,gdetail,category_id,status,release_time,first_img_path)" +
			"from t_goods where gtype=$1 and gname like $2 order by send_time asc"
		rows, err = db.Query(query, category_id, keyWord)
	}else {
		query = "select(gid,uid,gtype,gname,gprice,gdetail,category_id,status,release_time,first_img_path)" +
			"from t_goods where gtype=$1 and gname like $2 gid = (select gid from t_goods_label where label=$3)order by send_time asc"
		rows, err = db.Query(query, category_id, keyWord,label)
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
		var first_img_path string
		err = rows.Scan(&gid,&uid,&gtype,&gname,&gprice,&gdetail,&category_id,&status,&release_time,&first_img_path)
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
		good.FirstImgPath = first_img_path
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

// 获取货品的信息
func GetGoodInfo(gid int64) (good Goods, err error){
	queryStr := "select t_user.username, gid, uid, gname, gprice, gdetail,category_id,click_number, status, release_time,type_name,first_img_path from t_goods, t_user, t_goods_type "  +
		"where  gid=$1 and t_goods.uid=t_user.uid and t_goods.gt_id=t_goods_type.gt_id;"
	rows, err := db.Query(queryStr, gid)
	if err != nil {
		logs.Error(err)
		return
	}
	defer rows.Close()

	var username string
	var uid int64
	var gname string
	var gprice float64
	var gdetail string
	var status int64
	var click_number int64
	var release_time time.Time
	var type_name string
	var category_id int64
	var category string
	var first_img_path string
	for rows.Next() {

		err = rows.Scan(&username, &gid, &uid, &gname, &gprice, &gdetail, &category_id, &click_number, &status, &release_time,&type_name,&first_img_path)
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
		good.CategoryId = category_id
		switch category_id {
		case 0:
			category = "免费商品"
		case 1:
			category = "商品"
		case 2:
			category = "需求"
		}
		good.Category = category
		good.ClickNumber = click_number
		good.Status = status
		good.ReleaseTime = release_time
		good.TypeName = type_name
		good.FirstImgPath = first_img_path
	}
	return
}

// 获取指定货品gid的所有图片
func GetGoodImg(gid int64)(imgList *[]GoodsImg, err error)  {
	if gid < 1 {
		err = fmt.Errorf("illegal gid")
		logs.Error(err)
		return
	}
	logs.Info("ok")
    queryStr := "select id, gid, image_name,image_ext, save_path, image_size,release_time from t_goods_img where gid=$1"

	rows, err := db.Query(queryStr, gid)
	if err != nil {
		logs.Error(err)
		return
	}
	defer rows.Close()

	imgList = new([]GoodsImg)
	var img GoodsImg
	var id sql.NullInt64
	var image_name sql.NullString
	var image_ext sql.NullString
	var save_path sql.NullString
	var image_size sql.NullFloat64
	var release_time time.Time

	for rows.Next() {

		err = rows.Scan(&id, &gid, &image_name,&image_ext, &save_path, &image_size, &release_time)
		if err != nil {
			logs.Error(err)
			return
		}
		img.Id = id.Int64
		img.Gid = gid
		img.ImageName = image_name.String
		img.ImageExt = image_ext.String
		img.SavePath = save_path.String
		img.ImageSize = image_size.Float64
        img.ReleaseTime = release_time
		*imgList = append(*imgList, img)
	}
	logs.Info(imgList)
	return
}

// 货品管理
func GetGoodsByCategory(uid int64, category_id int64, key string)(goodsList *[]Goods, err error)  {
	logs.Info("ok")
	var rows *sql.Rows
    key = "%" + key + "%"
    logs.Info(key)

	rows, err = db.Query("select  gid, uid, gname, gprice, gdetail,click_number, status, release_time, gt_id,first_img_path from t_goods where  category_id=$1 and status=$2 and uid=$3 and gname like $4;", category_id, 1, uid, key)
	if err != nil {
		logs.Error(err)
		return
	}


	defer rows.Close()
	goodsList = new([]Goods)
	var good Goods
	var gid int64
	var gname string
	var gprice float64
	var gdetail string
	var status int64
	var click_number int64
	var gt_id int64
	var release_time time.Time
    var first_img_path string
	for rows.Next() {

		err = rows.Scan(&gid, &uid,&gname, &gprice, &gdetail, &click_number, &status, &release_time, &gt_id, &first_img_path)
		if err != nil {
			logs.Error(err)
			return
		}
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
		good.FirstImgPath = first_img_path
		*goodsList = append(*goodsList, good)
	}
	logs.Info(goodsList)
	return
}

// 改变货品状态
func ModifyGoodStatus(gid, good_status int64)(bool, error) {
	if gid < 1  {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	stmt, err := db.Prepare("update t_goods set status=$1 where gid=$2")
	if err != nil {
		logs.Error(err)
		return false, err
	}
	res, err := stmt.Exec(good_status, gid)
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