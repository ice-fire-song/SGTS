package controllers

import (
	"../models"
	"github.com/astaxie/beego/logs"
	"net/http"
)
func AddFavour(w http.ResponseWriter, r *http.Request) {
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
	// 读取用户名
	uid, ok1 := data["uid"].(float64)
	gid, ok2 := data["gid"].(float64)
	if !ok1 || !ok2 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}
    var respData ActionRes
	respData.IsSuccess, err = models.AddFavour(int(uid),int(gid))
	logs.Info(respData)
	if err != nil {
		logs.Error("数据库增加收藏失败, err:", err)
		ErrorResp(w, r, "数据库增加收藏失败", http.StatusInternalServerError)
		return
	}
	if err := NaturalResp(w, r, respData, "增加收藏成功", 0); err != nil {
		logs.Error(err)
		return
	}
}
func FavourControll(w http.ResponseWriter, r *http.Request) {
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

	action, ok1 := data["action"].(int) //0表示取消收藏，1表示添加收藏，2查看收藏状态
	uid, ok2 := data["uid"].(int) //用户id
	gid, ok3 := data["gid"].(int) //货品id
	fid, ok4 := data["fid"].(float64)
	if !ok1 || !ok2 || !ok3 || !ok4 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}
	
	var respData interface{}
	var msg string
	if action == 0 {
		_, err = models.CancelFavour(int(fid))
		if err != nil {
			msg = "取消收藏失败"
			logs.Error("取消收藏失败，err:", err)
		}else {
			msg = "取消收藏成功"
			respData = struct{ success bool }{ true }
		}
        
	}else if action == 1 {
		_, err = models.AddFavour(gid, uid)
		if err != nil {
			msg = "添加收藏失败"
			logs.Error("添加收藏失败，err:", err)
		}else {
			msg = "添加收藏成功"
			respData = struct{ success bool }{ true }
		}	
	}else if action == 2 {
		status, err := models.SeeFavourStatus(int64(gid), int64(uid))
		if err != nil {
			msg = "查看收藏状态失败"
			logs.Error("查看收藏状态失败，err:", err)
		}else {
			msg = "查看收藏状态成功"
			respData = struct{ status bool }{ status }//true表示已收藏
		}	
	}else {
		if err := NaturalResp(w, r, "", "请求信息有误", 0); err != nil {
			logs.Error(err)
			return
		}
	}
	if err != nil {
		ErrorResp(w, r, msg, http.StatusInternalServerError)
		return
	}
	if err := NaturalResp(w, r, respData, msg, 0); err != nil {
		logs.Error(err)
		return
	}
}

// 获取用户收藏夹列表
func GetFolder(w http.ResponseWriter, r *http.Request) {
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
	// 读取用户名

	uid, ok1 := data["uid"].(float64)
	if !ok1 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}

	respData, err := models.GetFolder(int(uid))
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

// 获取指定收藏夹的收藏货品
func GetFavourGoods(w http.ResponseWriter, r *http.Request) {
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
	//uid, ok1 := data["uid"].(float64)
	fdid, ok2 := data["fdid"].(float64)
	key, ok3 := data["key"].(string)
	if !ok2 || !ok3 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}
	str := "获取" + string(int(fdid)) + "中的收藏货品"
	respData, err := models.GetFavourGoods(int(fdid), key)
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
type favourStatus struct {
	IsFavour bool `json:"isFavour"`
}
func SeeFavourStatus (w http.ResponseWriter, r *http.Request) {
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
	gid, ok2 := data["gid"].(float64)
	uid, ok3 := data["uid"].(float64)
	if !ok2 || !ok3 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}
	var respData favourStatus
	respData.IsFavour, err = models.SeeFavourStatus(int64(gid), int64(uid))
	if err != nil {
		logs.Error("获取货品的收藏状态失败， err:", err)
		ErrorResp(w, r, "获取货品的收藏状态失败", http.StatusInternalServerError)
		return
	}

	if err := NaturalResp(w, r, respData, "获取货品的收藏状态成功", 0); err != nil {
		logs.Error(err)
		return
	}
}

// 取消指定收藏货品
func DeleteFavour(w http.ResponseWriter, r *http.Request) {
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
	fid, ok1 := data["fid"].(float64)

	if !ok1 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}
	var respData ActionRes
	respData.IsSuccess, err = models.CancelFavour(int(fid))
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

// 删除收藏夹
// 取消指定收藏货品
func DeleteDir(w http.ResponseWriter, r *http.Request) {
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
	fd_id, ok1 := data["fd_id"].(float64)

	if !ok1 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}
	var respData ActionRes
	respData.IsSuccess, err = models.DeleteDir(int(fd_id))
	logs.Info(respData)
	if err != nil {
		logs.Error("数据库删除收藏夹失败, err:", err)
		ErrorResp(w, r, "数据库删除收藏夹失败", http.StatusInternalServerError)
		return
	}
	if err := NaturalResp(w, r, respData, "收藏夹删除成功", 0); err != nil {
		logs.Error(err)
		return
	}
}