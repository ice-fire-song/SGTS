package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"time"
)

//后端响应数据通信协议
type ReplyProto struct {
	Status    int         `json:"status"` //状态 0正常，小于0出错，大于0可能有问题
	Msg       string      `json:"msg"`    //状态信息
	Data      interface{} `json:"data"`
	API       string      `json:"API"`        //api接口
	Method    string      `json:"method"`     //post,put,get,delete
	RowCount  int         `json:"rowCount"`   //Data若是数组，算其长度
	Time      int64       `json:"time"`       //请求响应时间，毫秒
	//CheckTime int64       `json:"check_time"` //检测时间，毫秒
}

//前端请求数据通讯协议
type ReqProto struct {
	Action   string      `json:"action"` //请求类型GET/POST/PUT/DELETE
	Data     interface{} `json:"data"`   //请求数据
	Sets     []string    `json:"sets"`
	OrderBy  string      `json:"orderBy"`  //排序要求
	Filter   string      `json:"filter"`   //筛选条件
	Page     int         `json:"page"`     //分页
	PageSize int         `json:"pageSize"` //分页大小
}

// 正常响应
func NaturalResp(w http.ResponseWriter, r *http.Request, data interface{}, msg string, rowCount int) error {
	if w == nil || r == nil || data == nil {
		err := fmt.Errorf("arguments can not be a nil value")
		logs.Error(err)
		return err
	}
	if rowCount < 0 {
		err := fmt.Errorf("rowCount must be positive")
		logs.Error(err)
		return err
	}
	startTime := r.Context().Value("startTime")
	last := time.Duration(0)
	if startTime != nil {
		_, ok := startTime.(time.Time)
		if ok {
			last = time.Since(startTime.(time.Time))
		}
	}
	resp := ReplyProto{}
	resp.Status = 0
	resp.Msg = msg
	resp.Data = data
	resp.API = r.RequestURI
	resp.Method = r.Method
	resp.RowCount = rowCount
	resp.Time = last.Nanoseconds() / 1000000
    logs.Info(data)
	logs.Info(resp)
	response, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = w.Write(response)
	if err != nil {
		return err
	}

	return nil
}
// 异常响应
func ErrorResp(w http.ResponseWriter, r *http.Request, msg string, statusCode int){
	if w == nil || r == nil {
		err := fmt.Errorf("arguments can not be a nil value")
		logs.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//响应时长
	startTime := r.Context().Value("startTime")
	last := time.Duration(0)
	if startTime != nil {
		_, ok := startTime.(time.Time)
		if ok {
			last = time.Since(startTime.(time.Time))
		}
	}

	resp := ReplyProto{}
	resp.Status = -1
	resp.Msg = msg
	resp.Data = nil
	resp.API = r.RequestURI
	resp.Method = r.Method
	resp.RowCount = 0
	resp.Time = last.Nanoseconds() / 1000000

	response, err := json.Marshal(resp)
	if err != nil {
		logs.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	_, err = w.Write(response)
	if err != nil {
		logs.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func GetBodyData(r *http.Request) (data map[string]interface{}, err error) {
	if r == nil {
		err = fmt.Errorf("Call GetBodyData with a empty r")
		logs.Error(err)
		return nil, err
	}
	logs.Info(r)
	logs.Info(r.Body)
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if len(body) == 0 {
		err = fmt.Errorf("r.Body is nil")
		logs.Error(err)
		return nil, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return data, nil
}
