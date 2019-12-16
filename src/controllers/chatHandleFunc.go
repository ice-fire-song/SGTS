package controllers

import (
	"github.com/astaxie/beego/logs"
	"net/http"
	"../models"
)

func GetSenderList(w http.ResponseWriter, r *http.Request) {
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

	respData, err := models.GetSenderList(int(uid))
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
func GetRecords(w http.ResponseWriter, r *http.Request) {
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
	user_id, ok2 := data["user_id"].(float64)
	friend_id, ok3 := data["friend_id"].(float64)
	if !ok2 || !ok3 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}

	respData, err := models.GetRecords(int64(user_id), int64(friend_id))
	logs.Info(respData)
	if err != nil {
		logs.Error("获取聊天记录失败，err:", err)
		ErrorResp(w, r, "获取聊天记录失败", http.StatusInternalServerError)
		return
	}

	if err := NaturalResp(w, r, respData, "获取聊天记录成功", 0); err != nil {
		logs.Error(err)
		return
	}
}

// 发送私信
func SendPrivateLetter(w http.ResponseWriter, r *http.Request) {
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
	sender_id, ok1 := data["user_id"].(float64)
	receiver_id, ok2 := data["friend_id"].(float64)
	messageContent, ok3 := data["message_content"].(string)
	if !ok1 || !ok2 || !ok3 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}
	var respData ActionRes
	respData.IsSuccess, err = models.SendPrivateLetter(int(sender_id), int(receiver_id),messageContent)
	logs.Info(respData)
	if err != nil {
		logs.Error("私信发送失败，err:", err)
		ErrorResp(w, r, "私信发送失败", http.StatusInternalServerError)
		return
	}

	if err := NaturalResp(w, r, respData, "私信发送成功", 0); err != nil {
		logs.Error(err)
		return
	}
}


// 将所有私信改为已读
func ToReaded(w http.ResponseWriter, r *http.Request) {
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
	// 读取user_id, friend_id
	user_id, ok1 := data["user_id"].(float64)
	friend_id, ok2 := data["friend_id"].(float64)

	if !ok1 || !ok2 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}
	var respData ActionRes
	respData.IsSuccess, err = models.ToReaded(int(user_id),int(friend_id))
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