package controllers

import (
	"net/http"

	"../common"
	"../models"
	"github.com/astaxie/beego/logs"
)

type  userStatus struct{
	Status bool `json:"status"`
	Username string `json:"username"`
}
func Login(w http.ResponseWriter, r *http.Request) {
	logs.Info(r)
	if r.Method == "POST" {
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
		// 读取用户名和密码
		username, ok1 := data["username"].(string)
		password, ok2 := data["password"].(string)
		if !ok1 || !ok2 {
			logs.Warn("读取表单参数错误")
			ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
			return
		}
		// 判断账号或密码是否正确
		isExist, err := models.LoginVerification(username, password)
		logs.Info(isExist)
		if !isExist {
			if err := NaturalResp(w, r, 0, "账号或密码不正确", 0); err != nil {
				logs.Error(err)
				return
			}
			return
		}
	
		// 判断用户角色
		isAdmin, err := models.IsAdmin(username)
		if err != nil {
			logs.Error("数据库检索验证用户信息失败, err:", err)
			ErrorResp(w, r, "数据库检索验证用户信息失败", http.StatusInternalServerError)
			return
		}
		// 生成cookie
		newCookie, err := common.NewCookie(username, isAdmin)
		if err != nil {
			logs.Error(err)
			ErrorResp(w, r, err.Error(), http.StatusInternalServerError) //500
			return
		}
		logs.Info(newCookie)
		w.Header().Set("Set-Cookie", newCookie)
		if err := NaturalResp(w, r,1, "登录成功", 0); err != nil {
			logs.Error(err)
			return
		}
	}
	if r.Method == "GET" {
		logs.Info("login get11111111111")
		var token string
		for _, cookie := range r.Cookies() {
			token = cookie.Value
		}
		logs.Info(token)
		var resp userStatus
		username, err := common.ParseToken(token,[]byte("secretkey"))
		if err != nil || username == "" {
			//token无效的业务流程，根据自己的需要修改

			resp.Status = false
			resp.Username = ""
			if err := NaturalResp(w, r, resp, "该用户token失效或不处于登录状态", 0); err != nil {
				logs.Error(err)
				return
			}
			return
		}
		////若有效则生刷新Token存入NewCookie
		//// 判断用户角色
		//isAdmin, err := models.IsAdmin(username)
		//if err != nil {
		//	logs.Error("数据库检索验证用户信息失败, err:", err)
		//	ErrorResp(w, r, "数据库检索验证用户信息失败", http.StatusInternalServerError)
		//	return
		//}
		//// 生成cookie
		//newCookie, err := common.NewCookie(username, isAdmin)
		//if err != nil {
		//	logs.Error(err)
		//	ErrorResp(w, r, err.Error(), http.StatusInternalServerError) //500
		//	return
		//}
		//logs.Info(newCookie)
		//w.Header().Set("Set-Cookie", newCookie)
		resp.Status = true
		resp.Username = username
		logs.Info(resp)
		if err := NaturalResp(w, r, resp, "该用户处于登录状态", 0); err != nil {
			logs.Error(err)
			return
		}
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
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
	// 读取用户名和密码
	username, ok1 := data["username"].(string)
	password, ok2 := data["password"].(string)
	if !ok1 || !ok2 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}
	isUserExist, err := models.IsUserExist(username)
	if err != nil {
		logs.Error("数据库检索验证用户信息失败, err:", err)
		ErrorResp(w, r, "数据库检索验证用户信息失败", http.StatusInternalServerError)
		return
	}
	respData1 := struct {
		IsUserExist bool
	}{
		isUserExist,
	}
	if isUserExist {
		if err := NaturalResp(w, r, respData1, "该用户名已被使用", 0); err != nil {
			logs.Error(err)
			return
		}
	}
	err = models.RegisterUser(username, password)
	if err != nil {
		logs.Error("注册失败，err:", err)
		ErrorResp(w, r, "注册失败", http.StatusInternalServerError)
		return
	}
	respData2 := struct {
		success bool
	}{
		true,
	}
	if err := NaturalResp(w, r, respData2, "注册成功", 0); err != nil {
		logs.Error(err)
		return
	}
}

// 注销
func Logout(w http.ResponseWriter, r *http.Request) {
	// 生成cookie
	newCookie, err := common.NewCookie("注销账号", false)
	if err != nil {
		logs.Error(err)
		ErrorResp(w, r, err.Error(), http.StatusInternalServerError) //500
		return
	}
	w.Header().Set("Set-Cookie", newCookie)
	if err := NaturalResp(w, r, "", "注销成功", 0); err != nil {
		logs.Error(err)
		return
	}
}

// 获取私信
func GetPrivateLetter(w http.ResponseWriter, r *http.Request) {
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
	// 读取用户id
	_, ok1 := data["uid"].(int)
	if !ok1 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}

	// 根据uid获取与其相关的所有私信
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
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
	// 读取用户名和密码
	username, ok1 := data["username"].(string)
	if !ok1  {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}
	user, err := models.GetUserInfo(username)
	if err != nil {
		logs.Error("数据库获取用户失败, err:", err)
		ErrorResp(w, r, "数据库获取用户失败", http.StatusInternalServerError)
		return
	}
	if err := NaturalResp(w, r, user, "用户信息获取成功", 0); err != nil {
		logs.Error(err)
		return
	}
}