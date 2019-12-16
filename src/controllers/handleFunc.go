package controllers

import (
	"net/http"
	"../common"
	"../models"
	"github.com/astaxie/beego/logs"
)

type  userStatus struct{ // 用户登录状态是否有效
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
		logs.Info("是否管理员",isExist,isAdmin)
		if err := NaturalResp(w, r,isAdmin, "登录成功", 0); err != nil {
			logs.Error(err)
			return
		}
	}
	if r.Method == "GET" {
		logs.Info("login get")
		var token string
		cookie, err := r.Cookie("token")
		if err != nil {
			logs.Error(err)
		}
		token = cookie.Value
		logs.Info("token", token)
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
	mailbox, ok3 := data["mailbox"].(string)
	if !ok1 || !ok2 || !ok3 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}
	// 判断该用户名是否以被使用
	respData := struct {
		IsUserExist bool `json:"is_user_exist`
		success bool `json:"success"`
	}{
		false,
		false,
	}
	isUserExist, err := models.IsUserExist(username)
	if err != nil {
		logs.Error("数据库检索验证用户信息失败, err:", err)
		ErrorResp(w, r, "数据库检索验证用户信息失败", http.StatusInternalServerError)
		return
	}

	if isUserExist {
		respData.IsUserExist = true
		if err := NaturalResp(w, r, respData, "该用户名已被使用", 0); err != nil {
			logs.Error(err)
			return
		}
	}

	// 注册
	err = models.RegisterUser(username, password,mailbox)
	if err != nil {
		logs.Error("注册失败，err:", err)
		ErrorResp(w, r, "注册失败", http.StatusInternalServerError)
		return
	}
    respData.success = true
	if err := NaturalResp(w, r, respData, "注册成功", 0); err != nil {
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

// 修改个人信息
func ModifyUserInfo(w http.ResponseWriter, r *http.Request) {
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
	uid, ok1 := data["uid"].(float64)
	head_sculpture_path, ok2 := data["head_sculpture_path"].(string)
	label, ok3 := data["label"].(string)
	oldPassword, ok4 := data["oldPassword"].(string)
	newPassword, ok5 := data["newPassword"].(string)
	mailbox, ok6 := data["mailbox"].(string)
	if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 || !ok6 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}

	var respData ActionRes
	respData.IsSuccess, err = models.ModifyUserInfo(int(uid), head_sculpture_path, label, mailbox)
	if err != nil {
		logs.Error("修改用户信息失败, err:", err)
		ErrorResp(w, r, "修改用户信息失败", http.StatusInternalServerError)
		return
	}

	pwdRes, err := models.ModifyPWD(int(uid),oldPassword, newPassword)
	if err != nil {
		logs.Error(err)
	}
	var msg string
	if pwdRes {
		msg = "修改用户信息成功、密码修改成功"
	}else {
		msg =  "修改用户信息成功、密码修改失败"
	}
	if err := NaturalResp(w, r, respData, msg, 0); err != nil {
		logs.Error(err)
		return
	}
}