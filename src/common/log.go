package common

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"log"
	"os"
)

// 初始化日志
func InitLogger() error {

	// 支持文件
	//err := os.Mkdir("./log/", os.ModePerm)
	//if err != nil {
	//	log.Fatal("创建日志目录失败")
	//	return err
	//}
	isExist, err := PathExists("./log/log.log")
	if err != nil {
		logs.Error(err)
		return err
	}
	if !isExist {
		fileName := "./log/log.log"
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal("创建日志文件失败")
			return err
		}
		defer f.Close()
	}

	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	config := make(map[string]interface{})
	config["filename"] = `./log/log.log`
	configStr, err := json.Marshal(config)
	if err != nil {
		log.Fatal("log file's config marshal failed, err:", err)
		return err
	}
	logs.SetLevel(logs.LevelDebug)
	err = logs.SetLogger(logs.AdapterFile, string(configStr))
	if err != nil {
		log.Fatal("set logger of file failed, err:", err)
		return err
	}
	// 支持控制台
	err = logs.SetLogger(logs.AdapterConsole, "")
	if err != nil {
		log.Fatal("set logger of console failed, err:", err)
		return err
	}
	return nil
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
