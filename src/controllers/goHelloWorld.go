package controllers

import (
	"../models"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
)

func GoHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("go hello world"))
}

// 获取货品种类
func GetGoodsType(w http.ResponseWriter, r *http.Request) {

	respData, err := models.GetGoodsType()
	logs.Info(respData)
	if err != nil {
		logs.Error("数据库中获取货品类别失败, err:", err)
		ErrorResp(w, r, "数据库中获取货品类别失败", http.StatusInternalServerError)
		return
	}
	if err := NaturalResp(w, r, respData, "获取货品类别成功", 0); err != nil {
		logs.Error(err)
		return
	}
}

// 主页：根据货品种类id、货品的类别获取货品
func GetGoodsByType(w http.ResponseWriter, r *http.Request) {
	body, err := GetBodyData(r)
	if err != nil {
		logs.Error(err)
		ErrorResp(w, r, err.Error(), http.StatusInternalServerError) //500
		return
	}
	logs.Info("调用GetGoodsByType")
	logs.Info(body)
	data, ok := body["data"].(map[string]interface{})
	if !ok {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest) //400
		return
	}
	logs.Info(data)
	key, _ := data["key"].(string)
	gt_id, ok2 := data["gt_id"].(float64)
	category_id, ok3 := data["category_id"].(float64)
	if !ok2 || !ok3 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}

	str := fmt.Sprintf("获取gt_id：%f, category_id：%f的货品", gt_id, category_id)
	respData, err := models.GetGoodsByType(int64(gt_id), int64(category_id), key)
	logs.Info(respData)
	if err != nil {
		str += "失败"
		logs.Error(str, " err:", err)
		ErrorResp(w, r, str, http.StatusInternalServerError)
		return
	}
	str += "成功"
	if err := NaturalResp(w, r, respData, str, 0); err != nil {
		logs.Error(err)
		return
	}
}