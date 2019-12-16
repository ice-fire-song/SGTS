package controllers

import (
	"github.com/astaxie/beego/logs"
	"net/http"
	"../models"
)

// 发布货品
func UploadGood(w http.ResponseWriter, r *http.Request) {
	body, err := GetBodyData(r)
	if err != nil {
		logs.Error(err)
		ErrorResp(w, r, err.Error(), http.StatusInternalServerError) //500
		return
	}
	logs.Info(body)
	data, ok := body["data"].(map[string]interface{})
	if !ok {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest) //400
		return
	}
	logs.Info(data)

	// 读取货品gid
	var good models.Goods

	gname, ok1 := data["gname"].(string)
	gprice, _ := data["gprice"].(float64)
	gdetail, ok3 := data["gdetail"].(string)
	category_id, ok4 := data["category_id"].(float64)
	gt_id, ok5 := data["gt_id"].(float64)
	//tabs, ok6 := data["tabs"]
	images, _ := data["images"].(*[]models.Image)
	logs.Info("image", images)
	first_img_path, _ := data["first_img_path"].(string)
	uid, ok8 := data["uid"].(float64)
	good.Uid = int64(uid)
	good.Gname = gname
	good.Gprice = gprice
	good.Gdetail = gdetail
	good.CategoryId = int64(category_id)
	good.Gtid = int64(gt_id)
	good.FirstImgPath = first_img_path
	if !ok1 ||  !ok3 || !ok4 || !ok5 || !ok8 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}

	_, err = models.AddGood(&good)
	if err != nil {
		logs.Error("数据库插入数据失败, err:", err)
		ErrorResp(w, r, "数据库插入数据失败", http.StatusInternalServerError)
		return
	}

	if err := NaturalResp(w, r, true, "数据库插入数据成功", 0); err != nil {
		logs.Error(err)
		return
	}
}
// 按照货品类型：免费商品，商品，需求
func GetGoodsByCategory(w http.ResponseWriter, r *http.Request) {
	body, err := GetBodyData(r)
	if err != nil {
		logs.Error(err)
		ErrorResp(w, r, err.Error(), http.StatusInternalServerError) //500
		return
	}
	logs.Info(body)
	data, ok := body["data"].(map[string]interface{})
	if !ok {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest) //400
		return
	}
	logs.Info(data)
	// 读取货品gid
	uid, ok1 := data["uid"].(float64)
	category_id, ok2 := data["category_id"].(float64)
	key, ok3 := data["key"].(string)
	if !ok1 || !ok2 || !ok3 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}

	respData, err := models.GetGoodsByCategory(int64(uid), int64(category_id), key)
	logs.Info(respData)
	if err != nil {
		logs.Error("数据库获取用户货品列表失败, err:", err)
		ErrorResp(w, r, "数据库获取用户货品失败", http.StatusInternalServerError)
		return
	}
	if err := NaturalResp(w, r, respData, "用户货品列表获取成功", 0); err != nil {
		logs.Error(err)
		return
	}
}
func GetGoodInfo(w http.ResponseWriter, r *http.Request) {
	body, err := GetBodyData(r)
	if err != nil {
		logs.Error(err)
		ErrorResp(w, r, err.Error(), http.StatusInternalServerError) //500
		return
	}
	logs.Info(body)
	data, ok := body["data"].(map[string]interface{})
	if !ok {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest) //400
		return
	}
	logs.Info(data)
	// 读取货品gid
	gid, ok1 := data["gid"].(float64)
	if !ok1 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}

	respData, err := models.GetGoodInfo(int64(gid))
	logs.Info(respData)
	if err != nil {
		logs.Error("数据库获取收藏列表失败, err:", err)
		ErrorResp(w, r, "数据库获取收藏列表失败", http.StatusInternalServerError)
		return
	}
	if err := NaturalResp(w, r, respData, "收藏列表获取成功", 0); err != nil {
		logs.Error(err)
		return
	}
}

// 获取指定货品的所有图片
func GetGoodImg(w http.ResponseWriter, r *http.Request) {
	body, err := GetBodyData(r)
	if err != nil {
		logs.Error(err)
		ErrorResp(w, r, err.Error(), http.StatusInternalServerError) //500
		return
	}
	logs.Info(body)
	data, ok := body["data"].(map[string]interface{})
	if !ok {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest) //400
		return
	}
	logs.Info(data)
	// 读取货品gid
	gid, ok1 := data["gid"].(float64)
	if !ok1 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}

	respData, err := models.GetGoodImg(int64(gid))
	logs.Info(respData)
	if err != nil {
		logs.Error("数据库获取收藏列表失败, err:", err)
		ErrorResp(w, r, "数据库获取收藏列表失败", http.StatusInternalServerError)
		return
	}
	if err := NaturalResp(w, r, respData, "收藏列表获取成功", 0); err != nil {
		logs.Error(err)
		return
	}
}

type ActionRes struct {
	IsSuccess bool `json:"isSuccess"`
}
// 改变货品的状态
func ModifyGoodStatus(w http.ResponseWriter, r *http.Request) {
	body, err := GetBodyData(r)
	if err != nil {
		logs.Error(err)
		ErrorResp(w, r, err.Error(), http.StatusInternalServerError) //500
		return
	}
	logs.Info(body)
	data, ok := body["data"].(map[string]interface{})
	if !ok {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest) //400
		return
	}
	logs.Info(data)
	// 读取货品gid
	gid, ok1 := data["gid"].(float64)
	good_status, ok2 := data["good_status"].(float64)
	if !ok1 || !ok2 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}
	var respData ActionRes
	respData.IsSuccess, err = models.ModifyGoodStatus(int64(gid), int64(good_status))
	logs.Info(respData)
	if err != nil {
		logs.Error("数据库获取收藏列表失败, err:", err)
		ErrorResp(w, r, "数据库获取收藏列表失败", http.StatusInternalServerError)
		return
	}
	if err := NaturalResp(w, r, respData, "收藏列表获取成功", 0); err != nil {
		logs.Error(err)
		return
	}
}