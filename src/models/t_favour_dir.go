package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

func AddFavourDir(foldername string, sketch string, authorityLevel int) {

}

func RemoveFavourDir() {
	
}

func GetFolder(uid int)(favourDirList *[]FavourDir, err error)  {
	if uid < 1 {
		err = fmt.Errorf("illegal gid")
		logs.Error(err)
		return
	}
	rows, err := db.Query("select fd_id, foldername, sketch, authority_level, create_time from t_favour_dir where uid=$1", uid)
	if err != nil {
		logs.Error(err)
		return
	}
	defer rows.Close()
	favourDirList = new([]FavourDir)
	var dir FavourDir
	var fdid  sql.NullInt64
	var foldername  sql.NullString
	var createTime  time.Time
	var sketch  sql.NullString
	var authorityLevel sql.NullInt64
	for rows.Next() {

		err = rows.Scan(&fdid, &foldername, &sketch, &authorityLevel, &createTime)
		if err != nil {
			logs.Error(err)
			return
		}
		dir.Fdid = fdid.Int64
		dir.Foldername = foldername.String
		dir.Uid = int64(uid)
		dir.Sketch = sketch.String
		dir.AuthorityLevel = authorityLevel.Int64
		dir.CreateTime = createTime
		*favourDirList = append(*favourDirList, dir)
	}
	return
}

type favourGood struct {
	Gid  int `json:"gid"`
	Gname  string `json:"gname"`
	CreateTime time.Time `json:"create_time"`
}
func GetFavourGoods(fdid int, key string)(favourGoodList *[]favourGood, err error)  {
	if fdid < 1 {
		err = fmt.Errorf("illegal gid")
		logs.Error(err)
		return
	}
	logs.Info("ok")
	key = "%" + key + "%"
	rows, err := db.Query("select t_goods.gid, gname,  t_favour.create_time from t_goods, t_favour where fd_id=$1 and t_favour.gid = t_goods.gid and gname like $2;", fdid, key)
	if err != nil {
		logs.Error(err)
		return
	}
	defer rows.Close()
	favourGoodList = new([]favourGood)
	var good favourGood

	var createTime  time.Time
	var gname  string
	var gid int
	for rows.Next() {

		err = rows.Scan(&gid, &gname, &createTime)
		if err != nil {
			logs.Error(err)
			return
		}
		good.Gid = gid
		good.Gname = gname
		good.CreateTime = createTime
		*favourGoodList = append(*favourGoodList, good)
		logs.Info(favourGoodList)
	}
	logs.Info(favourGoodList)
	return
}

