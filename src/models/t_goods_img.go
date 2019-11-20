package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

// 添加图片
func AddImage(gid int,imageName, imageExt, savePath string, imageSize float64) (bool, error){
	if gid < 1 || len(imageName) == 0 || len(savePath) == 0 || imageSize <= 0{
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	stmt, err := db.Exec("insert into t_goods_img(gid, image_name,image_ext,save_path,image_size) VALUES ($1,$2,$3,$4,$5)",
		gid,imageName,imageExt,savePath,imageSize)
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
	return true, nil
}
// 删除图片
// id 图片id
func DeleteImage(id int)(bool, error) {
	if id < 1 {
		err := fmt.Errorf("illegal param")
		logs.Error(err)
		return false, err
	}
	//删除数据
	stmt, err := db.Prepare("delete from t_goods_img where id=$1")
	if err != nil {
		logs.Error(err)
		return false, err
	}

	res, err := stmt.Exec(id)
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
// 获取图片
func GetImages(goodId int)(goodImagesList *[]GoodsImg, err error) {
	if goodId < 1 {
		err = fmt.Errorf("illegal goodId")
		logs.Error(err)
		return
	}
	rows, err := db.Query("select id, gid, image_name,image_ext,save_path,image_size,release_time from t_goods_img where gid=$1", goodId)
	if err != nil {
		logs.Error(err)
		return
	}
	defer rows.Close()
	goodImagesList = new([]GoodsImg)
	var gi GoodsImg
	for rows.Next() {
		var id int
		var gid int
		var image_name string
		var image_ext string
		var save_path string
		var image_size float64
		var release_time time.Time
		err = rows.Scan(&id,&gid,&image_name,&image_ext,&save_path,&image_size,&release_time,&goodId)
		if err != nil {
			logs.Error(err)
			return
		}
		gi.Id = id
		gi.Gid= gid
		gi.ImageName = image_name
		gi.ImageExt = image_ext
		gi.SavePath = save_path
		gi.ImageSize = image_size
		gi.ReleaseTime = release_time
		*goodImagesList = append(*goodImagesList, gi)
	}
	return
}